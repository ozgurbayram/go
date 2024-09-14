package main

import (
	"log"
	userHttp "mygogo/internal/user/adapters/inbound/http"
	userApp "mygogo/internal/user/application"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	userService := userApp.NewUserService()
	userController := userHttp.NewUserController(userService)

	router := mux.NewRouter()

	userHttp.UserRouter(router, userController)

	log.Println("Server is running on port 8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
