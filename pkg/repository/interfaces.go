package repository

import "main.go/Entity"

type IHomeworkRepository interface {
	Create(homework Entity.Homework) (int, error)
}
