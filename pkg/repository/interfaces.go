package repository

import "main.go/entity"

type IHomeworkRepository interface {
	Create(homework entity.Homework) (int, error)
}
