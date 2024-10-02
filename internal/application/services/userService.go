package services

import (
	"github.com/google/uuid"
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
	user.Id = uuid.New()
	return s.repos.IUserRepository.Create(user)
}

func (s *UserService) Update(user domain.User) error {
	return s.repos.IUserRepository.Update(user)
}

func (s *UserService) GetByUsername(username string) (domain.User, error) {
	return s.repos.IUserRepository.GetByUsername(username)
}
