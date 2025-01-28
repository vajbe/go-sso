package handler

import (
	"log"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	log.Print("In Add user")
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

}
