package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router, userController *UserController) {
	r := router.PathPrefix("/user").Subrouter()

	r.HandleFunc("/{id}", userController.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("", userController.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("", userController.CreateUser).Methods(http.MethodPost)
}
