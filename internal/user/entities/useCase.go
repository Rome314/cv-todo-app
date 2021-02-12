package userEntities

import "cv-todo-app/cmd/modules"

type UseCase interface {
	modules.UseCase
	Create(input CreateInput) (user User, err error)
	Update(input UpdateInput) (user User, err error)
	GetOne(id string) (user User, err error)
	DeleteOne(id string) (err error)
}
