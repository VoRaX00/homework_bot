package interfaces

import (
	"main.go/pkg/entity"
)

type IHomeworkRepository interface {
	Create(homework entity.Homework) (int, error)
	GetByTags(tags []string) ([]entity.HomeworkToGet, error)
	GetByName(name string) ([]entity.HomeworkToGet, error)
	GetByWeek() ([]entity.HomeworkToGet, error)
	GetById(id int) (entity.HomeworkToGet, error)
	GetAll() ([]entity.HomeworkToGet, error)
	Update(homeworkToUpdate entity.HomeworkToUpdate) (entity.Homework, error)
}
