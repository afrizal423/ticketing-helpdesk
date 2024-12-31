package tele

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/redis/go-redis/v9"
	"go.mau.fi/whatsmeow"
)

type InitTele struct {
	ClientWA *whatsmeow.Client
	Db       *sql.DB
	Chatid   int64
	Userid   int64
	Token    string
	Rdb      *redis.Client
}

func Mulai(ctx context.Context, db *sql.DB, b *bot.Bot, client *whatsmeow.Client, telex InitTele) error {
	fmt.Println("ini tele")
	// ctx2, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, telex.HelloHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/list_tiket", bot.MatchTypeExact, telex.listTiket)
	b.Start(ctx)
	return nil
}
