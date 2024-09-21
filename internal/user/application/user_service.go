package application

import (
	"context"
	"mygogo/internal/user/domain"
)

type UserService interface {
	GetUserById(id string) (*domain.User, error)
	CreateNewUser(name, email string) (*domain.User, error)
}

type userServiceImpl struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) GetUserById(id string) (*domain.User, error) {
	user, err := s.userRepository.GetUserById(context.Background(), id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) CreateNewUser(name, email string) (*domain.User, error) {
	user := domain.NewUser(name, email)

	err := s.userRepository.CreateUser(context.Background(), user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
