package usersWeb

import (
	userEntities "cv-todo-app/internal/user/entities"
	userTwirpDelivery "cv-todo-app/internal/user/web/twirp"
	usersRPC "cv-todo-app/rpc/user"
)

func GetTwirp() func(uc userEntities.UseCase) usersRPC.UsersService {
	return userTwirpDelivery.New
}
