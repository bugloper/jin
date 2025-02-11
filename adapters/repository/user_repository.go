package repository

import (
	"errors"
	"jin/domain"
	"jin/ports"
)

type InMemoryUserRepository struct {
	users []domain.User
}

func NewInMemoryUserRepository() ports.UserRepository {
	return &InMemoryUserRepository{}
}

func (r *InMemoryUserRepository) GetAllUsers() ([]domain.User, error) {
	return r.users, nil
}

func (r *InMemoryUserRepository) GetUserByID(id int) (*domain.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) CreateUser(user domain.User) (*domain.User, error) {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return &user, nil
}
