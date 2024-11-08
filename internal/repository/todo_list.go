package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"todo/internal/domain"
)

type TodoListRepository struct {
	db *sqlx.DB
}

func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

func (r *TodoListRepository) Create(userID int, list domain.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createTodoListQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		TodoListsTable,
	)
	row := tx.QueryRow(createTodoListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserListQuery := fmt.Sprintf(
		"INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id",
		UsersListsTable,
	)
	_, err = tx.Exec(createUserListQuery, userID, id)
	if err != nil {
		tx.Rollback()
	}
	return id, tx.Commit()
}

func (r *TodoListRepository) All(userID int) ([]domain.TodoList, error) {
	var lists []domain.TodoList
	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description 
				FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id 
                WHERE ul.user_id = $1`,
		TodoListsTable,
		UsersListsTable,
	)
	err := r.db.Select(&lists, query, userID)
	return lists, err
}

func (r *TodoListRepository) ListByID(userID int, listID int) (domain.TodoList, error) {
	var list domain.TodoList
	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description
				FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id 
				WHERE ul.user_id = $1 AND ul.list_id = $2`,
		TodoListsTable,
		UsersListsTable,
	)
	err := r.db.Get(&list, query, userID, listID)
	return list, err
}

func (r *TodoListRepository) Update(userID int, listID int, input domain.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ",")
	query := fmt.Sprintf(
		"UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		TodoListsTable, setQuery, UsersListsTable, argID, argID+1,
	)

	args = append(args, listID, userID)
	logrus.Debugf("Update query: %s", query)
	logrus.Debugf("Update args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoListRepository) Delete(userID int, listID int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		TodoListsTable,
		UsersListsTable,
	)
	_, err := r.db.Exec(query, userID, listID)

	return err
}
