package wa

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/afrizal423/ticketing-helpdesk/internal/repository/payload"
	handlerwa "github.com/afrizal423/ticketing-helpdesk/internal/tele/handler_wa"
	"github.com/go-telegram/bot"
	"github.com/redis/go-redis/v9"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func KirimdariTeleHandler(ctx context.Context, teleGo *bot.Bot, client *whatsmeow.Client, msgTele string, kepada string) {
	aku, _ := parseJID(kepada)
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

func KirimdariTeleHandlerSendImage(ctx context.Context, cli *whatsmeow.Client, kepada string, filePath string, caption string) {
	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	ke, _ := parseJID(kepada)

	// Upload the file
	uploaded, err := cli.Upload(context.Background(), data, whatsmeow.MediaImage) // Change MediaImage to MediaDocument for PDFs
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	// Create the message
	msg := &waProto.Message{
		ImageMessage: &waProto.ImageMessage{
			Caption:       proto.String(caption),
			URL:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSHA256: uploaded.FileEncSHA256,
			FileSHA256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
		},
	}

	// Send the message
	resp, err := cli.SendMessage(context.Background(), ke, &waProto.Message{ImageMessage: msg.ImageMessage})
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	} else {
		log.Printf("Message sent (server timestamp: %s)", resp.Timestamp)
	}
}

func KirimdariTeleHandlersendPDF(ctx context.Context, cli *whatsmeow.Client, kepada string, filePath string, caption string) {
	// Read the PDF file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read PDF file: %v", err)
	}

	ke, _ := parseJID(kepada)

	// Upload the PDF file
	uploaded, err := cli.Upload(context.Background(), data, whatsmeow.MediaDocument)
	if err != nil {
		log.Fatalf("Failed to upload PDF file: %v", err)
	}

	// Create the message for the PDF document
	msg := &waProto.Message{
		DocumentMessage: &waProto.DocumentMessage{
			Title:         proto.String(caption), // Title of the document
			URL:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Mimetype:      proto.String("application/pdf"), // Set MIME type for PDF
			FileEncSHA256: uploaded.FileEncSHA256,
			FileSHA256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
		},
	}

	// Send the message
	resp, err := cli.SendMessage(context.Background(), ke, &waProto.Message{DocumentMessage: msg.DocumentMessage})
	if err != nil {
		log.Fatalf("Error sending PDF message: %v", err)
	} else {
		log.Printf("PDF sent successfully (server timestamp: %s)", resp.Timestamp)
	}
}

func GetEventHandler(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, rdb *redis.Client) func(interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			if !v.Info.IsFromMe && !v.Info.IsGroup {
				log.Printf("Ada pesan dari %s", v.Info.Sender)
				cekUser, _ := json.Marshal(v.Info.Sender)
				fmt.Println(strings.ReplaceAll(string(cekUser), `"`, ""))
				var pesan, no_hp_client string
				var dariOrang types.JID
				if v.Info.MediaType != "image" && v.Info.MediaType != "document" {
					if strings.Contains(string(cekUser), ":") {
						pesan = *v.Message.ExtendedTextMessage.Text
					} else {
						pesan = v.Message.GetConversation()
					}
				} else {
					if v.Info.MediaType == "image" {
						pesan = v.Message.ImageMessage.GetCaption()
					} else if v.Info.MediaType == "document" {
						pesan = v.Message.DocumentMessage.GetCaption()
					}

				}

				log.Printf("pesannyaa %s", pesan)
				dariOrang, _ = parseJID(getNomorHP(strings.ReplaceAll(string(cekUser), `"`, "")))
				// Tentukan timestamp (waktu saat ini)
				// timestamp := time.Now()
				// tandai biru
				// client.MarkRead([]types.MessageID{v.Info.ID}, timestamp, v.Info.Chat, v.Info.Sender)
				jsonData, _ := json.Marshal(dariOrang)
				fmt.Println(string(jsonData))
				no_hp_client = getNomorHP(strings.ReplaceAll(string(cekUser), `"`, ""))

				// handlerwa.KirimdariWA(ctx, teleGo, pesan)

				// start disini
				var resp string
				if strings.Contains(pesan, "!start") || strings.Contains(pesan, "!mulai") {
					cek := cekSudahDaftar(db, no_hp_client)
					if cek == 0 {
						resp += "*Mohon maaf*, nomor anda belum terdaftar disistem kami.\n"
						resp += "Silahkan ketik perintah !daftar untuk mendaftarkan diri anda terlebih dahulu"
						client.SendMessage(context.Background(), dariOrang, &waProto.Message{
							Conversation: proto.String(resp),
						})
					} else {
						nama, lokasi := GetIdentitasCLient(db, no_hp_client)
						resp = ""
						resp += fmt.Sprintf("*Hai %s dari toko %s*, Selamat datang di Helpdesk ICT.\n", nama, lokasi)
						// resp += "*Hai nama*, Selamat datang di Helpdesk ICT.\n"
						resp += "Silahkan ketik perintah !buattiket untuk membuat tiket"
						client.SendMessage(context.Background(), dariOrang, &waProto.Message{
							Conversation: proto.String(resp),
						})
					}
				} else if strings.Contains(pesan, "!register") || strings.Contains(pesan, "!daftar") {
					cek := cekSudahDaftar(db, no_hp_client)
					if cek > 0 {
						client.SendMessage(context.Background(), dariOrang, &waProto.Message{
							Conversation: proto.String("*Mohon maaf* anda sudah terdaftar disistem kami.\nSilahkan ketik perintah !buattiket untuk membuat tiket."),
						})
					} else {
						daftarAwalClient(ctx, client, teleGo, db, dariOrang, rdb)
						setDaftarNama(ctx, rdb, no_hp_client)
						pesan = ""
					}
				} else if strings.Contains(pesan, "!buattiket") || strings.Contains(pesan, "!createticket") {
					tiketAwal(ctx, client, teleGo, db, dariOrang, rdb)
					setJudulTiket(ctx, rdb, no_hp_client)
					pesan = ""
				} else if strings.Contains(pesan, "!tiketku") || strings.Contains(pesan, "!myticket") {
					// Memisahkan string berdasarkan spasi
					parts := strings.Fields(pesan)
					if len(parts) > 1 {
						// Mengambil elemen kedua yang merupakan kode
						code := parts[1]
						fmt.Println("Kode:", code)
					} else {
						// fmt.Println("Kode tidak ditemukan")
						myTiket(ctx, client, teleGo, db, dariOrang, rdb, no_hp_client)
					}
					pesan = ""
				} else {

					// else ini mengecek pesan bukan dair perintah
					// pengecekkan bisa waktu daftar maupun ngecek mana masuk chat ticket yang aktif

					if waCekJikaOnChatDanBlmDone(db, no_hp_client) == 1 {
						if v.Info.MediaType == "image" {
							// Download the image
							data, err := client.Download(v.Message.GetImageMessage())
							if err != nil {
								fmt.Println("Failed to download image:", err)
								return
							}

							// Save the image to a file
							path := "storage/dari_wa"
							nama_file := "image_" + no_hp_client + "_" + time.Now().Format("20060102_150405") + ".jpg"
							err = saveFile(path, nama_file, data)
							if err != nil {
								fmt.Println("Failed to save image:", err)
								return
							}

							fmt.Println("Image saved successfully!")

							// kirim tele
							no_tiket, notele := waGetTiketOnChat(db, no_hp_client)
							var insert payload.WaInsertChat
							insert.NoTiket = no_tiket
							insert.Dari = no_hp_client
							insert.Pesan = pesan
							insert.Attch = path + "/" + nama_file
							insert.Kepada = notele
							waSimpanChatOn(db, insert)
							pesan = escapeSpecialChars(pesan)
							handlerwa.KirimTeleDokumenDariWA(ctx, teleGo, pesan, notele, path+"/"+nama_file, nama_file)

						} else if v.Info.MediaType == "document" {
							// Check if the document is a PDF
							doc := v.Message.GetDocumentMessage()
							if doc.GetMimetype() == "application/pdf" {
								// Download the PDF
								data, err := client.Download(doc)
								if err != nil {
									fmt.Println("Failed to download PDF:", err)
									return
								}

								// Save the PDF to a file
								path := "storage/dari_wa"
								nama_file := "document_" + no_hp_client + "_" + time.Now().Format("20060102_150405") + ".pdf"
								err = saveFile(path, nama_file, data)
								if err != nil {
									fmt.Println("Failed to save PDF:", err)
									return
								}

								fmt.Println("PDF saved successfully!")

								// kirim tele
								no_tiket, notele := waGetTiketOnChat(db, no_hp_client)
								var insert payload.WaInsertChat
								insert.NoTiket = no_tiket
								insert.Dari = no_hp_client
								insert.Pesan = pesan
								insert.Attch = path + "/" + nama_file
								insert.Kepada = notele
								waSimpanChatOn(db, insert)
								pesan = escapeSpecialChars(pesan)
								handlerwa.KirimTeleDokumenDariWA(ctx, teleGo, pesan, notele, path+"/"+nama_file, nama_file)
							}
						} else {
							no_tiket, notele := waGetTiketOnChat(db, no_hp_client)
							var insert payload.WaInsertChat
							insert.NoTiket = no_tiket
							insert.Dari = no_hp_client
							insert.Pesan = pesan
							insert.Attch = ""
							insert.Kepada = notele
							waSimpanChatOn(db, insert)
							pesan = escapeSpecialChars(pesan)
							handlerwa.KirimTeledariWA(ctx, teleGo, pesan, notele)
						}
					} else
					// cekan posisi daftar isi nama
					if CekPosisiDaftarNama(ctx, rdb, no_hp_client) {
						if len(pesan) > 2 {
							simpanDaftarNama(ctx, rdb, no_hp_client, pesan)
							daftarAwalClientLokasi(ctx, client, teleGo, db, dariOrang, rdb)
							setLokasiAsal(ctx, rdb, no_hp_client)
							pesan = ""
						}
					} else
					// cekan posisi daftar isi lokasi
					if CekPosisiLokasiAsal(ctx, rdb, no_hp_client) {
						// fmt.Println(pesan)
						if len(pesan) > 2 {
							simpanLokasiAsal(ctx, rdb, no_hp_client, pesan)
							akhiriDaftar(ctx, client, teleGo, db, dariOrang, rdb)
							hapusStateDaftar(db, ctx, rdb, no_hp_client)
							pesan = ""
						}
					} else
					// cekan posisi isi judul tiket
					if CekPosisiJudulTiket(ctx, rdb, no_hp_client) {
						if len(pesan) > 2 {
							simpanJudulTiket(ctx, rdb, no_hp_client, pesan)
							tiketDeskripsi(ctx, client, teleGo, db, dariOrang, rdb)
							setIsiTiket(ctx, rdb, no_hp_client)
							pesan = ""
						}
					} else
					// cekan posisi isi tiket
					if CekPosisiIsiTiket(ctx, rdb, no_hp_client) {
						if len(pesan) > 2 {
							simpanDetailTiket(ctx, rdb, no_hp_client, pesan)
							tiketLampiran(ctx, client, teleGo, db, dariOrang, rdb)
							setAttchTiket(ctx, rdb, no_hp_client)
							pesan = ""
						}
					} else
					// cekan posisi attach tiket
					if CekPosisiAttcTiket(ctx, rdb, no_hp_client) {
						if len(pesan) >= 2 && (pesan == "tidak" || pesan == "no") {
							notiket := simpanTiketTanpaAttch(db, ctx, rdb, no_hp_client)
							akhiriBuatTiket(ctx, client, teleGo, db, dariOrang, rdb, notiket)
							// hapusStateDaftar(db, ctx, rdb, no_hp_client)

							pesan = ""
						} else if len(pesan) >= 2 && (pesan != "tidak" || pesan != "no") {
							client.SendMessage(context.Background(), dariOrang, &waProto.Message{
								Conversation: proto.String("*Mohon maaf* tidak sesuai format !!!."),
							})
							pesan = ""
						} else {
							// disini upload file
						}
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
