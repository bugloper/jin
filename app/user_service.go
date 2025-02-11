package app

import (
	"jin/domain"
	"jin/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user domain.User) (*domain.User, error) {
	return s.repo.CreateUser(user)
}
