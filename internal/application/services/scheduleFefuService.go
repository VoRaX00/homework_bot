package services

import "homework_bot/internal/domain"

type ScheduleFefuService struct {
}

func NewScheduleFefuService() *ScheduleFefuService {
	return &ScheduleFefuService{}
}

func (s *ScheduleFefuService) GetOnDate() domain.Schedule {
	return domain.Schedule{}
}

func (s *ScheduleFefuService) GetOnWeek() domain.Schedule {
	return domain.Schedule{}
}

func (s *ScheduleFefuService) GetOnToday() domain.Schedule {
	return domain.Schedule{}
}

func (s *ScheduleFefuService) GetOnTomorrow() domain.Schedule {
	return domain.Schedule{}
}
