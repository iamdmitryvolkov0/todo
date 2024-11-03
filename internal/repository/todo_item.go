package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"todo/internal/domain"
)

type ToDoItemRepository struct {
	db *sqlx.DB
}

func NewToDoItemRepository(db *sqlx.DB) *ToDoItemRepository {
	return &ToDoItemRepository{db: db}
}

func (r *ToDoItemRepository) Create(listId int, item domain.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", TodoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", ListsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *ToDoItemRepository) All(userID int, listID int) ([]domain.TodoItem, error) {
	var items []domain.TodoItem
	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
        INNER JOIN %s li on li.item_id = ti.id 
    	INNER JOIN %s ul on ul.list_id = li.list_id
    	WHERE li.list_id = $1 AND ul.user_id = $2`,
		TodoItemsTable, ListsItemsTable, UsersListsTable)

	err := r.db.Select(&items, query, listID, userID)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ToDoItemRepository) ItemByID(userID int, itemID int) (domain.TodoItem, error) {
	var item domain.TodoItem
	query := fmt.Sprintf(
		`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
        INNER JOIN %s li on li.item_id = ti.id
		INNER JOIN %s ul on ul.list_id = li.list_id 
        WHERE ti.id = $1 AND ul.user_id = $2`,
		TodoItemsTable, ListsItemsTable, UsersListsTable)
	if err := r.db.Get(&item, query, itemID, userID); err != nil {
		return item, err
	}

	return item, nil
}

func (r *ToDoItemRepository) Update(userID int, itemID int, input domain.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
		WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		TodoItemsTable, setQuery, ListsItemsTable, UsersListsTable, argId, argId+1)
	args = append(args, userID, itemID)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ToDoItemRepository) Delete(userID int, itemID int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
		WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		TodoItemsTable, ListsItemsTable, UsersListsTable)
	_, err := r.db.Exec(query, userID, itemID)
	return err
}
