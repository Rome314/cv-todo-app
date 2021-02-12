package config

import "fmt"

type config struct {
	Server   serverConfig   `yaml:"server"`
	Mongo    mongoConfig    `yaml:"mongo"`
	Postgres postgresConfig `yaml:"postgres"`
}

type serverConfig struct {
	ApiPort string `yaml:"apiPort"`
}

type mongoConfig struct {
	Host     string `yaml:"host"`
	Db       string `yaml:"db"`
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
}

func (m mongoConfig) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s/%s", m.Login, m.Password, m.Host, m.Db)
}

type redisConfig struct {
	HostName string `yaml:"hostName"`
	Password string `yaml:"password"`
}

type postgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

func (p postgresConfig) GetConnString() string {
	return fmt.Sprintf("host=%s port=%s database=%s user=%s password=%s", p.Host, p.Port, p.Db, p.User, p.Password)
}
