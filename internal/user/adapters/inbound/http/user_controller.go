package http

import (
	"encoding/json"
	"mygogo/internal/user/application"
	"mygogo/internal/user/domain"
	"net/http"
)

type UserController struct {
	userService application.UserService
}

func NewUserController(userService application.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	usr := c.userService.CreateNewUser("ozgur", "bayram")

	json.NewEncoder(w).Encode(map[string]domain.User{"ok": *usr})
}
