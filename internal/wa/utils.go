package wa

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
	"go.mau.fi/whatsmeow/types"
)

func getNomorHP(arg string) (phoneNumber string) {
	if strings.Contains(string(arg), ":") {
		parts := strings.Split(arg, "@")
		phoneParts := strings.Split(parts[0], ":")
		phoneNumber = phoneParts[0]
	} else {
		parts := strings.Split(arg, "@")
		phoneNumber = parts[0]
	}
	return
}

func parseJID(arg string) (types.JID, bool) {
	if arg[0] == '+' {
		arg = arg[1:]
	}
	if !strings.ContainsRune(arg, '@') {
		return types.NewJID(arg, types.DefaultUserServer), true
	} else {
		recipient, err := types.ParseJID(arg)
		if err != nil {
			// log.Error().Err(err).Msg("Invalid JID")
			panic(err)
			return recipient, false
		} else if recipient.User == "" {
			// log.Error().Err(err).Msg("Invalid JID no server specified")
			panic(err)
			return recipient, false
		}
		return recipient, true
	}
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
