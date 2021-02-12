package usersUseCase

import (
	"strings"

	"emperror.dev/errors"

	userEntities "cv-todo-app/internal/user/entities"
	userValidators "cv-todo-app/internal/user/entities/validators"
	"cv-todo-app/pkg/checkers"
	"cv-todo-app/pkg/regexes"
)

type v struct {
}

func (v *v) ValidToCreate(input *userEntities.CreateInput) (err error) {

	if input.Name == "" {
		return errors.Errorf("name not provided")
	}

	if input.Mail == "" {
		return errors.Errorf("mail not provided")
	}

	if input.PhoneNumber == "" {
		return errors.Errorf("phone number not provided")
	}

	if !checkers.ValidMail(input.Mail) {
		return errors.Errorf("invalid email provided: %s", input.Mail)
	}

	if !checkers.ValidPhone(input.PhoneNumber) {
		return errors.Errorf("invalid phone number provided: %s", input.Mail)
	}

	input.Mail = strings.ToLower(input.Mail)
	input.PhoneNumber = regexes.NotNumbers.ReplaceAllLiteralString(input.PhoneNumber, "")
	return nil
}

func (v *v) ValidToUpdate(input *userEntities.UpdateInput) (err error) {
	if input.Id == "" {
		return errors.Errorf("id not provided")
	}

	if input.Name == "" {
		return errors.Errorf("name not provided")
	}

	if input.Mail == "" {
		return errors.Errorf("mail not provided")
	}

	if input.PhoneNumber == "" {
		return errors.Errorf("phone number not provided")
	}

	if !checkers.ValidMail(input.Mail) {
		return errors.Errorf("invalid email provided: %s", input.Mail)
	}

	if !checkers.ValidPhone(input.PhoneNumber) {
		return errors.Errorf("invalid phone number provided: %s", input.Mail)
	}

	input.Mail = strings.ToLower(input.Mail)
	input.PhoneNumber = regexes.NotNumbers.ReplaceAllLiteralString(input.PhoneNumber, "")
	return nil
}

func newValidator() userValidators.UseCaseValidator {
	return &v{}
}
