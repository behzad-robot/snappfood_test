package trips

import (
	"snappfood/ordering/common"
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"
	"snappfood/ordering/internal/repositories"
	"snappfood/ordering/services/analytics"
	"snappfood/ordering/services/orders"
	"time"
)

type TripService interface {
	GetOneByOrderID(orderID uint) (*models.Trip, *common.ServiceError)
	Create(command CreateTripCommand) (*models.Trip, *common.ServiceError)
	ChangeStatus(command ChangeTripStatusCommand) *common.ServiceError
}

func NewTripService(tripsRepo repositories.TripRepository, delayResultsService analytics.AnalyticsService, orderService orders.OrderService) TripService {
	return &tripService{
		tripsRepo:           tripsRepo,
		delayResultsService: delayResultsService,
		orderService:        orderService,
	}
}

type tripService struct {
	tripsRepo           repositories.TripRepository
	delayResultsService analytics.AnalyticsService
	orderService        orders.OrderService
}

func (service *tripService) GetOneByOrderID(orderID uint) (*models.Trip, *common.ServiceError) {
	trip, err := service.tripsRepo.FindOne(map[string]any{"order_id": orderID})
	if err == repo.ErrRecordNotFound {
		return nil, common.NotFound
	}
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return trip, nil
}
func (service *tripService) Create(command CreateTripCommand) (*models.Trip, *common.ServiceError) {
	oldItem, err := service.tripsRepo.FindOne(map[string]any{"order_id": command.OrderID})
	if err != nil && err != repo.ErrRecordNotFound {
		return nil, common.NewServiceError(500, err)
	}
	if err == nil {
		return oldItem, nil
	}
	item := &models.Trip{
		OrderID: command.OrderID,
		BikeID:  command.BikeID,
		Status:  models.TRIP_STATUS_NONE,
	}
	err = service.tripsRepo.Insert(item)
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return item, nil
}
func (service *tripService) ChangeStatus(command ChangeTripStatusCommand) *common.ServiceError {
	trip, serviceErr := service.GetOneByOrderID(command.OrderID)
	if serviceErr != nil {
		return serviceErr
	}
	trip.Status = command.Status
	if command.Status == models.TRIP_STATUS_DELIVERED {
		t := time.Now()
		trip.DeliveredAt = &t
	}
	err := service.tripsRepo.Edit(trip)
	if err != nil {
		return common.NewServiceError(500, err)
	}
	//check if we need delay result:
	if command.Status == models.TRIP_STATUS_DELIVERED {
		order, serviceErr := service.orderService.GetOne(trip.OrderID)
		if serviceErr == nil {
			//order had delay:
			if order.HasDelay() {
				service.delayResultsService.CreateDelayResult(analytics.CreateDelayResultCommand{
					OrderID:   order.ID,
					VendorID:  order.VendorID,
					DelayTime: int64(order.GetDelay()),
				})
			}
		}
	}
	return nil
}
