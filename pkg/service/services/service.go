package services

import (
	"main.go/pkg/repository/repository"
	"main.go/pkg/service/interfaces"
)

type Service struct {
	interfaces.IHomeworkService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		IHomeworkService: NewHomeworkService(repos),
	}
}
