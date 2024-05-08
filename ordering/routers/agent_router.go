package routers

import (
	"snappfood/ordering/services/agents"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewAgentRouter(fiberRouter fiber.Router, service agents.AgentService) {
	group := fiberRouter.Group("/agents")
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
		command := agents.CreateAgentCommand{}
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
