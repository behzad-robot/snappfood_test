package main

import (
	"fmt"
	"snappfood/ordering/env"
	"snappfood/ordering/infrastructure/database"
	"snappfood/ordering/internal/repositories"
)

func main() {
	fmt.Println("SnappFood Ordering Migrate")
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
	delayReportsRepo := repositories.NewDelayReportRepository(db)
	delayResultsRepo := repositories.NewDelayResultRepository(db)

	if err := vendorsRepo.Migrate(); err != nil {
		panic(err)
	}
	if err := ordersRepo.Migrate(); err != nil {
		panic(err)
	}
	if err := tripsRepo.Migrate(); err != nil {
		panic(err)
	}

	if err := agentsRepo.Migrate(); err != nil {
		panic(err)
	}
	if err := delayReportsRepo.Migrate(); err != nil {
		panic(err)
	}
	if err := delayResultsRepo.Migrate(); err != nil {
		panic(err)
	}

	fmt.Println("SnappFood Ordering Migrate Done!")
}
