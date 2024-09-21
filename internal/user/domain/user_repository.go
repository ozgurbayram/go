package domain

import "context"

type UserRepository interface {
	GetUserById(ctx context.Context, id string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}
