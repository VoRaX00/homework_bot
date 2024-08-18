package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	IHomeworkRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IHomeworkRepository: NewHomeworkRepository(db),
	}
}
