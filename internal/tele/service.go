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
