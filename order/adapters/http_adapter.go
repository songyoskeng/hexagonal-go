package adapters

import (
	"hexagonalgo/order/core"

	"github.com/gofiber/fiber/v3"
)

type HttpOrderHandler struct {
	service core.OrderService
}

func NewHttpOrderHandler(service core.OrderService) *HttpOrderHandler {
	return &HttpOrderHandler{service: service}
}

func (h *HttpOrderHandler) CreateOrder(c fiber.Ctx) error {
	var order core.Order

	if err := c.Bind().Body(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	if err := h.service.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
