package tele

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/afrizal423/ticketing-helpdesk/internal/repository/payload"
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
	app.Chatid = update.Message.Chat.ID
	app.Userid = update.Message.From.ID
	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID: update.Message.Chat.ID,
	// 	Text:   "Say /hello",
	// })
	// fmt.Println(update.Message.Text)

	if teleCekJikaOnChatDanBlmDone(app.Db, strconv.FormatInt(update.Message.From.ID, 10)) == 1 {
		app.on_chat(ctx, b, update, strconv.FormatInt(update.Message.From.ID, 10))
	} else if strings.Contains(update.Message.Text, "/grab_tiket") && teleCekJikaOnChatDanBlmDone(app.Db, strconv.FormatInt(update.Message.From.ID, 10)) == 0 {
		parts := strings.Fields(update.Message.Text)
		if len(parts) > 1 {
			// Mengambil elemen kedua yang merupakan kode
			code := parts[1]
			app.grab_tiket(ctx, b, update, code)
			// fmt.Println("Kode:", code)
		} else {
			fmt.Println("Kode tidak ditemukan")
		}
	} else {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "Anda tidak memiliki tiket aktif.",
			ParseMode: models.ParseModeMarkdown,
		})
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
	// wa.KirimdariTeleHandler(ctx, b, app.ClientWA, update.Message.Text)
}

func (app *InitTele) on_chat(ctx context.Context, b *bot.Bot, update *models.Update, emp string) {
	no_tiket, nowa := teleGetTiketOnChat(app.Db, emp)
	var insert payload.TeleInsertChat
	insert.NoTiket = no_tiket
	insert.Dari = emp
	insert.Pesan = update.Message.Text
	insert.Attch = ""
	insert.Kepada = nowa
	teleSimpanChatOn(app.Db, insert)
	fmt.Println(no_tiket + " " + emp + " " + update.Message.Text)
	wa.KirimdariTeleHandler(ctx, b, app.ClientWA, update.Message.Text, nowa)
}

func (app *InitTele) grab_tiket(ctx context.Context, b *bot.Bot, update *models.Update, tiket string) {
	if cekTiketIsOpen(app.Db, tiket) == 0 {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Maaf tiket tidak ada atau sudah closed.",
		})
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	} else {
		var res string
		res = "Berikut adalah tiket dari nomor " + tiket + ".\n\n"
		res += GrabTiketAktif(app.Db, tiket)
		res += "Seluruh percakapan akan terekam oleh sistem.\nPergunakanlah dengan bijak dan sopan.\nSemangat bekerja."
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   res,
		})
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
		updateOnChatConversationTiket(app.Db, tiket, strconv.FormatInt(update.Message.From.ID, 10))
	}

}

func (app *InitTele) HelloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// jsonData, _ := json.Marshal(update)
	// fmt.Println(string(jsonData))
	jum := cekSudahDaftar(app.Db, strconv.FormatInt(update.Message.From.ID, 10))
	if jum == 0 {
		simpanDataEmployee(app.Db, update)
	}

	app.Chatid = update.Message.Chat.ID
	app.Userid = update.Message.From.ID

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Hello, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "* " + strconv.FormatInt(update.Message.From.ID, 10),
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func (app *InitTele) listTiket(ctx context.Context, b *bot.Bot, update *models.Update) {
	// jsonData, _ := json.Marshal(update)
	// fmt.Println(string(jsonData))
	var res string
	res = "Tiket yang belum di respon:\n\n"
	res += ListTiketAktif(app.Db) + "\n"

	// Menambahkan backslash sebelum karakter '-' dan '*'
	res = strings.ReplaceAll(res, "-", "\\-")

	// fmt.Println(res)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      res,
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
