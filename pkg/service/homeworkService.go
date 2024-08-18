package service

import (
	"main.go/Entity"
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

func (s *HomeworkService) Create(homework Entity.Homework) (int, error) {
	return s.repos.Create(homework)
}
