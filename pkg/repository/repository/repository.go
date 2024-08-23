package repository

import (
	"github.com/jmoiron/sqlx"
	"main.go/pkg/service/interfaces"
)

type Repository struct {
	interfaces.IHomeworkRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IHomeworkRepository: NewHomeworkRepository(db),
	}
}
