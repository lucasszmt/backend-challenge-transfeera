package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/receiver"
	"github.com/lucasszmt/transfeera-challenge/utils"
	"net/http"
)

type ReceiverHandler interface {
	Create() fiber.Handler
}

type receiverHandler struct {
	recvService receiver.UseCase
}

func NewReceiverHandler(useCase receiver.UseCase) ReceiverHandler {
	return &receiverHandler{recvService: useCase}
}

func (r *receiverHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dtos.CreateUserRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"errors": "invalid data request",
			})
		}
		if err := utils.ValidateStruct(req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"errors": fmt.Sprintf("invalid data request: %s", err),
			})
		}
		resp, err := r.recvService.CreateReceiver(req)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"errors": fmt.Sprintf("unable to create receiver cause %s", err),
			})
		}
		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"status": true,
			"data":   fmt.Sprintf("user with id %s created", resp.Id()),
		})
	}
}
