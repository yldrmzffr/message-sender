package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"message-sender/config"
	"message-sender/internal/pkg/logger"
)

var (
	cfg config.Config
)

func configSetup() {
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println("Error loading config")
		return
	}
}

func main() {
	configSetup()

	logger.InitLogger(&cfg.Log)

	logger.Info("Starting application...")
}
