package wa

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/afrizal423/ticketing-helpdesk/internal/repository/payload"
	"github.com/afrizal423/ticketing-helpdesk/internal/repository/query"
	"github.com/redis/go-redis/v9"
)

func cekSudahDaftar(db *sql.DB, nowa string) int {
	jum := query.CekSudahDaftarWA(db, nowa)
	// fmt.Println(jum)
	return jum
}

func GetIdentitasCLient(db *sql.DB, nowa string) (nama string, lokasi string) {
	data, err := query.GetNamaClient(db, nowa)
	if err != nil {
		fmt.Println("Error setting value:", err)
	}
	nama = data.Nama
	lokasi = data.Lokasi
	return
}

func simpanDaftarNama(ctx context.Context, rdb *redis.Client, nowa string, isi string) {
	err := rdb.Set(ctx, nowa+"_nama", isi, 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func setDaftarNama(ctx context.Context, rdb *redis.Client, nowa string) {
	err := rdb.Set(ctx, nowa, "set-nama", 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func CekPosisiDaftarNama(ctx context.Context, rdb *redis.Client, nowa string) bool {
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

	if val == "set-nama" {
		return true
	}

	return false
}

func simpanLokasiAsal(ctx context.Context, rdb *redis.Client, nowa string, isi string) {
	err := rdb.Set(ctx, nowa+"_lokasi", isi, 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func setLokasiAsal(ctx context.Context, rdb *redis.Client, nowa string) {
	err := rdb.Set(ctx, nowa, "set-lokasi", 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func CekPosisiLokasiAsal(ctx context.Context, rdb *redis.Client, nowa string) bool {
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

	if val == "set-lokasi" {
		return true
	}

	return false
}

func hapusStateDaftar(db *sql.DB, ctx context.Context, rdb *redis.Client, nowa string) {
	// insert data dahulu
	val, err := rdb.Get(ctx, nowa+"_nama").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key tidak ditemukan.")
		} else {
			fmt.Println("Error getting value:", err)
		}
	}
	nama := val

	val, err = rdb.Get(ctx, nowa+"_lokasi").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key tidak ditemukan.")
		} else {
			fmt.Println("Error getting value:", err)
		}
	}
	lokasi := val

	// fmt.Println(nama + lokasi)
	var res payload.SimpanDataClient
	res.Nowa = nowa
	res.Nama = nama
	res.Lokasi = lokasi

	query.SimpanDataClient(db, res)

	hapusSesi(ctx, rdb, nowa)
}

func setJudulTiket(ctx context.Context, rdb *redis.Client, nowa string) {
	err := rdb.Set(ctx, nowa, "set-judul", 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func CekPosisiJudulTiket(ctx context.Context, rdb *redis.Client, nowa string) bool {
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

	if val == "set-judul" {
		return true
	}
	return false
}

func simpanJudulTiket(ctx context.Context, rdb *redis.Client, nowa string, isi string) {
	err := rdb.Set(ctx, nowa+"_judul", isi, 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func setIsiTiket(ctx context.Context, rdb *redis.Client, nowa string) {
	err := rdb.Set(ctx, nowa, "set-isi", 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func CekPosisiIsiTiket(ctx context.Context, rdb *redis.Client, nowa string) bool {
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

	if val == "set-isi" {
		return true
	}
	return false
}

func simpanDetailTiket(ctx context.Context, rdb *redis.Client, nowa string, isi string) {
	err := rdb.Set(ctx, nowa+"_detail", isi, 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func setAttchTiket(ctx context.Context, rdb *redis.Client, nowa string) {
	err := rdb.Set(ctx, nowa, "set-attc", 5*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return
	}
}

func CekPosisiAttcTiket(ctx context.Context, rdb *redis.Client, nowa string) bool {
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

	if val == "set-attc" {
		return true
	}
	return false
}

func simpanTiketTanpaAttch(db *sql.DB, ctx context.Context, rdb *redis.Client, nowa string) (nomor string) {
	// insert data dahulu
	val, err := rdb.Get(ctx, nowa+"_judul").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key tidak ditemukan.")
		} else {
			fmt.Println("Error getting value:", err)
		}
	}
	judul := val

	val, err = rdb.Get(ctx, nowa+"_detail").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("Key tidak ditemukan.")
		} else {
			fmt.Println("Error getting value:", err)
		}
	}
	isi := val

	nomor = query.GenerateKodeTiket(db, nowa)

	// fmt.Println(nama + lokasi)
	var res payload.SimpanTiketClient
	res.Nowa = nowa
	res.NoTiket = nomor
	res.Judul = judul
	res.Isi = isi

	query.SimpanTiketClient(db, res)

	hapusSesi(ctx, rdb, nowa)

	return
}

func listMyKodeTiket(db *sql.DB, nowa string) (res string) {
	res = query.ListMyKodeTiket(db, nowa)
	return
}

func waCekJikaOnChatDanBlmDone(db *sql.DB, emp string) int {
	tkt := query.WACekJikaOnChatDanBlmDone(db, emp)
	// fmt.Println(jum)
	return tkt
}

func waGetTiketOnChat(db *sql.DB, emp string) (string, string) {
	tkt, wa := query.WaGetTiketOnChat(db, emp)
	// fmt.Println(jum)
	return tkt, wa
}

func waSimpanChatOn(db *sql.DB, arg payload.WaInsertChat) {
	query.WaSimpanChatOn(db, arg)
}
