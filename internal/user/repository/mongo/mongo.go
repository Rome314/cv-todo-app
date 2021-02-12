package userMongoRepository

import (
	"go.mongodb.org/mongo-driver/mongo"

	userEntities "cv-todo-app/internal/user/entities"
	userAdapters "cv-todo-app/internal/user/entities/adapters"
)

type repo struct {
	collection *mongo.Collection
	adapter    userAdapters.Bson
}

func (r *repo) IsModuleRepository() {
}

func New(database *mongo.Database) userEntities.Repository {
	return &repo{
		collection: database.Collection("users"),
		adapter:    nil,
	}
}

func (r *repo) Store(input userEntities.User) (user userEntities.User, err error) {
	panic("implement me")
}

func (r *repo) GetOne(id string) (user userEntities.User, err error) {
	panic("implement me")
}

func (r *repo) Update(input userEntities.User) (user userEntities.User, err error) {
	panic("implement me")
}

func (r *repo) DeleteOne(id string) (err error) {
	panic("implement me")
}



