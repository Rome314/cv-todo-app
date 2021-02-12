package userAdapters

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	userEntities "cv-todo-app/internal/user/entities"
)

type Bson interface {
	FromBson(input []byte) (user userEntities.User, err error)
	ToUpdate(input userEntities.User) (update primitive.M)
}
