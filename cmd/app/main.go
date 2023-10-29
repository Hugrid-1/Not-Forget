package main

import (
	"log"

	"github.com/Hugrid-1/Not-Forget/config"
	"github.com/Hugrid-1/Not-Forget/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
