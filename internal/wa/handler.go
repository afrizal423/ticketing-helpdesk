package wa

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/redis/go-redis/v9"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
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

func GetEventHandler(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, rdb *redis.Client) func(interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			if !v.Info.IsFromMe && !v.Info.IsGroup {
				log.Printf("Ada pesan dari %s", v.Info.Sender)
				cekUser, _ := json.Marshal(v.Info.Sender)
				fmt.Println(strings.ReplaceAll(string(cekUser), `"`, ""))
				var pesan string
				var dariOrang types.JID
				if strings.Contains(string(cekUser), ":") {
					pesan = *v.Message.ExtendedTextMessage.Text
				} else {
					pesan = v.Message.GetConversation()
				}
				log.Printf("pesannyaa %s", pesan)
				dariOrang, _ = parseJID(getNomorHP(strings.ReplaceAll(string(cekUser), `"`, "")))
				// Tentukan timestamp (waktu saat ini)
				// timestamp := time.Now()
				// tandai biru
				// client.MarkRead([]types.MessageID{v.Info.ID}, timestamp, v.Info.Chat, v.Info.Sender)
				jsonData, _ := json.Marshal(dariOrang)
				fmt.Println(string(jsonData))

				// handlerwa.KirimdariWA(ctx, teleGo, pesan)

				// start disini
				var resp string
				if strings.Contains(pesan, "!start") || strings.Contains(pesan, "!mulai") {
					cek := cekSudahDaftar(db, getNomorHP(strings.ReplaceAll(string(cekUser), `"`, "")))
					if cek == 0 {
						resp += "*Mohon maaf*, nomor anda belum terdaftar disistem kami.\n"
						resp += "Silahkan ketik perintah !daftar untuk mendaftarkan diri anda terlebih dahulu"
						client.SendMessage(context.Background(), dariOrang, &waProto.Message{
							Conversation: proto.String(resp),
						})
					} else {
						resp = ""
						resp += "*Hai nama*, Selamat datang di Helpdesk ICT.\n"
						resp += "Silahkan ketik perintah !buattiket untuk membuat tiket"
						client.SendMessage(context.Background(), dariOrang, &waProto.Message{
							Conversation: proto.String(resp),
						})
					}
				}

				if strings.Contains(pesan, "!register") || strings.Contains(pesan, "!daftar") {
					daftarAwalClient(ctx, client, teleGo, db, dariOrang, rdb)
					setDaftarNama(ctx, rdb, getNomorHP(strings.ReplaceAll(string(cekUser), `"`, "")))
					pesan = ""
				}

				if CekPosisiDaftarNama(ctx, rdb, getNomorHP(strings.ReplaceAll(string(cekUser), `"`, ""))) {
					if len(pesan) > 2 {
						simpanDaftarNama(ctx, rdb, getNomorHP(strings.ReplaceAll(string(cekUser), `"`, "")), pesan)
						daftarAwalClientLokasi(ctx, client, teleGo, db, dariOrang, rdb)
						setLokasiAsal(ctx, rdb, getNomorHP(strings.ReplaceAll(string(cekUser), `"`, "")))
						pesan = ""
					}
				}

				if CekPosisiLokasiAsal(ctx, rdb, getNomorHP(strings.ReplaceAll(string(cekUser), `"`, ""))) {
					// fmt.Println(pesan)
					if len(pesan) > 2 {
						simpanLokasiAsal(ctx, rdb, getNomorHP(strings.ReplaceAll(string(cekUser), `"`, "")), pesan)
						akhiriDaftar(ctx, client, teleGo, db, dariOrang, rdb)
						hapusStateDaftar(db, ctx, rdb, getNomorHP(strings.ReplaceAll(string(cekUser), `"`, "")))
						pesan = ""
					}
				}

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
