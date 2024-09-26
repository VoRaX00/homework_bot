package interfaces

import (
	"homework_bot/internal/domain"
	"time"
)

type IHomeworkService interface {
	Create(homework domain.Homework) (int, error)
	GetByTags(tags []string) ([]domain.HomeworkToGet, error)
	GetById(id int) (domain.HomeworkToGet, error)
	GetAll() ([]domain.HomeworkToGet, error)
	GetByName(name string) ([]domain.HomeworkToGet, error)
	GetByWeek() ([]domain.HomeworkToGet, error)
	GetByToday() ([]domain.HomeworkToGet, error)
	GetByTomorrow() ([]domain.HomeworkToGet, error)
	GetByDate(date time.Time) ([]domain.HomeworkToGet, error)
	Update(homeworkToUpdate domain.HomeworkToUpdate) (domain.Homework, error)
	Delete(id int) error
}

type IScheduleService interface {
	GetOnDate() domain.Schedule
	GetOnWeek() domain.Schedule
	GetOnToday() domain.Schedule
	GetOnTomorrow() domain.Schedule
}
