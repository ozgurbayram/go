package http

import (
	"encoding/json"
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

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Name and email are required"})
		return
	}

	usr, err := c.userService.CreateNewUser(name, email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]domain.User{"ok": *usr})
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
