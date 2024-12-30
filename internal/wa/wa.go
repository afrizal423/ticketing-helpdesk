package wa

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-telegram/bot"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"github.com/redis/go-redis/v9"
	"go.mau.fi/whatsmeow"
)

func Mulai(ctx context.Context, db *sql.DB, teleGo *bot.Bot, client *whatsmeow.Client, rdb *redis.Client) error {
	fmt.Println("ini wa")
	client.AddEventHandler(GetEventHandler(ctx, client, teleGo, db, rdb))

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err := client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				// fmt.Println("QR code:", evt.Code)
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err := client.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()

	return nil
}
