package db

import (
	"database/sql"
	"fmt"
)

type UserStore interface {
	CreateUser(name, email, password string) error
}
type PostGresUserStore struct {
	conn *sql.DB
}

func NewPostGresUserStore(conn *sql.DB) *PostGresUserStore {
	return &PostGresUserStore{
		conn: conn,
	}
}

func (p *PostGresUserStore) CreateUser(name, email, password string) error {
	sql := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);`
	_, err := p.conn.Exec(sql)
	if err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}

	sql = `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err = p.conn.Exec(sql, name, email, password)
	if err != nil {
		return err
	}

	fmt.Println("User created successfully!")
	return nil
}
