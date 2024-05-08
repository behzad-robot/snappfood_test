package routers

import (
	"snappfood/ordering/services/trips"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewTripRouter(fiberRouter fiber.Router, service trips.TripService) {
	group := fiberRouter.Group("/trips")
	group.Get("/by-order-id/:orderID", func(c *fiber.Ctx) error {
		orderIDStr := c.Params("orderID")
		orderID, e := strconv.Atoi(orderIDStr)
		if e != nil {
			return badParameters(c)
		}
		result, serviceErr := service.GetOneByOrderID(uint(orderID))
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(result)
	})
	group.Post("/", func(c *fiber.Ctx) error {
		command := trips.CreateTripCommand{}
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
	group.Post("/change-status", func(c *fiber.Ctx) error {
		command := trips.ChangeTripStatusCommand{}
		e := c.BodyParser(&command)
		if e != nil {
			return badParameters(c)
		}
		serviceErr := service.ChangeStatus(command)
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(map[string]any{"success": true})
	})
}
