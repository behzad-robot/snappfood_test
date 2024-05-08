package orders

import (
	"snappfood/ordering/common"
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"
	"snappfood/ordering/internal/repositories"
	"time"
)

type OrderService interface {
	GetOne(ID uint) (*models.Order, *common.ServiceError)
	Create(CreateOrderCommand) (*models.Order, *common.ServiceError)
	//inside system functionalities:
	AddToDeliveryTime(orderID uint, duration time.Duration) *common.ServiceError
}

func NewOrderService(ordersRepo repositories.OrderRepository) OrderService {
	return &orderService{ordersRepo: ordersRepo}
}

type orderService struct {
	ordersRepo repositories.OrderRepository
}

func (service *orderService) GetOne(ID uint) (*models.Order, *common.ServiceError) {
	item, err := service.ordersRepo.FindByID(ID)
	if err == repo.ErrRecordNotFound {
		return nil, common.NotFound
	}
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return item, nil
}
func (service *orderService) Create(command CreateOrderCommand) (*models.Order, *common.ServiceError) {
	item := &models.Order{
		UserID:              command.UserID,
		VendorID:            command.VendorID,
		DeliveryTime:        command.DeliveryTime,
		InitialDeliveryTime: command.DeliveryTime,
	}
	err := service.ordersRepo.Insert(item)
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return item, nil
}
func (service *orderService) AddToDeliveryTime(orderID uint, duration time.Duration) *common.ServiceError {
	order, serviceErr := service.GetOne(orderID)
	if serviceErr != nil {
		return serviceErr
	}
	order.DeliveryTime += int64(duration)
	err := service.ordersRepo.Edit(order)
	if err != nil {
		return common.NewServiceError(500, err)
	}
	return nil
}
