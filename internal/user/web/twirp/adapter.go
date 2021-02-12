package userTwirpDelivery

import (
	userEntities "cv-todo-app/internal/user/entities"
	userAdapters "cv-todo-app/internal/user/entities/adapters"
	usersRPC "cv-todo-app/rpc/user"
)

type a struct {
}

func (a *a) CreateInputFromProto(input *usersRPC.CreateRequest) (resp userEntities.CreateInput) {
	return userEntities.CreateInput{
		PhoneNumber: input.PhoneNumber,
		Mail:        input.Email,
		Name:        input.Name,
	}
}

func (a *a) UpdateInputFromProto(input *usersRPC.UpdateRequest) (resp userEntities.UpdateInput) {
	return userEntities.UpdateInput{
		Id:          input.Id,
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Mail:        input.Email,
	}
}

func (a *a) UserToProto(input userEntities.User) (resp *usersRPC.User) {
	return &usersRPC.User{
		Id:          input.Id,
		Name:        input.Name,
		Email:       input.Mail,
		PhoneNumber: input.PhoneNumber,
		Created:     input.Created.Unix(),
		LastUpdated: input.LastUpdated.Unix(),
	}
}

func newAdapter() userAdapters.Proto {
	return &a{}
}
