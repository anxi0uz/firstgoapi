package handlers

import (
	"firstapi/internal/storage"
	"net/http"
)

type UserHandler struct {
	store *storage.UserStorage
}

func NewUserHandler(st *storage.TodoStorage) *UserHandler {
	return &UserHandler{store: (*storage.UserStorage)(st)}
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

}
