package domain

import "github.com/google/uuid"

type User struct {
	Id            uuid.UUID `db:"id"`
	Username      string    `db:"name"`
	CodeDirection string    `db:"code_direction"`
	StudyGroup    string    `db:"study_group"`
}
