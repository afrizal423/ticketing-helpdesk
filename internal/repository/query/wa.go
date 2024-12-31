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

func GenerateKodeTiket(db *sql.DB, nowa string) string {
	var nomor string
	// Use QueryRow to fetch a single row
	q := `
	SELECT 
		TRIM(SUBSTR(NOD, 1, 5)) 
		|| TRIM(TO_CHAR(TO_NUMBER(SUBSTR(NOD, 6)) + 1, '0000000')) NO_LANJUT
	FROM 
	(
		SELECT NVL(MAX(NO_TIKET), 'IHD' ||TO_CHAR(SYSDATE,'YY')|| '0000000') NOD
		FROM IHD_TIKET
		WHERE NO_TIKET LIKE 'IHD' || TO_CHAR(SYSDATE,'YY')|| '%'
	) AP
	`
	err := db.QueryRow(q).Scan(&nomor)
	if err != nil {
		if err == sql.ErrNoRows {
			return "" // No user found
		}
		return "" // Return the error if something else went wrong
	}
	return nomor // Return the user struct
}

func SimpanTiketClient(db *sql.DB, arg payload.SimpanTiketClient) {
	_, err := db.Exec(`INSERT INTO IHD_TIKET (NO_TIKET, NO_WA_CLIENT, JUDUL, ISI, LOG_TGL) 
			VALUES (:1, :2, :3, :4, SYSDATE)`, arg.NoTiket, arg.Nowa, arg.Judul, arg.Isi)
	if err != nil {
		log.Fatal(err)
	}
}

func ListMyKodeTiket(db *sql.DB, nowa string) string {
	rows, err := db.Query(`SELECT NO_TIKET,
			CASE WHEN IS_DONE = 'T' THEN 'OPEN' ELSE 'CLOSE' END STATUS
		FROM IHD_TIKET 
		WHERE NO_WA_CLIENT=:1
		ORDER BY LOG_TGL`, nowa)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var res string
	for rows.Next() {
		var datanya payload.ListTiketHeader
		if err := rows.Scan(&datanya.NoTiket, &datanya.Status); err != nil {
			log.Fatal(err)
		}
		// users = append(users, user)
		res += "- *" + datanya.NoTiket + "* status: *" + datanya.Status + "*\n"
	}

	return res
}

func WACekJikaOnChatDanBlmDone(db *sql.DB, emp string) int {
	rows, err := db.Query(`
			SELECT COUNT(1) JUM FROM IHD_TIKET
			WHERE IS_DONE='T' AND ON_CHAT='Y' AND NO_WA_CLIENT=:1
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

func WaGetTiketOnChat(db *sql.DB, emp string) (string, string) {
	rows, err := db.Query(`
			SELECT NO_TIKET JUM, EMPLOYEE TELE FROM IHD_TIKET
			WHERE IS_DONE='T' AND ON_CHAT='Y' AND NO_WA_CLIENT=:1
			`, emp)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var JUM, TELE string
	for rows.Next() {
		err := rows.Scan(&JUM, &TELE)
		if err != nil {
			log.Fatal(err)
		}
	}
	return JUM, TELE
}

func WaSimpanChatOn(db *sql.DB, arg payload.WaInsertChat) {
	_, err := db.Exec(`INSERT INTO IHD_TIKET_DISCUSS (NO_TIKET, DARI, PESAN, IS_CLIENT, LOG_TGL, KEPADA, ATTCHMENT) VALUES (:1, :2, :3, 'Y', SYSDATE, :4, :5)`, arg.NoTiket, arg.Dari, arg.Pesan, arg.Kepada, arg.Attch)
	if err != nil {
		log.Fatal(err)
	}
}
