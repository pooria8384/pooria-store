package handlers

import (
	"fmt"
	"my-project/auth"
	"my-project/storer"
	"my-project/types"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	user storer.User
}

func NewUser(user storer.User) *User {
	return &User{
		user: user,
	}
}

func (s *User) AuthenticateUsers(email, password string) (*types.User, error) {
	user, err := s.user.FindUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %v", err)
	}

	return user, nil
}

func (s *User) GenerateTokens(userID int) (string, error) {
	return auth.GenerateJWT(userID)
}
