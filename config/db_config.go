package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDefaultDB() *sql.DB {
	LoadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to default database: ", err)
	}

	return db
}

func ConnectToSeedUserDB() *sql.DB {
	// اتصال به دیتابیس `seed-user`
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to seed-user database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB unreachable: ", err)
	}

	fmt.Println("Successfully connected to the seed-user database!")
	return db
}

func CreateDatabaseIfNotExists(db *sql.DB, dbName string) error {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('%s'))", dbName)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if database exists: %v", err)
	}

	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbName))
		if err != nil {
			return fmt.Errorf("error creating database: %v", err)
		}
		fmt.Printf("Database '%s' created successfully!\n", dbName)
	} else {
		fmt.Printf("Database '%s' already exists.\n", dbName)
	}
	return nil
}
