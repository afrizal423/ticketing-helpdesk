package query

import (
	"database/sql"
	"log"

	"github.com/afrizal423/ticketing-helpdesk/internal/repository/payload"
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
