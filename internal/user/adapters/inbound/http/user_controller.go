package http

import (
	"encoding/json"
	"mygogo/internal/core/utils"
	"mygogo/internal/user/application"
	"mygogo/internal/user/domain"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService application.UserService
}

func NewUserController(userService application.UserService) *UserController {
	return &UserController{userService: userService}
}

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request UserCreateRequest

	if !utils.DecodeAndValidateRequest(r, &request, w) {
		return
	}

	usr, err := c.userService.CreateNewUser(request.Name, request.Email, request.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]domain.User{"ok": *usr})
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userService.GetAllUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string][]*domain.User{"users": users})
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	usr, err := c.userService.GetUserById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]domain.User{"ok": *usr})
}
