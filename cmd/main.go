package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/afrizal423/ticketing-helpdesk/internal/wa"
	"github.com/afrizal423/ticketing-helpdesk/pkg/config"
	"github.com/afrizal423/ticketing-helpdesk/pkg/database"
)

func main() {
	appContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// dbnya
	db, err := database.Konek(appContext, database.Config(cfg.Database))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Create a channel to listen for interrupt signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)

	// wa
	// Start WhatsApp bot
	go func() {
		if err := wa.Mulai(appContext, db); err != nil {
			log.Fatalf("WhatsApp bot failed: %v", err)
		}
	}()

	// Block main goroutine
	// select {}
	log.Println("Tekan Ctrl+C untuk exit...")

	// Wait for interrupt signal
	<-interruptChan

	log.Println("Shutting down sukses...")
	// Perform any cleanup here if necessary
	cancel() // Cancel the context to signal goroutines to stop
}
