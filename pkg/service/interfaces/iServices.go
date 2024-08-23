package interfaces

import (
	"main.go/pkg/entity"
)

type IHomeworkService interface {
	Create(homework entity.Homework) (int, error)
	GetByTags(tags []string) ([]entity.Homework, error)
	GetById(id int) (entity.Homework, error)
	GetAll() ([]entity.Homework, error)
	GetByName(name string) ([]entity.Homework, error)
	GetByWeek() ([]entity.Homework, error)
	Update(homeworkToUpdate entity.HomeworkToUpdate) (entity.Homework, error)
}
