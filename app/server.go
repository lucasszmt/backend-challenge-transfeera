package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lucasszmt/transfeera-challenge/domain/receiver"
	"os"
)

type Server struct {
	app             *fiber.App
	receiverService receiver.UseCase
}

func NewServer(receiverService receiver.UseCase) *Server {
	server := &Server{
		app:             fiber.New(),
		receiverService: receiverService,
	}
	server.app.Use(logger.New())
	server.router()
	return server
}

func (s *Server) Run() {
	err := s.app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		panic(err)
	}
}

func (s *Server) Close() error {
	return s.app.Shutdown()
}
