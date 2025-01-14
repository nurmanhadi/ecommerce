package repository

import "ecommerce/internal/model"

type UserRepository interface {
	RegisterUser(user *model.User) error
	CountUser(email *string) (int, error)
	GetUserById(userId *string) (*model.User, error)
	GetUserByEmail(email *string) (*model.User, error)
}
