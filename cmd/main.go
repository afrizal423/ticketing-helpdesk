package main

import (
	"log"

	"github.com/afrizal423/ticketing-helpdesk/pkg/config"
	"github.com/afrizal423/ticketing-helpdesk/pkg/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// dbnya
	db, err := database.Konek(database.Config(cfg.Database))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

}
