package api

import (
	"fmt"
	"my-project/config"
	"my-project/db"
	"my-project/types"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	store db.UserStore
}

func NewUserHandler(store db.UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (s *UserHandler) AuthenticateUser(email, password string) (*types.User, error) {
	user, err := s.store.FindUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %v", err)
	}

	return user, nil
}

func (s *UserHandler) GenerateToken(userID int) (string, error) {
	return config.GenerateJWT(userID)
}
