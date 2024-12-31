package handlerwa

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func KirimTeledariWA(ctx context.Context, b *bot.Bot, pesan string, kepada string) {
	jsonData, _ := json.Marshal(&bot.SendMessageParams{
		ChatID:    kepada,
		Text:      pesan,
		ParseMode: models.ParseModeMarkdown,
	})
	fmt.Println(string(jsonData))

	fmt.Println(pesan)

	_, err := b.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID:    kepada,
		Text:      pesan,
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
