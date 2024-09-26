package scheduleParser

import "homework_bot/internal/domain"

type IFefuParser interface {
	ParseSchedule(link string) (domain.Schedule, error)
}
