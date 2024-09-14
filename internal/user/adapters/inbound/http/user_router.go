package http

import "github.com/gorilla/mux"

func UserRouter(router *mux.Router, userController *UserController) {
	r := router.PathPrefix("/user").Subrouter()

	r.HandleFunc("", userController.CreateUser).Methods("GET")
}
