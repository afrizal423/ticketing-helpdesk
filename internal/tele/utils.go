package tele

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-telegram/bot"
	"github.com/redis/go-redis/v9"
)

func generateFileName(extension string) string {
	timestamp := time.Now().Unix() // Mendapatkan timestamp dalam detik
	return filepath.Join("storage", "dari_tele", strconv.FormatInt(timestamp, 10)+"."+extension)
}

// escapeSpecialChars menambahkan backslash sebelum karakter khusus
func escapeSpecialChars(input string) string {
	// Daftar karakter yang perlu di-escape
	specialChars := []string{".", "-", "!", "(", ")", "_", "*", "[", "]", "~", "`", ">", "#", "+", "=", "|", "{", "}"}

	// Gantikan setiap karakter khusus dengan versi yang di-escape
	for _, char := range specialChars {
		input = strings.ReplaceAll(input, char, "\\"+char)
	}

	return input
}

func hapusSesi(ctx context.Context, rdb *redis.Client, nowa string) {
	var cursor uint64
	for {
		// Mengambil key
		keys, nextCursor, err := rdb.Scan(ctx, cursor, "*"+nowa+"*", 0).Result()
		if err != nil {
			log.Fatalf("Error scanning keys: %v", err)
		}

		// Menghapus key yang ditemukan
		for _, key := range keys {
			err := rdb.Del(ctx, key).Err()
			if err != nil {
				log.Printf("Error deleting key %s: %v", key, err)
			} else {
				fmt.Printf("Deleted key: %s\n", key)
			}
		}

		// Jika cursor 0, berarti sudah tidak ada lagi key yang tersisa
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	fmt.Println("Selesai menghapus key.")
}

func downloadAndSaveFile(ctx context.Context, b *bot.Bot, fileID string, filename string, token string) {
	file, err := b.GetFile(ctx, &bot.GetFileParams{
		FileID: fileID,
	})
	if err != nil {
		panic(err)
	}

	// Membuat folder jika belum ada
	os.MkdirAll(filepath.Dir(filename), os.ModePerm)

	fmt.Println(file.FilePath)

	// Membuat file di server
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Mengunduh file
	resp, err := http.Get("https://api.telegram.org/file/bot" + token + "/" + file.FilePath)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Menulis file ke server
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
