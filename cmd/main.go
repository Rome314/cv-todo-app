package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"cv-todo-app/cmd/config"
	"cv-todo-app/cmd/connections"
	"cv-todo-app/cmd/logger"
	"cv-todo-app/cmd/modules"
	userModule "cv-todo-app/cmd/modules/user"
)

func checker(step string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", step, err.Error())
	}
}

func init() {

	formatter := &log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	}

	log.SetFormatter(formatter)

	if sentryDSN := viper.GetString("SENTRY_DSN"); sentryDSN != "" {
		e := sentry.Init(sentry.ClientOptions{
			Dsn: sentryDSN,
		})
		checker("sentry.init", e)
		log.Infof("sentry inizialized")
	}

	if viper.GetBool("app.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

}

var (
	g errgroup.Group
)

func main() {

	cfg := config.Config

	// mongoDb, err := connections.GetMongoDatabase(cfg.Mongo.GetURI(), cfg.Mongo.Db)
	// checker("mongo.connect", err)

	postgres, err := connections.GetPostgresDatabase(cfg.Postgres.GetConnString())
	checker("postgres.connect", err)

	uM := userModule.GetPgBased(postgres)

	modulesList := []modules.Module{uM}

	log.Infof("Configuring server...")

	router := gin.Default()
	router.Use(logger.GinLogger(log.StandardLogger()))

	for _, module := range modulesList {
		log.Infof("[%s]\n", module.GetName())
		for _, handler := range module.GetGinHandlers() {
			if !handler.Secured {
				router.POST(handler.Route, handler.Worker)
				log.Infof("-- Gin ✅ \n")

			}
		}

		log.Info("Running...")

	}
	err = router.Run(":" + cfg.Server.ApiPort)
	checker("server.running", err)
}
