package services

import (
	"homework_bot/internal/domain"
	"homework_bot/pkg/scheduleParser"
	"os"
	"time"
)

type ScheduleFefuService struct {
	parser scheduleParser.IFefuParser
}

func NewScheduleFefuService() *ScheduleFefuService {
	return &ScheduleFefuService{
		parser: scheduleParser.NewFefuParser(),
	}
}

func (s *ScheduleFefuService) GetOnDate(date time.Time) domain.Schedule {
	link := os.Getenv("SCHEDULE_LINK")
	res, err := s.parser.ParseSchedule(link)
	if err != nil {
		return domain.Schedule{}
	}
	return res
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
