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

	err = rdb.Del(ctx, nowa+"*").Err()
	if err != nil {
		fmt.Println("Error deleting value:", err)
		return
	}
}
