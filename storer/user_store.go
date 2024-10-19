package storer

import (
	"database/sql"
	"fmt"
	"pooria-store/types"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	CreateUser(name, email, password string) error
	FindUserByEmail(email string) (*types.User, error)
}
type PostGresUser struct {
	conn *sql.DB
}

var userInstance *PostGresUser
var userOnce sync.Once

func NewPostGresUserStore(conn *sql.DB) *PostGresUser {
	userOnce.Do(func() {
		userInstance = &PostGresUser{
			conn: conn,
		}
	})
	return userInstance
}

func (p *PostGresUser) CreateUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}
	sql := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	);`
	_, err = p.conn.Exec(sql)
	if err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}

	sql = `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err = p.conn.Exec(sql, name, email, hashedPassword)
	if err != nil {
		return err
	}

	fmt.Println("User created successfully!")
	return nil
}

func (p *PostGresUser) FindUserByEmail(email string) (*types.User, error) {
	user := &types.User{}
	query := "SELECT id, name, email, password FROM users WHERE email=$1"
	err := p.conn.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
