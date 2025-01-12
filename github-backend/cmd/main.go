package main

import (
	"log"

	"github.com/MarlomSouza/go-git/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	log.Printf("Server running on port %s", cfg.Port)

}
