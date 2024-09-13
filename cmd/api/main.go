package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"message-sender/config"
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

	fmt.Println("Hello, World!")
}
