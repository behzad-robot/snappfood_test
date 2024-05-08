package main

import (
	"fmt"
	"snappfood/ordering/env"
	"snappfood/ordering/infrastructure/database"
	"snappfood/ordering/internal/models"
	"snappfood/ordering/internal/repositories"
	"time"
)

func main() {
	fmt.Println("SnappFood Ordering Test Data")
	env, err := env.CreateEnv()
	if err != nil {
		panic(err)
	}
	db, err := database.CreateGormConnection(env.PostgresqlDatabase)
	if err != nil {
		panic(err)
	}
	//DI:

	vendorsRepo := repositories.NewVendorRepository(db)
	ordersRepo := repositories.NewOrderRepository(db)
	tripsRepo := repositories.NewTripRepository(db)

	agentsRepo := repositories.NewAgentRepository(db)
	// delayReportsRepo := repositories.NewDelayReportRepository(db)

	vendorsRepo.Insert(&models.Vendor{Name: "vendor test"})
	vendorsRepo.Insert(&models.Vendor{Name: "vendor test 2"})
	agentsRepo.Insert(&models.Agent{Name: "agent test"})
	agentsRepo.Insert(&models.Agent{Name: "agent test 2"})
	order := &models.Order{
		UserID:              1,
		VendorID:            1,
		DeliveryTime:        int64(time.Minute * 50),
		InitialDeliveryTime: int64(time.Minute * 50),
	}
	ordersRepo.Insert(order)
	tripsRepo.Insert(&models.Trip{
		OrderID: order.ID,
		BikeID:  1,
		Status:  models.TRIP_STATUS_NONE,
	})
	ordersRepo.Insert(&models.Order{
		UserID:              1,
		VendorID:            2,
		DeliveryTime:        int64(time.Minute * 50),
		InitialDeliveryTime: int64(time.Minute * 50),
	})
}
