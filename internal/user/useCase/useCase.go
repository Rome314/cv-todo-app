package usersUseCase

import (
	"time"

	"emperror.dev/errors"

	userEntities "cv-todo-app/internal/user/entities"
	userValidators "cv-todo-app/internal/user/entities/validators"
)

type uc struct {
	repo      userEntities.Repository
	validator userValidators.UseCaseValidator
}

func (u *uc) IsModuleUseCase() {
}

func (u *uc) Create(input userEntities.CreateInput) (user userEntities.User, err error) {
	user = userEntities.User{}
	if err = u.validator.ValidToCreate(&input); err != nil {
		err = errors.WithDetails(userEntities.BadRequestError, err)
		return
	}

	toCreate := userEntities.User{
		Name:        input.Name,
		Mail:        input.Mail,
		PhoneNumber: input.PhoneNumber,
		Created:     time.Now(),
		LastUpdated: time.Now(),
	}

	user, err = u.repo.Store(toCreate)
	if err != nil {
		return
	}
	return user, nil

}

func (u *uc) Update(input userEntities.UpdateInput) (user userEntities.User, err error) {
	user = userEntities.User{}
	if err = u.validator.ValidToUpdate(&input); err != nil {
		err = errors.WithDetails(userEntities.BadRequestError, err)
		return
	}

	toUpdate := userEntities.User{
		Id:          input.Id,
		Name:        input.Name,
		Mail:        input.Mail,
		PhoneNumber: input.PhoneNumber,
		LastUpdated: time.Now(),
	}

	user, err = u.repo.Update(toUpdate)
	if err != nil {
		return
	}

	return user, nil
}

func (u *uc) GetOne(id string) (user userEntities.User, err error) {
	if id == "" {
		err = errors.WithDetails(userEntities.BadRequestError, "id not provided")
		return
	}

	user, err = u.repo.GetOne(id)
	if err != nil {
		return
	}
	return user, nil
}

func (u *uc) DeleteOne(id string) (err error) {
	if id == "" {
		err = errors.WithDetails(userEntities.BadRequestError, "id not provided")
		return
	}

	return u.repo.DeleteOne(id)
}

func New(repo userEntities.Repository) userEntities.UseCase {
	return &uc{
		repo:      repo,
		validator: newValidator(),
	}
}
