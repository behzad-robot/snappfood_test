package routers

import (
	"snappfood/ordering/services/analytics"

	"github.com/gofiber/fiber/v2"
)

func NewAnalyticsRouter(fiberRouter fiber.Router, service analytics.AnalyticsService) {
	group := fiberRouter.Group("/analytics")
	group.Get("/vendor-delays-weekly-report", func(c *fiber.Ctx) error {
		result, serviceErr := service.GetWeeklyVendorDelaysReport()
		if serviceErr != nil {
			return sendServiceError(c, serviceErr)
		}
		return c.JSON(result)
	})
}
