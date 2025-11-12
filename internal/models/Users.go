package models

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)
type User struct{
	id int  	`json:"id"`
	name string `json:"name"`
}
func (u *User) GetAll(conn *pgx.Conn) ([]User, error){
	rows,err := conn.Query(context.Background(),"SELECT * FROM users")
	if err!= nil{
		return nil, errors.New("failed to get users")
	}
	defer rows.Close()
	users,err:= pgx.CollectRows(rows, pgx.RowToStructByName[User])
	return users,nil
}
func (u * User) CreateUser(conn *pgx.Conn, name string) error{
	_,err:= conn.Exec(context.Background(), "INSERT INTO users(name) VALUES ($1)",name)
	if err!= nil{
		return errors.New("Failed to insert user")
	}
	return nil
}
func (u *User) UpdateUser(conn *pgx.Conn, id int,name string) error{
	_,err:=
}
