package service

import "main.go/pkg/repository"

type Service struct {
	IHomeworkService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		IHomeworkService: NewHomeworkService(repos),
	}
}
