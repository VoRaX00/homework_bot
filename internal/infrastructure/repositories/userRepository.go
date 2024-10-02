package repositories

import (
	"github.com/jmoiron/sqlx"
	"homework_bot/internal/domain"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user domain.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO users (id, username, code_direction, study_group) VALUES ($1, $2, $3, $4)"
	row := tx.QueryRow(query, user.Id, user.Username, user.CodeDirection, user.StudyGroup)
	if row.Err() != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *UserRepository) Update(user domain.User) error {
	query := "UPDATE users SET code_direction = $1, study_group = $2 WHERE username = $3"
	_, err := r.db.Exec(query, user.CodeDirection, user.StudyGroup, user.Username)
	return err
}

func (r *UserRepository) GetByUsername(username string) (domain.User, error) {
	query := "SELECT * FROM users WHERE username = $1"
	var user domain.User
	err := r.db.Get(&user, query, username)
	return user, err
}
