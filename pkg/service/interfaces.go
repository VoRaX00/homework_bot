package service

import "main.go/Entity"

type IHomeworkService interface {
	Create(homework Entity.Homework) (int, error)
}
