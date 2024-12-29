package wa

import (
	"context"
	"log"
	"time"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func GetEventHandler(client *whatsmeow.Client) func(interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			if !v.Info.IsFromMe && !v.Info.IsGroup {
				log.Printf("Ada pesan dari %s", v.Info.Sender)
				// Tentukan timestamp (waktu saat ini)
				timestamp := time.Now()
				// tandai biru
				client.MarkRead([]types.MessageID{v.Info.ID}, timestamp, v.Info.Chat, v.Info.Sender)

				var pesan = "> ⓘ _This number was temporarily banned from WhatsApp for participating in a group of sad single men on Saturday nights. This WhatsApp was confiscated by the Republic of Indonesia Police Institution._"
				client.SendMessage(context.Background(), v.Info.Sender, &waProto.Message{
					Conversation: proto.String(pesan),
				})
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
