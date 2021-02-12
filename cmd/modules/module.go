package modules

import "github.com/gin-gonic/gin"

type UseCase interface {
	IsModuleUseCase()
}
type Repository interface {
	IsModuleRepository()
}

type Handler struct {
	Secured bool
	Route   string
	Worker  gin.HandlerFunc
}

func NewHandler(secured bool, route string, worker gin.HandlerFunc) *Handler {
	return &Handler{Secured: secured, Route: route, Worker: worker}
}

type Module interface {
	GetName() string
	GetUC() UseCase
	GetRepo() Repository
	GetGinHandlers() []*Handler
}
