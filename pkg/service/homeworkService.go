package service

import (
	"main.go/entity"
	"main.go/pkg/repository"
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
	return s.repos.Create(homework)
}

func (s *HomeworkService) GetByTags(tags []string) ([]entity.Homework, error) {
	return s.repos.GetByTags(tags)
}

func (s *HomeworkService) GetById(id int) (entity.Homework, error) {
	return s.repos.GetById(id)
}

func (s *HomeworkService) GetAll() ([]entity.Homework, error) {
	return s.repos.GetAll()
}

func (s *HomeworkService) GetByName(name string) ([]entity.Homework, error) {
	return s.repos.GetByName(name)
}

func (s *HomeworkService) GetByWeek() ([]entity.Homework, error) {
	return s.repos.GetByWeek()
}

func (s *HomeworkService) Update(homeworkToUpdate entity.HomeworkToUpdate) (entity.Homework, error) {
	return s.repos.Update(homeworkToUpdate)
}
