package routes

import (
	"go-sso/internal/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	userHandler := handler.NewUserHandler()

	routes := []struct {
		method      string
		path        string
		handlerFunc func(http.ResponseWriter, *http.Request)
	}{
		{"POST", "/register", userHandler.AddUser},
		{"GET", "/oauth/callback", userHandler.AuthCallbackHandler},
		{"GET", "/login", userHandler.SSOHandler},
		/* {"POST", "/login", userHandler.Login}, */

	}

	for _, route := range routes {
		router.HandleFunc(route.path, route.handlerFunc).Methods(route.method)
	}
}
