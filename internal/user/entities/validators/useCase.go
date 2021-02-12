package userValidators

import userEntities "cv-todo-app/internal/user/entities"

type UseCaseValidator interface {
	ValidToCreate(input *userEntities.CreateInput) (err error)
	ValidToUpdate(input *userEntities.UpdateInput) (err error)
}
