package storage

import (
	"context"
	"errors"
	"firstapi/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserStorage struct {
	c *pgx.Conn
}

func (u *UserStorage) GetAll() ([]models.User, error) {
	rows, err := u.c.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, errors.New("failed to get users")
	}
	defer rows.Close()
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, errors.New("failed to get users")
	}
	return users, nil
}
func (u *UserStorage) CreateUser(name string) error {
	_, err := u.c.Exec(context.Background(), "INSERT INTO users(name) VALUES ($1)", name)
	if err != nil {
		return errors.New("Failed to insert user")
	}
	return nil
}
func (u *UserStorage) UpdateUser(id int, name string) error {
	_, err := u.c.Exec(context.Background(), "UPDATE users SET name =$1 WHERE id = $2", name, id)
	if err != nil {
		return errors.New("failed to update user")
	}
	return nil
}
func (u *UserStorage) DeleteUser(id int) error {
	_, err := u.c.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return errors.New("Failed to delete user")
	}
	return nil
}
