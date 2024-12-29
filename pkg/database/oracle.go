package database

import (
	"database/sql"
	"fmt"
	"log"

	go_ora "github.com/sijms/go-ora/v2"
)

type Config struct {
	SID      string
	Username string
	Password string
}

func Konek(cfg Config) (*sql.DB, error) {
	connStr := cfg.SID
	databaseUrl := go_ora.BuildJDBC(cfg.Username, cfg.Password, connStr, map[string]string{})
	fmt.Println(databaseUrl)
	db, err := sql.Open("oracle", databaseUrl)

	if err != nil {
		log.Fatal(err)
	}
	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected to Oracle database")
	return db, nil
}
