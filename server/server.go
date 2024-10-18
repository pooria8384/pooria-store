package server

import (
	"my-project/handlers"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

func NewServer(userHandler *handlers.User, authHandler *handlers.AuthUser) *Server {
	app := fiber.New()
	app.Post("/login", authHandler.Login)

	return &Server{
		App: app,
	}
}

func (s *Server) Start(port string) error {
	return s.App.Listen(port)
}
