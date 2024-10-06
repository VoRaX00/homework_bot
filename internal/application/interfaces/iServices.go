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
	GetOnDate(user domain.User, date time.Time) domain.Schedule
	GetOnWeek(user domain.User) domain.Schedule
	GetOnNextWeek(user domain.User) domain.Schedule
	GetOnToday(user domain.User) domain.Schedule
	GetOnTomorrow(user domain.User) domain.Schedule
}

type IUserService interface {
	Create(user domain.User) error
	Update(user domain.User) error
	GetByUsername(username string) (domain.User, error)
}
