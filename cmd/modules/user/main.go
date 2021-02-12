package userModule

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"

	"cv-todo-app/cmd/modules"
	userEntities "cv-todo-app/internal/user/entities"
	userMongoRepository "cv-todo-app/internal/user/repository/mongo"
	userPostgresRepository "cv-todo-app/internal/user/repository/postgres"
	usersUseCase "cv-todo-app/internal/user/useCase"
	usersWeb "cv-todo-app/internal/user/web"
	usersRPC "cv-todo-app/rpc/user"
)

type userModule struct {
	uc       userEntities.UseCase
	repo     userEntities.Repository
	handlers []*modules.Handler
}

func (u *userModule) GetName() string {
	return name
}

func (u *userModule) GetUC() modules.UseCase {

	return u.uc
}

func (u *userModule) GetRepo() modules.Repository {

	return u.repo
}

func (u *userModule) GetGinHandlers() []*modules.Handler {

	return u.handlers
}

const name = "Users"

func GetMongoBased(db *mongo.Database) modules.Module {

	repo := userMongoRepository.New(db)
	uc := usersUseCase.New(repo)

	twirpHandler := usersWeb.GetTwirp()(uc)
	twirpServer := usersRPC.NewUsersServiceServer(twirpHandler, nil)

	ginF := gin.WrapH(twirpServer)

	handler := modules.NewHandler(false, twirpServer.PathPrefix()+"/*w", ginF)

	return &userModule{
		uc:       uc,
		repo:     repo,
		handlers: []*modules.Handler{handler},
	}
}

func GetPgBased(pool *pgxpool.Pool) modules.Module {

	repo := userPostgresRepository.New(pool)
	uc := usersUseCase.New(repo)

	twirpHandler := usersWeb.GetTwirp()(uc)
	twirpServer := usersRPC.NewUsersServiceServer(twirpHandler, nil)

	ginF := gin.WrapH(twirpServer)

	handler := modules.NewHandler(false, twirpServer.PathPrefix()+"/*w", ginF)

	return &userModule{
		uc:       uc,
		repo:     repo,
		handlers: []*modules.Handler{handler},
	}
}
