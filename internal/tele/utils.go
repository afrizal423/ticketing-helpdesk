package tele

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-telegram/bot"
)

func generateFileName(extension string) string {
	timestamp := time.Now().Unix() // Mendapatkan timestamp dalam detik
	return filepath.Join("storage", "dari_tele", strconv.FormatInt(timestamp, 10)+"."+extension)
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
