package config

import "fmt"

type Config struct {
	Service  ServiceConfig
	Log      LogConfig
	Database DatabaseConfig
}

type ServiceConfig struct {
	ENV  string `split_words:"true" required:"true" default:"dev"`
	Name string `split_words:"true" required:"true" default:"boilerplate"`
	Port string `split_words:"true" required:"true" default:"8080"`
}

type LogConfig struct {
	Level string `split_words:"true" default:"INFO"`
}

type DatabaseConfig struct {
	Host     string `split_words:"true" required:"true"`
	Port     int    `split_words:"true" required:"true"`
	User     string `split_words:"true" required:"true"`
	Password string `split_words:"true" required:"true"`
	Database string `split_words:"true" required:"true"`
}

func (d DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.Database)
}
