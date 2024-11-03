package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type AuthUser struct {
	AuthUser *User
}

func NewAuthUser(authUser *User) *AuthUser {
	return &AuthUser{
		AuthUser: authUser,
	}
}

func (h *AuthUser) Login(c *fiber.Ctx) error {
	data := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	user, err := h.AuthUser.AuthenticateUsers(data.Email, data.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	token, err := h.AuthUser.GenerateTokens(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token, "user": user})
}

func Register(c *fiber.Ctx) error {
	return c.JSON("user register seccusfully")
}
