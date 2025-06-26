package main

import (
	"github.com/smile-ko/go-ddd-template/config"
	"github.com/smile-ko/go-ddd-template/internal"
)

func main() {
	// Initialize the application configuration
	cfg, err := config.NewConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Run the application with the configuration
	internal.Run(cfg)
}
