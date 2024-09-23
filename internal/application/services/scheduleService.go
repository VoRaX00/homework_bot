package services

import (
	"homework_bot/internal/domain/models"
	"homework_bot/internal/infrastructure/repositories"
	"time"
)

type ScheduleService struct {
	repo *repositories.Repository
}

func NewScheduleService(repo *repositories.Repository) *ScheduleService {
	return &ScheduleService{
		repo: repo,
	}
}

func (s *ScheduleService) Get(date time.Time) (models.Schedule, error) {
	return s.repo.IScheduleRepository.Get(date)
}

func (s *ScheduleService) GetOnWeek() ([]models.Schedule, error) {
	return s.repo.IScheduleRepository.GetOnWeek()
}

func (s *ScheduleService) Add(date time.Time) error {
	return s.repo.IScheduleRepository.Add(date)
}
