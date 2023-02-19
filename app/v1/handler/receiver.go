package handler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/receiver"
	"github.com/lucasszmt/transfeera-challenge/utils"
	"net/http"
)

type ReceiverHandler interface {
	Create() fiber.Handler
	Update() fiber.Handler
	List() fiber.Handler
	Get() fiber.Handler
	Search() fiber.Handler
	Delete() fiber.Handler
}

type receiverHandler struct {
	recvService receiver.UseCase
}

func NewReceiverHandler(useCase receiver.UseCase) ReceiverHandler {
	return &receiverHandler{recvService: useCase}
}

func (r *receiverHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dtos.CreateReceiverRequest{}
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
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
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

func (r *receiverHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dtos.UpdateReceiverRequest{}
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
		err = r.recvService.UpdateReceiver(req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"status": false,
				"errors": fmt.Sprintf("unable to update the receiver requested: %s", err),
			})
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status": true,
			"data":   fmt.Sprintf("user with id %s updated", req.Id),
		})
	}
}

func (r *receiverHandler) List() fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			Page uint
		}{}
		if err := c.QueryParser(&param); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": true,
				"error":  fmt.Sprintf("invalid param for pages, it should be an integer"),
			})
		}
		receivers, err := r.recvService.ListReceivers(int(param.Page))
		if err != nil {
			return err
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":    true,
			"receivers": receivers,
		})
	}
}

func (r *receiverHandler) Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := struct {
			Id string `params:"id"`
		}{}
		if err := c.ParamsParser(&param); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"errors": "invalid data request",
			})
		}
		resp, err := r.recvService.GetReceiver(param.Id)
		if err != nil {
			if errors.Is(err, receiver.ErrReceiverNotFound) {
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"status": false,
					"errors": "receiver not found",
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"status": false,
				"errors": "some unexpected err has happened",
			})
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":    true,
			"receivers": resp,
		})
	}
}

func (r *receiverHandler) Search() fiber.Handler {
	return func(c *fiber.Ctx) error {
		params := dtos.SearchRequest{}
		if err := c.QueryParser(&params); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"error":  "invalid query params",
			})
		}
		if err := utils.ValidateStruct(params); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"error":  "query param required",
			})
		}
		//success case
		recv, err := r.recvService.SearchReceivers(params)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"error":  "receivers not found",
			})
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":    true,
			"receivers": recv,
		})
	}
}

func (r *receiverHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dtos.DeleReceiverRequest{}
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
				"errors": "ids field is required, and it needs to be an array",
			})
		}
		if err := r.recvService.DeleteReceivers(req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"status": false,
				"errors": fmt.Sprintf("an err has happened while deliting the following items %v", req.Ids),
			})
		}
		return c.SendStatus(http.StatusNoContent)
	}
}
