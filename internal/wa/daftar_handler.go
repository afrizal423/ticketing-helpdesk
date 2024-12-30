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

func daftarAwalClient(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, dariOrang types.JID, rdb *redis.Client) {
	client.SendMessage(context.Background(), dariOrang, &waProto.Message{
		Conversation: proto.String("Siapakah nama anda ?"),
	})
}

func daftarAwalClientLokasi(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, dariOrang types.JID, rdb *redis.Client) {
	client.SendMessage(context.Background(), dariOrang, &waProto.Message{
		Conversation: proto.String("Dari toko manakah anda berasal?"),
	})
}

func akhiriDaftar(ctx context.Context, client *whatsmeow.Client, teleGo *bot.Bot, db *sql.DB, dariOrang types.JID, rdb *redis.Client) {
	var resp string
	resp = "*Terima kasih telah mendaftar kedalam sistem kami*\n"
	resp += "Silahkan untuk membuat tiket dengan perintah !buattiket"
	client.SendMessage(context.Background(), dariOrang, &waProto.Message{
		Conversation: proto.String(resp),
	})
}
