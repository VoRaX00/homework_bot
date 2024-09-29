package services

import (
	"homework_bot/internal/domain"
	"homework_bot/internal/infrastructure/repositories"
)

type UserService struct {
	repos *repositories.Repository
}

func NewUserService(repos *repositories.Repository) *UserService {
	return &UserService{
		repos: repos,
	}
}

func (s *UserService) Create(user domain.User) error {
	return s.repos.IUserRepository.Create(user)
}

func (s *UserService) Update(user domain.User) error {
	return s.repos.IUserRepository.Update(user)
}

func (s *UserService) GetByUsername(username string) (domain.User, error) {
	return s.repos.IUserRepository.GetByUsername(username)
}
