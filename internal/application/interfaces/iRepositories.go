package interfaces

import (
	"homework_bot/internal/domain/models"
	"time"
)

type IHomeworkRepository interface {
	Create(homework models.Homework) (int, error)
	GetByTags(tags []string) ([]models.HomeworkToGet, error)
	GetByName(name string) ([]models.HomeworkToGet, error)
	GetByWeek() ([]models.HomeworkToGet, error)
	GetById(id int) (models.HomeworkToGet, error)
	GetAll() ([]models.HomeworkToGet, error)
	GetByToday() ([]models.HomeworkToGet, error)
	GetByTomorrow() ([]models.HomeworkToGet, error)
	GetByDate(date time.Time) ([]models.HomeworkToGet, error)
	Update(homeworkToUpdate models.HomeworkToUpdate) (models.Homework, error)
	Delete(id int) error
}

type IScheduleRepository interface {
	Get(date time.Time) (models.Schedule, error)
	GetOnWeek() ([]models.Schedule, error)
	Add(date time.Time) error
}
