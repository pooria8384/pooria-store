package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var Db *sql.DB
var once sync.Once

func ConnectDB() *sql.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"))

		db, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Fatal("Failed to connect to database: ", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal("DB unreachable: ", err)
		}

		fmt.Println("Successfully connected to database!")
		Db = db
	})

	return Db
}
