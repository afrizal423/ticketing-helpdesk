package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/afrizal423/ticketing-helpdesk/internal/tele"
	"github.com/afrizal423/ticketing-helpdesk/internal/wa"
	"github.com/afrizal423/ticketing-helpdesk/pkg/config"
	"github.com/afrizal423/ticketing-helpdesk/pkg/database"
	"github.com/go-telegram/bot"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	appContext, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
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

	// redis
	rdb, err := database.Redis(appContext, database.ConfigRedis(cfg.Redis))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Create a channel to listen for interrupt signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)

	// init wa
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	clientWA := whatsmeow.NewClient(deviceStore, clientLog)

	// init tele
	telex := &tele.InitTele{
		ClientWA: clientWA,
	}
	opts := []bot.Option{
		bot.WithDefaultHandler(telex.DefaultHandler),
	}
	teleGo, err := bot.New(cfg.Telegram.Token, opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	// wa
	// Start WhatsApp bot
	go func() {
		if err := wa.Mulai(appContext, db, teleGo, clientWA, rdb); err != nil {
			log.Fatalf("WhatsApp bot failed: %v", err)
		}
	}()

	//tele
	// Start Telegram bot
	go func() {
		if err := tele.Mulai(appContext, db, teleGo, clientWA); err != nil {
			log.Fatalf("Telegram bot failed: %v", err)
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
