package services

import (
	"main.go/pkg/entity"
	"main.go/pkg/repository/repository"
	"time"
)

type HomeworkService struct {
	repos *repository.Repository
}

func NewHomeworkService(repos *repository.Repository) *HomeworkService {
	return &HomeworkService{
		repos: repos,
	}
}

func (s *HomeworkService) Create(homework entity.Homework) (int, error) {
	homework.CreatedAt = time.Now()
	homework.UpdatedAt = time.Now()
	return s.repos.Create(homework)
}

func (s *HomeworkService) GetByTags(tags []string) ([]entity.HomeworkToGet, error) {
	return s.repos.GetByTags(tags)
}

func (s *HomeworkService) GetById(id int) (entity.HomeworkToGet, error) {
	return s.repos.GetById(id)
}

func (s *HomeworkService) GetAll() ([]entity.HomeworkToGet, error) {
	return s.repos.GetAll()
}

func (s *HomeworkService) GetByName(name string) ([]entity.HomeworkToGet, error) {
	return s.repos.GetByName(name)
}

func (s *HomeworkService) GetByWeek() ([]entity.HomeworkToGet, error) {
	return s.repos.GetByWeek()
}

func (s *HomeworkService) Update(homeworkToUpdate entity.HomeworkToUpdate) (entity.Homework, error) {
	return s.repos.Update(homeworkToUpdate)
}
