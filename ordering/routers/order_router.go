package routers

import (
	"snappfood/ordering/services/orders"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewOrderRouter(fiberRouter fiber.Router, service orders.OrderService) {
	group := fiberRouter.Group("/orders")
	group.Get("/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, e := strconv.Atoi(idStr)
		if e != nil {
			return badParameters(c)
		}
		result, serviceErr := service.GetOne(uint(id))
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(result)
	})
	group.Post("/", func(c *fiber.Ctx) error {
		command := orders.CreateOrderCommand{}
		e := c.BodyParser(&command)
		if e != nil {
			return badParameters(c)
		}
		result, serviceErr := service.Create(command)
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(result)
	})
}
