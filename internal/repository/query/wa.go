package query

import (
	"database/sql"
	"log"

	"github.com/afrizal423/ticketing-helpdesk/internal/repository/payload"
)

func CekSudahDaftarWA(db *sql.DB, nowa string) int {
	rows, err := db.Query(`
			SELECT COUNT(1) JUM FROM IHD_CLIENT_WA
			WHERE NO_WA=:1`, nowa)
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

func GetNamaClient(db *sql.DB, nowa string) (*payload.DataClient, error) {
	var user payload.DataClient
	// Use QueryRow to fetch a single row
	err := db.QueryRow(`SELECT NAMA, LOKASI FROM IHD_CLIENT_WA WHERE NO_WA = :1`, nowa).Scan(&user.Nama, &user.Lokasi)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Return the error if something else went wrong
	}
	return &user, nil // Return the user struct
}

func SimpanDataClient(db *sql.DB, arg payload.SimpanDataClient) {
	_, err := db.Exec(`INSERT INTO IHD_CLIENT_WA (NO_WA, NAMA, LOKASI, LOG_TGL) VALUES (:1, :2, :3, SYSDATE)`, arg.Nowa, arg.Nama, arg.Lokasi)
	if err != nil {
		log.Fatal(err)
	}
}
