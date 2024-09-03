package interfaces

import (
	"main.go/pkg/entity"
	"time"
)

type IHomeworkService interface {
	Create(homework entity.Homework) (int, error)
	GetByTags(tags []string) ([]entity.HomeworkToGet, error)
	GetById(id int) (entity.HomeworkToGet, error)
	GetAll() ([]entity.HomeworkToGet, error)
	GetByName(name string) ([]entity.HomeworkToGet, error)
	GetByWeek() ([]entity.HomeworkToGet, error)
	GetByToday() ([]entity.HomeworkToGet, error)
	GetByTomorrow() ([]entity.HomeworkToGet, error)
	GetByDate(date time.Time) ([]entity.HomeworkToGet, error)
	Update(homeworkToUpdate entity.HomeworkToUpdate) (entity.Homework, error)
	Delete(id int) error
}
