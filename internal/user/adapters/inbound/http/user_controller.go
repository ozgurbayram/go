package http

import (
	"log"
	"mygogo/internal/user/application"
	"net/http"
)

type UserController struct {
	userService application.UserService
}

func NewUserController(userService application.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}
