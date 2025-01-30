package handler

import (
	"encoding/json"
	"fmt"
	"go-sso/internal/db"
	res "go-sso/internal/middleware"
	"go-sso/internal/types"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := db.AddUser(newUser)
	if err != nil {
		res.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been added successfully.", resp)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin types.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		res.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := db.Login(userLogin)
	if err != nil {
		res.Error(w, fmt.Sprintf("failed to login: %s ", err.Error()), http.StatusInternalServerError)
		return
	}
	res.Success(w, "User has been logged in successfully.", resp)
}
