package tele

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/afrizal423/ticketing-helpdesk/internal/wa"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func KirimdariWA(ctx context.Context, b *bot.Bot, pesan string) {
	jsonData, _ := json.Marshal(&bot.SendMessageParams{
		ChatID:    "576507972",
		Text:      pesan,
		ParseMode: models.ParseModeMarkdown,
	})
	fmt.Println(string(jsonData))

	fmt.Println(pesan)

	_, err := b.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID:    "576507972",
		Text:      pesan,
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func (app *InitTele) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: update.Message.Chat.ID,
	// 	Text:   "Say /hello",
	// })
	// fmt.Println(update.Message.Text)
	wa.KirimdariTeleHandler(ctx, b, app.ClientWA, update.Message.Text)
}

func HelloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	jsonData, _ := json.Marshal(update)
	fmt.Println(string(jsonData))

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Hello, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "* " + strconv.FormatInt(update.Message.From.ID, 10),
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
