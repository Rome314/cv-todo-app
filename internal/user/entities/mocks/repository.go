package userMocks

import userEntities "cv-todo-app/internal/user/entities"

type Repository struct {
	StoreFn     func(input userEntities.User) (user userEntities.User, err error)
	GetOneFn    func(id string) (user userEntities.User, err error)
	UpdateFn    func(input userEntities.User) (user userEntities.User, err error)
	DeleteOneFn func(id string) (err error)
}

func( Repository)IsModuleRepository(){}

func (r Repository) Store(input userEntities.User) (user userEntities.User, err error) {
	return r.StoreFn(input)
}

func (r Repository) GetOne(id string) (user userEntities.User, err error) {
	return r.GetOneFn(id)
}

func (r Repository) Update(input userEntities.User) (user userEntities.User, err error) {
	return r.UpdateFn(input)
}

func (r Repository) DeleteOne(id string) (err error) {
	return r.DeleteOneFn(id)
}
