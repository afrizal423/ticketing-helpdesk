package wa

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	handlerwa "github.com/afrizal423/ticketing-helpdesk/internal/tele/handler_wa"
	"github.com/go-telegram/bot"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

func KirimdariTeleHandler(ctx context.Context, teleGo *bot.Bot, client *whatsmeow.Client, msgTele string) {
	aku, _ := parseJID("NOMORKU")
	pesan := msgTele
	var pointerToPesan *string = &pesan

	msg := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: pointerToPesan,
		},
	}
	// var pesan = "> ⓘ _This number was temporarily banned from WhatsApp for participating in a group of sad single men on Saturday nights. This WhatsApp was confiscated by the Republic of Indonesia Police Institution._"
	client.SendMessage(ctx, aku, msg)
	fmt.Println("masukk")
}

func GetEventHandler(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot) func(interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			if !v.Info.IsFromMe && !v.Info.IsGroup {
				log.Printf("Ada pesan dari %s", v.Info.Sender)
				cekUser, _ := json.Marshal(v.Info.Sender)
				// fmt.Println(string(cekUser))
				var pesan string
				if strings.Contains(string(cekUser), ":") {
					pesan = *v.Message.ExtendedTextMessage.Text
				} else {
					pesan = v.Message.GetConversation()
				}
				log.Printf("pesannyaa %s", pesan)
				// Tentukan timestamp (waktu saat ini)
				// timestamp := time.Now()
				// tandai biru
				// client.MarkRead([]types.MessageID{v.Info.ID}, timestamp, v.Info.Chat, v.Info.Sender)
				jsonData, _ := json.Marshal(v.Info.Sender)
				fmt.Println(string(jsonData))

				handlerwa.KirimdariWA(ctx, teleGo, pesan)

				// telegramBot.SendMessage(ctx, )

				// aku, _ := parseJID("NOMORKU")
				// msg := &waProto.Message{
				// 	Conversation: proto.String(pesan),
				// }
				// var pesan = "> ⓘ _This number was temporarily banned from WhatsApp for participating in a group of sad single men on Saturday nights. This WhatsApp was confiscated by the Republic of Indonesia Police Institution._"
				// client.SendMessage(context.Background(), aku, msg)

				// if strings.Contains(v.Message.GetConversation(), "-ask") {
				// 	gagal, jwban := Chat(apiKey, strings.Replace(v.Message.GetConversation(), "-ask", "", -1))
				// 	if gagal != nil {
				// 		client.SendMessage(context.Background(), v.Info.Sender, &waProto.Message{
				// 			Conversation: proto.String(gagal.Error()),
				// 		})
				// 	} else {
				// 		client.SendMessage(context.Background(), v.Info.Sender, &waProto.Message{
				// 			Conversation: proto.String(jwban),
				// 		})
				// 	}

				// } else {
				// 	client.SendMessage(context.Background(), v.Info.Sender, &waProto.Message{
				// 		Conversation: proto.String("Jangan lupa pakai trigger -ask"),
				// 	})
				// }

			}
		}
	}
}
