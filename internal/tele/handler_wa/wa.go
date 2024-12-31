package handlerwa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

func KirimTeleDokumenDariWA(ctx context.Context, b *bot.Bot, pesan string, kepada string, alamat string, namafile string) {

	fileData, errReadFile := os.ReadFile(alamat)
	if errReadFile != nil {
		fmt.Printf("error read file, %v\n", errReadFile)
		return
	}

	params := &bot.SendDocumentParams{
		ChatID:   kepada,
		Document: &models.InputFileUpload{Filename: namafile, Data: bytes.NewReader(fileData)},
		Caption:  pesan,
	}

	_, err := b.SendDocument(ctx, params)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
