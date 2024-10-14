package api

import "github.com/gofiber/fiber/v2"

type AuthHandler struct {
	UserHandler *UserHandler
}

func NewAuthHandler(userHandler *UserHandler) *AuthHandler {
	return &AuthHandler{
		UserHandler: userHandler,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	data := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	user, err := h.UserHandler.AuthenticateUser(data.Email, data.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	token, err := h.UserHandler.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token, "user": user})
}
