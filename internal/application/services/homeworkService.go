package services

import (
	"homework_bot/internal/domain/models"
	"homework_bot/internal/infrastructure/repositories"
	"time"
)

type HomeworkService struct {
	repos *repositories.Repository
}

func NewHomeworkService(repos *repositories.Repository) *HomeworkService {
	return &HomeworkService{
		repos: repos,
	}
}

func (s *HomeworkService) Create(homework models.Homework) (int, error) {
	homework.CreatedAt = time.Now()
	homework.UpdatedAt = time.Now()
	return s.repos.Create(homework)
}

func (s *HomeworkService) GetByTags(tags []string) ([]models.HomeworkToGet, error) {
	return s.repos.GetByTags(tags)
}

func (s *HomeworkService) GetById(id int) (models.HomeworkToGet, error) {
	return s.repos.GetById(id)
}

func (s *HomeworkService) GetAll() ([]models.HomeworkToGet, error) {
	return s.repos.GetAll()
}

func (s *HomeworkService) GetByName(name string) ([]models.HomeworkToGet, error) {
	return s.repos.GetByName(name)
}

func (s *HomeworkService) GetByWeek() ([]models.HomeworkToGet, error) {
	return s.repos.GetByWeek()
}

func (s *HomeworkService) GetByToday() ([]models.HomeworkToGet, error) {
	return s.repos.GetByToday()
}

func (s *HomeworkService) GetByTomorrow() ([]models.HomeworkToGet, error) {
	return s.repos.GetByTomorrow()
}

func (s *HomeworkService) GetByDate(date time.Time) ([]models.HomeworkToGet, error) {
	return s.repos.GetByDate(date)
}

func (s *HomeworkService) Update(homeworkToUpdate models.HomeworkToUpdate) (models.Homework, error) {
	return s.repos.Update(homeworkToUpdate)
}

func (s *HomeworkService) Delete(id int) error {
	return s.repos.Delete(id)
}
