package tele

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/afrizal423/ticketing-helpdesk/internal/repository/payload"
	"github.com/afrizal423/ticketing-helpdesk/internal/repository/query"
	"github.com/go-telegram/bot/models"
	"github.com/redis/go-redis/v9"
)

func cekSudahDaftar(db *sql.DB, uid string) int {
	jum := query.CekSudahDaftarTele(db, uid)
	// fmt.Println(jum)
	return jum
}

func simpanDataEmployee(db *sql.DB, datax *models.Update) {
	var input payload.SimpanDataEmployee
	input.Userid = datax.Message.From.ID
	input.Username = datax.Message.Chat.Username
	input.FirstName = datax.Message.From.FirstName
	input.LastName = datax.Message.From.LastName
	input.Chat_id = datax.Message.Chat.ID

	query.SimpanDataEmployee(db, input)
}

func ListTiketAktif(db *sql.DB) (res string) {
	res = query.ListTiketAktif(db)
	return
}

func GrabTiketAktif(db *sql.DB, notiket string) (res string) {
	res = query.GrabTiket(db, notiket)
	return
}

func cekTiketIsOpen(db *sql.DB, notiket string) int {
	jum := query.CekTiketIsOpen(db, notiket)
	// fmt.Println(jum)
	return jum
}

func updateOnChatConversationTiket(db *sql.DB, notiket string, emp string) {
	query.UpdateOnChatConversationTiket(db, notiket, emp)
}

func updateDoneOnChatConversationTiket(db *sql.DB, ctx context.Context, rdb *redis.Client, notiket string, emp string) {
	query.UpdateDoneOnChatConversationTiket(db, notiket, emp)

	hapusSesi(ctx, rdb, emp)
}

func teleCekJikaOnChatDanBlmDone(db *sql.DB, emp string) int {
	jum := query.TeleCekJikaOnChatDanBlmDone(db, emp)
	// fmt.Println(jum)
	return jum
}

func teleGetTiketOnChat(db *sql.DB, emp string) (string, string) {
	tkt, wa := query.TeleGetTiketOnChat(db, emp)
	// fmt.Println(jum)
	return tkt, wa
}

func teleSimpanChatOn(db *sql.DB, arg payload.TeleInsertChat) {
	query.TeleSimpanChatOn(db, arg)
}

func hapusStateDoneTiket(ctx context.Context, rdb *redis.Client, emp string) {
	hapusSesi(ctx, rdb, emp)
}

func setDoneTiket(ctx context.Context, rdb *redis.Client, nowa string) {
	err := rdb.Set(ctx, nowa, "set-done", 1*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func setDoneTiketNomornya(ctx context.Context, rdb *redis.Client, nowa string, tiket string) {
	err := rdb.Set(ctx, nowa+"_done", tiket, 1*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func getTiketDoneNomornya(ctx context.Context, rdb *redis.Client, nowa string) (val string) {
	val, err := rdb.Get(ctx, nowa+"_done").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key tidak ditemukan.")
			val = ""
		} else {
			fmt.Println("Error getting value:", err)
			val = ""
		}
	}
	return
}

func cekPosisiDoneTiket(ctx context.Context, rdb *redis.Client, nowa string) bool {
	val, err := rdb.Get(ctx, nowa).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key tidak ditemukan.")
			return false
		} else {
			fmt.Println("Error getting value:", err)
			return false
		}
	}

	if val == "set-done" {
		return true
	}

	return false
}
