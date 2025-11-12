package storage

import (
	"context"
	"firstapi/internal/models"

	"github.com/jackc/pgx/v5"
)

type TodoStorage struct {
	c *pgx.Conn
}

func (t *TodoStorage) GetAll() ([]models.TodoResponse, error) {
	rows, err := t.c.Query(context.Background(), "SELECT t.id,t.name,t.description,u.name FROM todo t inner join users u on u.id = t.user_id")
	if err != nil {
		return nil, err
	}
	todos, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TodoResponse])
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodoStorage) CreateTodo(r models.TodoRequest) error {
	_, err := t.c.Exec(context.Background(), "INSERT INTO todo(name, description, user_id) VALUES ($1,$2,$3)")
	if err != nil {
		return err
	}
	return nil
}

func (t *TodoStorage) UpdateTodo(r models.TodoRequest, id int) error {
	_, err := t.c.Exec(context.Background(), "UPDATE todo SET name=$1,description=$2,user_id=$3", r.Name, r.Description, r.User_id)
	if err != nil {
		return err
	}
	return nil
}
