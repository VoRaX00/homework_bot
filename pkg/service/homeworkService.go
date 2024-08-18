package service

import (
	"main.go/entity"
	"main.go/pkg/repository"
)

type HomeworkService struct {
	repos *repository.Repository
}

func NewHomeworkService(repos *repository.Repository) *HomeworkService {
	return &HomeworkService{
		repos: repos,
	}
}

func (s *HomeworkService) Create(homework entity.Homework) (int, error) {
	return s.repos.Create(homework)
}
