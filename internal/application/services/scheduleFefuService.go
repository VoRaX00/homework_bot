package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"homework_bot/internal/domain"
	"homework_bot/pkg/scheduleParser"
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

func generateLink(typeSchedule string, firstDate, secondDate time.Time) string {
	link := fmt.Sprintf("https://univer.dvfu.ru/schedule/get?type=%s&start=%s", typeSchedule, firstDate.Format("2006-01-02"))
	link += "T14%3A00%3A00.000Z&" + fmt.Sprintf("end=%s", secondDate.Format("2006-01-02")) + "T14%3A00%3A00.000Z&groups%5B%5D=5560&ppsId=&facilityId=0"
	return link
}

func (s *ScheduleFefuService) GetOnDate(date time.Time) domain.Schedule {
	link := generateLink("agendaDay", date.Add(-24*time.Hour), date)
	res, err := s.parser.ParseSchedule(link)
	if err != nil {
		return domain.Schedule{}
	}
	return res
}

func getDatesForWeek() (time.Time, time.Time) {
	date := time.Now()
	dayOfWeek := int(date.Weekday())
	lastSunday := date.AddDate(0, 0, -dayOfWeek)
	saturday := date.AddDate(0, 0, 6-dayOfWeek)
	return lastSunday, saturday
}

func (s *ScheduleFefuService) GetOnWeek() domain.Schedule {
	lastSunday, saturday := getDatesForWeek()
	link := generateLink("agendaWeek", lastSunday, saturday)

	res, err := s.parser.ParseSchedule(link)
	if err != nil {
		logrus.Errorf("Error in parse schedule, %v", err)
		return domain.Schedule{}
	}
	return res
}

func (s *ScheduleFefuService) GetOnToday() domain.Schedule {
	return s.GetOnDate(time.Now())
}

func (s *ScheduleFefuService) GetOnTomorrow() domain.Schedule {
	return s.GetOnDate(time.Now().AddDate(0, 0, 1))
}
