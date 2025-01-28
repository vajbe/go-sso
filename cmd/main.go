package main

import (
	"go-sso/internal/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterUserRoutes(router)

	http.ListenAndServe(":8080", router)
}
