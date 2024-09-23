package repositories

import (
	"github.com/jmoiron/sqlx"
	"homework_bot/internal/application/interfaces"
)

type Repository struct {
	interfaces.IHomeworkRepository
	interfaces.IScheduleRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IHomeworkRepository: NewHomeworkRepository(db),
		IScheduleRepository: NewScheduleRepository(db),
	}
}
