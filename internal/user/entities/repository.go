package userEntities

import "cv-todo-app/cmd/modules"

type Repository interface {
	modules.Repository
	Store(input User) (user User, err error)
	GetOne(id string) (user User, err error)
	Update(input User) (user User, err error)
	DeleteOne(id string) (err error)
}
