package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
)

func Newconnect(connstring string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		return nil, errors.New("Error connecting to db")
	}
	return conn, err
}
