package config

import "fmt"

type Config struct {
	Service      ServiceConfig
	Log          LogConfig
	Database     DatabaseConfig
	Redis        RedisConfig
	Notification NotificationConfig
	Messages     MessagesConfig
}

type ServiceConfig struct {
	ENV  string `split_words:"true" required:"true" default:"dev"`
	Name string `split_words:"true" required:"true" default:"boilerplate"`
	Port string `split_words:"true" required:"true" default:"8080"`
}

type LogConfig struct {
	Level string `split_words:"true" default:"DEBUG"`
}

type DatabaseConfig struct {
	Host     string `split_words:"true" required:"true"`
	Port     int    `split_words:"true" required:"true"`
	User     string `split_words:"true" required:"true"`
	Password string `split_words:"true" required:"true"`
	Database string `split_words:"true" required:"true"`
}

type NotificationConfig struct {
	Provider string `split_words:"true" required:"true" default:"mock"`
}

type MessagesConfig struct {
	BatchSize int  `split_words:"true" required:"true" default:"2"`
	AutoStart bool `split_words:"true" required:"true" default:"true"`
	Interval  int  `split_words:"true" required:"true" default:"120"` // seconds
}

type RedisConfig struct {
	Url string `split_words:"true" required:"true"`
}

func (d DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.Database)
}

func (d DatabaseConfig) GetURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		d.User, d.Password, d.Host, d.Port, d.Database)
}
