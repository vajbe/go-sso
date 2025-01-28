package handler

import "net/http"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

}
