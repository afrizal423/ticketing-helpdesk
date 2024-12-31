package wa

import (
	"context"
	"database/sql"

	"github.com/go-telegram/bot"
	"github.com/redis/go-redis/v9"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func tiketAwal(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, dariOrang types.JID, rdb *redis.Client) {
	var res string
	res = "Masukkan judul tiket anda.\n\n"
	res += "*_Contoh:_*\n"
	res += "*_- Halaman laporan error_*"
	client.SendMessage(context.Background(), dariOrang, &waProto.Message{
		Conversation: proto.String(res),
	})
}

func tiketDeskripsi(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, dariOrang types.JID, rdb *redis.Client) {
	var res string
	res = "Masukkan isi detail tiket anda.\n"
	res += "Isikan dengan detail dan rinci dari kendala yang anda alami dari penggunaan siskom dari ICT."
	client.SendMessage(context.Background(), dariOrang, &waProto.Message{
		Conversation: proto.String(res),
	})
}

func tiketLampiran(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, dariOrang types.JID, rdb *redis.Client) {
	var res string
	res = "Sisipkan lampiran file pendukung.\n"
	res += "Lampiran file pendukung untuk mengetahui kendala yang dialami oleh anda.\n\n"
	res += "File format yang didukung jpg, png, pdf.\n\n*Lampiran file hanya file yang dilampirkan tanpa ada text caption/deskripsi/chat*.\n\n"
	res += "_Jika tidak ingin melampikan file, harap balas dengan kata *tidak* atau *no*_"
	client.SendMessage(context.Background(), dariOrang, &waProto.Message{
		Conversation: proto.String(res),
	})
}

func akhiriBuatTiket(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, dariOrang types.JID, rdb *redis.Client, notiket string) {
	var resp string
	resp = "*Berhasil membuat tiket*\n"
	resp += "Berikut adalah nomor tiket anda *" + notiket + "*\n\n"
	resp += "Untuk mengetahui status tiket anda, kirimkan perintah `!tiketku <nomor_tiket_anda>`\n"
	resp += "Untuk mengetahui tiket yang anda miliki, kirimkan perintah `!tiketku`\n"
	client.SendMessage(context.Background(), dariOrang, &waProto.Message{
		Conversation: proto.String(resp),
	})
}

func myTiket(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, dariOrang types.JID, rdb *redis.Client, nowa string) {
	listnya := listMyKodeTiket(db, nowa)
	var resp string
	resp = "*Berikut ini adalah tiket yang anda miliki*\n\n"
	resp += listnya
	resp += "\nUntuk mengetahui detail tiket anda, kirimkan perintah `!tiketku <nomor_tiket_anda>`"
	client.SendMessage(context.Background(), dariOrang, &waProto.Message{
		Conversation: proto.String(resp),
	})
}
