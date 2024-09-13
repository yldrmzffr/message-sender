package config

type Config struct {
	Service ServiceConfig
}

type ServiceConfig struct {
	ENV  string `split_words:"true" required:"true" default:"dev"`
	Name string `split_words:"true" required:"true" default:"boilerplate"`
	Port string `split_words:"true" required:"true" default:"8080"`
}

type LogConfig struct {
	Level string `split_words:"true" default:"INFO"`
}
