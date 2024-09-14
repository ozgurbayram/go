package application

import (
	"mygogo/internal/user/domain"

	"github.com/google/uuid"
)

type UserService interface {
	GetUserById(id uuid.UUID) *domain.User
	CreateNewUser(name, email string) *domain.User
}

type userServiceImpl struct {
}

func NewUserService() UserService {
	return &userServiceImpl{}
}

func (s *userServiceImpl) GetUserById(id uuid.UUID) *domain.User {
	user := domain.NewUser("ozgur", "dsa")

	return user
}

func (s *userServiceImpl) CreateNewUser(name, email string) *domain.User {
	user := domain.NewUser("ozgur", "dsada")

	return user
}
