migrate up:
	goose -dir schema postgres "postgresql://postgres:postgres@127.0.0.1:5432/todo?sslmode=disable" up

#reverse only one migration per use
migrate down:
	goose -dir schema postgres "postgresql://postgres:postgres@127.0.0.1:5432/todo?sslmode=disable" down