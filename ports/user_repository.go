package ports

import "jin/domain"

type UserRepository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	CreateUser(user domain.User) (*domain.User, error)
}
