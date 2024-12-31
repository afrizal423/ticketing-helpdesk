package tele

import (
	"database/sql"

	"github.com/afrizal423/ticketing-helpdesk/internal/repository/payload"
	"github.com/afrizal423/ticketing-helpdesk/internal/repository/query"
	"github.com/go-telegram/bot/models"
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
