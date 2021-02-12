package userMocks

import userEntities "cv-todo-app/internal/user/entities"

type UseCase struct {
	CreateFn    func(input userEntities.CreateInput) (user userEntities.User, err error)
	UpdateFn    func(input userEntities.UpdateInput) (user userEntities.User, err error)
	GetOneFn    func(id string) (user userEntities.User, err error)
	DeleteOneFn func(id string) (err error)
}

func (u UseCase) Create(input userEntities.CreateInput) (user userEntities.User, err error) {
	return u.CreateFn(input)
}

func (u UseCase) Update(input userEntities.UpdateInput) (user userEntities.User, err error) {
	return u.UpdateFn(input)
}

func (u UseCase) GetOne(id string) (user userEntities.User, err error) {
	return u.GetOneFn(id)
}

func (u UseCase) DeleteOne(id string) (err error) {
	return u.DeleteOneFn(id)
}
