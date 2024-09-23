package services

import (
	"homework_bot/internal/application/interfaces"
	"homework_bot/internal/infrastructure/repositories"
)

type Service struct {
	interfaces.IHomeworkService
	interfaces.IScheduleService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IHomeworkService: NewHomeworkService(repos),
		IScheduleService: NewScheduleService(repos),
	}
}
