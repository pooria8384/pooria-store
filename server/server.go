package server

import (
	"pooria-store/handlers"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

var serverInstance *Server
var serverOnce sync.Once

func NewServer(userHandler *handlers.User, authHandler *handlers.AuthUser) *Server {
	serverOnce.Do(func() {
		app := fiber.New()
		app.Post("/login", authHandler.Login)

		serverInstance = &Server{
			App: app,
		}
	})
	return serverInstance
}

func (s *Server) Start(port string) error {
	return s.App.Listen(port)
}
