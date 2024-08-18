package service

import "main.go/entity"

type IHomeworkService interface {
	Create(homework entity.Homework) (int, error)
}
