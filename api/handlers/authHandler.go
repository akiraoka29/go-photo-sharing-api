package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	db *database.MySQLDB
}

func NewUserHandler(db *database.MySQLDB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login API")
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Register API")
}