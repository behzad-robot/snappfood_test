package routers

import (
	"snappfood/ordering/services/support"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewSupportRouter(fiberRouter fiber.Router, service support.SupportService) {
	group := fiberRouter.Group("/support/")
	group.Get("/order-status/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, e := strconv.Atoi(idStr)
		if e != nil {
			return badParameters(c)
		}
		result, serviceErr := service.GetOrderDelayStatus(uint(id))
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(result)
	})
	group.Post("/report-order-delay", func(c *fiber.Ctx) error {
		command := support.ReportOrderDelayCommand{}
		e := c.BodyParser(&command)
		if e != nil {
			return badParameters(c)
		}
		result, serviceErr := service.ReportOrderDelay(command)
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(result)
	})
	group.Get("/get-task-for-agent/:agentID", func(c *fiber.Ctx) error {
		idStr := c.Params("agentID")
		id, e := strconv.Atoi(idStr)
		if e != nil {
			return badParameters(c)
		}
		result, serviceErr := service.GetTaskForAgent(uint(id))
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(result)
	})
	group.Post("/update-task-for-agent", func(c *fiber.Ctx) error {
		command := support.UpdateTaskForAgentCommand{}
		e := c.BodyParser(&command)
		if e != nil {
			return badParameters(c)
		}
		serviceErr := service.UpdateTaskForAgent(command)
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(command)
	})
}
