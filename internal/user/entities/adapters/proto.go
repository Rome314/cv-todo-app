package userAdapters

import (
	userEntities "cv-todo-app/internal/user/entities"
	usersRPC "cv-todo-app/rpc/user"
)

type Proto interface {
	CreateInputFromProto(input *usersRPC.CreateRequest) (resp userEntities.CreateInput)
	UpdateInputFromProto(input *usersRPC.UpdateRequest) (resp userEntities.UpdateInput)
	UserToProto(input userEntities.User) (resp *usersRPC.User)
}
