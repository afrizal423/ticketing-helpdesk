package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type ConfigRedis struct {
	Host string
}

func Redis(ctx context.Context, cfg ConfigRedis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host, // alamat Redis server
		Password: "",       // tidak ada password secara default
		DB:       0,        // menggunakan database default
	})

	// Tes ping ke Redis
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error pinging Redis:", err)
		return nil, nil
	}

	fmt.Println("Ping response from Redis:", pong) // Harusnya mencetak "PONG"
	return rdb, nil
}
