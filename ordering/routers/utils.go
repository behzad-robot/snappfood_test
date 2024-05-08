package routers

import (
	"snappfood/ordering/common"

	"github.com/gofiber/fiber/v2"
)

func badParameters(c *fiber.Ctx) error {
	return sendServiceError(c, common.BadParameters)
}
func sendServiceError(c *fiber.Ctx, err *common.ServiceError) error {
	return c.Status(err.Code).JSON(map[string]any{"error": err.Error.Error(), "statusCode": err.Code})
}
