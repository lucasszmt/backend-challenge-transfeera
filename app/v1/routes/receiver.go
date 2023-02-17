package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasszmt/transfeera-challenge/app/v1/handler"
)

const (
	receiverV1Route = "api/v1/receiver"
)

func ReceiverRoutes(route *fiber.App, handler handler.ReceiverHandler) {
	receiverRoutes := route.Group(receiverV1Route)
	receiverRoutes.Post("/", handler.Create())
	receiverRoutes.Patch("/", handler.Update())
	receiverRoutes.Get("/", handler.List())
	receiverRoutes.Get("/:id", handler.Get())
	receiverRoutes.Delete("/", handler.Delete())
}
