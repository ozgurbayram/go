package http

import (
	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router, userController *UserController) {
	r := router.PathPrefix("/user").Subrouter()

	r.HandleFunc("/{id}", userController.GetUserById).Methods("GET")
	r.HandleFunc("", userController.CreateUser).Methods("POST")
}
