package tele

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-telegram/bot"
	"go.mau.fi/whatsmeow"
)

type InitTele struct {
	ClientWA *whatsmeow.Client
}

func Mulai(ctx context.Context, db *sql.DB, b *bot.Bot, client *whatsmeow.Client) error {
	fmt.Println("ini tele")
	// ctx2, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()

	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, HelloHandler)
	b.Start(ctx)
	return nil
}
