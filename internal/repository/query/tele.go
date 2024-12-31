package query

import (
	"database/sql"
	"log"

	"github.com/afrizal423/ticketing-helpdesk/internal/repository/payload"
	"github.com/go-telegram/bot"
)

func CekSudahDaftarTele(db *sql.DB, uid string) int {
	rows, err := db.Query(`
			SELECT COUNT(1) JUM FROM IHD_EMPLOYEE
			WHERE USER_ID=:1`, uid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var JUM int
	for rows.Next() {
		err := rows.Scan(&JUM)
		if err != nil {
			log.Fatal(err)
		}
	}
	return JUM
}

func SimpanDataEmployee(db *sql.DB, arg payload.SimpanDataEmployee) {
	_, err := db.Exec(`INSERT INTO IHD_EMPLOYEE (USER_ID, USERNAME, FIRSTNAME, LASTNAME, CHAT_ID, LOG_TGL) VALUES (:1, :2, :3, :4, :5, SYSDATE)`, arg.Userid, arg.Username, arg.FirstName, arg.LastName, arg.Chat_id)
	if err != nil {
		log.Fatal(err)
	}
}

func ListTiketAktif(db *sql.DB) string {
	rows, err := db.Query(`SELECT NO_TIKET,
			NO_WA_CLIENT NOWA
		FROM IHD_TIKET 
		WHERE IS_DONE='T'
		ORDER BY LOG_TGL`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var res string
	for rows.Next() {
		var datanya payload.ListTiketAktif
		if err := rows.Scan(&datanya.NoTiket, &datanya.Nowa); err != nil {
			log.Fatal(err)
		}
		// users = append(users, user)
		res += "- *" + bot.EscapeMarkdown(datanya.NoTiket) + "* Usere: *" + datanya.Nowa + "*\n"
	}

	return res
}

func GrabTiket(db *sql.DB, notiket string) string {
	rows, err := db.Query(`SELECT NO_TIKET, JUDUL, ISI,
			NO_WA_CLIENT NOWA
		FROM IHD_TIKET 
		WHERE IS_DONE='T'
		AND NO_TIKET=:1`, notiket)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var res string
	for rows.Next() {
		var datanya payload.GrabTiketAktif
		if err := rows.Scan(&datanya.NoTiket, &datanya.Judul, &datanya.Isi, &datanya.Nowa); err != nil {
			log.Fatal(err)
		}
		// users = append(users, user)
		res += "Judul Tiket:\n" + datanya.Judul + "\n\n"
		res += "Isi Tiket:\n" + datanya.Isi + "\n\n"
	}

	return res
}

func CekTiketIsOpen(db *sql.DB, uid string) int {
	rows, err := db.Query(`
			SELECT COUNT(1) JUM FROM IHD_TIKET
			WHERE IS_DONE='T'
		AND NO_TIKET=:1`, uid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var JUM int
	for rows.Next() {
		err := rows.Scan(&JUM)
		if err != nil {
			log.Fatal(err)
		}
	}
	return JUM
}

func UpdateOnChatConversationTiket(db *sql.DB, tiket string, emp string) {
	_, err := db.Exec(`UPDATE IHD_TIKET SET START_CONVERSATION=SYSDATE, ON_CHAT='Y', EMPLOYEE=:1 WHERE NO_TIKET=:2`, emp, tiket)
	if err != nil {
		log.Fatal(err)
	}
}

func TeleCekJikaOnChatDanBlmDone(db *sql.DB, emp string) int {
	rows, err := db.Query(`
			SELECT COUNT(1) JUM FROM IHD_TIKET
			WHERE IS_DONE='T' AND ON_CHAT='Y' AND EMPLOYEE=:1
			`, emp)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var JUM int
	for rows.Next() {
		err := rows.Scan(&JUM)
		if err != nil {
			log.Fatal(err)
		}
	}
	return JUM
}

func TeleGetTiketOnChat(db *sql.DB, emp string) (string, string) {
	rows, err := db.Query(`
			SELECT NO_TIKET JUM, NO_WA_CLIENT NOWA FROM IHD_TIKET
			WHERE IS_DONE='T' AND ON_CHAT='Y' AND EMPLOYEE=:1
			`, emp)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var JUM, NOWA string
	for rows.Next() {
		err := rows.Scan(&JUM, &NOWA)
		if err != nil {
			log.Fatal(err)
		}
	}
	return JUM, NOWA
}

func TeleSimpanChatOn(db *sql.DB, arg payload.TeleInsertChat) {
	_, err := db.Exec(`INSERT INTO IHD_TIKET_DISCUSS (NO_TIKET, DARI, PESAN, IS_EMPLOYEE, LOG_TGL, KEPADA, ATTCHMENT) VALUES (:1, :2, :3, 'Y', SYSDATE, :4, :5)`, arg.NoTiket, arg.Dari, arg.Pesan, arg.Kepada, arg.Attch)
	if err != nil {
		log.Fatal(err)
	}
}
