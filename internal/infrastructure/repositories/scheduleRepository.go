package repositories

import (
	"github.com/jmoiron/sqlx"
	"homework_bot/internal/domain/models"
	"time"
)

type ScheduleRepository struct {
	db *sqlx.DB
}

func NewScheduleRepository(db *sqlx.DB) *ScheduleRepository {
	return &ScheduleRepository{
		db: db,
	}
}

func (repo *ScheduleRepository) Get(date time.Time) (models.Schedule, error) {
	return models.Schedule{}, nil
}

func (repo *ScheduleRepository) GetOnWeek() ([]models.Schedule, error) {
	return nil, nil
}

func (repo *ScheduleRepository) Add(date time.Time) error {
	return nil
}
