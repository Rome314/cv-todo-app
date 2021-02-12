package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config *config

func init() {
	initViperDefaults()
	Config = &config{
		Server: serverConfig{
			ApiPort: viper.GetString("server.api_port"),
		},
		Mongo: mongoConfig{
			Host:     viper.GetString("mongo.host"),
			Db:       viper.GetString("mongo.db"),
			Login:    viper.GetString("mongo.login"),
			Password: viper.GetString("mongo.password"),
		},
		Postgres: postgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetString("postgres.port"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			Db:       viper.GetString("postgres.db"),
		},
	}

}

func initViperDefaults() {

	viper.SetDefault("server.api_port", "1234")

	viper.SetDefault("mongo.host", "localhost")
	viper.SetDefault("mongo.db", "showcase")
	viper.SetDefault("mongo.login", "root")
	viper.SetDefault("mongo.password", "pass12345")

	viper.SetDefault("postgres.host","localhost")
	viper.SetDefault("postgres.port","5432")
	viper.SetDefault("postgres.user","postgres")
	viper.SetDefault("postgres.password","secret")
	viper.SetDefault("postgres.db","todo")

	viper.SetDefault("app.debug", true)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AddConfigPath(".")

	if e := viper.ReadInConfig(); e != nil {
		log.Info("Config file not found, using env.")
		viper.AutomaticEnv()

		if viper.GetBool("app.debug") {
			ee := viper.WriteConfigAs("./config.toml")
			if ee != nil {
				log.Error(ee)
			} else {
				log.Info("Debug mode on, thus config file created")
			}
		}

	} else {
		log.Info("Config file used: ", viper.ConfigFileUsed())
	}

}
