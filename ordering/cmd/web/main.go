package main

import (
	"fmt"
	"log"
	"snappfood/ordering/env"
	"snappfood/ordering/infrastructure/database"
	"snappfood/ordering/internal/repositories"
	"snappfood/ordering/routers"
	"snappfood/ordering/services/agents"
	"snappfood/ordering/services/analytics"
	"snappfood/ordering/services/orders"
	"snappfood/ordering/services/rabbit"
	"snappfood/ordering/services/support"
	"snappfood/ordering/services/trips"
	"snappfood/ordering/services/vendors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/timeout"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("SnappFood Ordering Web")
	env, err := env.CreateEnv()
	if err != nil {
		panic(err)
	}
	db, err := database.CreateGormConnection(env.PostgresqlDatabase)
	if err != nil {
		panic(err)
	}
	rabbitMQConnection, err := amqp.Dial(env.RabbitMQConnection)
	if err != nil {
		panic(err)
	}
	defer rabbitMQConnection.Close()
	//DI:
	vendorsRepo := repositories.NewVendorRepository(db)
	ordersRepo := repositories.NewOrderRepository(db)
	tripsRepo := repositories.NewTripRepository(db)
	agentsRepo := repositories.NewAgentRepository(db)
	delayReportsRepo := repositories.NewDelayReportRepository(db)
	delayResultsRepo := repositories.NewDelayResultRepository(db)

	rabbitHelperService := rabbit.NewRabbitHelperService(rabbitMQConnection)
	defer rabbitHelperService.Close()

	analyticsService := analytics.NewAnalyticsService(delayResultsRepo)
	vendorsService := vendors.NewVendorService(vendorsRepo)
	orderService := orders.NewOrderService(ordersRepo)
	tripService := trips.NewTripService(tripsRepo, analyticsService, orderService)

	agentsService := agents.NewAgentService(agentsRepo)
	supportService := support.NewSupportService(delayReportsRepo, agentsService, orderService, tripService, rabbitHelperService)

	//setup fiber:
	fiberApp := fiber.New()
	fiberApp.Use(cors.New(cors.Config{
		// AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowHeaders:     "*",
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	fiberApp.Use(logger.New(logger.Config{
		Format: "[${ip}] ${status} - ${method} ${path} - ${latency}\n",
	}))
	fiberApp.Use(timeout.NewWithContext(func(c *fiber.Ctx) (err error) { return c.Next() }, 15*time.Second))
	fiberApp.Use(func(c *fiber.Ctx) error {
		c.Response().Header.Add("Access-Control-Allow-Origin", "*")
		return c.Next()
	})
	fiberApp.Use(recover.New())
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SnappFood Ordering Web Service")
	})
	fiberApp.Get("/metrics", monitor.New(monitor.Config{
		Title: "Metrics Page",
	}))
	apiGroup := fiberApp.Group("/api")
	//add routers:
	routers.NewOrderRouter(apiGroup, orderService)
	routers.NewTripRouter(apiGroup, tripService)
	routers.NewVendorRouter(apiGroup, vendorsService)
	routers.NewAgentRouter(apiGroup, agentsService)
	routers.NewSupportRouter(apiGroup, supportService)
	routers.NewAnalyticsRouter(apiGroup, analyticsService)
	//listen fiber:
	log.Fatal(fiberApp.Listen(":" + strconv.Itoa(env.Port)))
}
