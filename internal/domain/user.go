package domain

import "github.com/google/uuid"

type User struct {
	Id            uuid.UUID `db:"id"`
	Username      string    `db:"username"`
	CodeDirection string    `db:"code_direction"`
	StudyGroup    int       `db:"study_group"`
}

func NewUser(username, codeDirection string, studyGroup int) *User {
	return &User{
		Username:      username,
		CodeDirection: codeDirection,
		StudyGroup:    studyGroup,
	}
}
