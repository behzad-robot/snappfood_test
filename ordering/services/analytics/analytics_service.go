package analytics

import (
	"snappfood/ordering/common"
	"snappfood/ordering/internal/repositories"
)

type AnalyticsService interface {
	GetWeeklyVendorDelaysReport() (any, *common.ServiceError)
}

func NewAnalyticsService(orderRepo repositories.OrderRepository) AnalyticsService {
	return &analyticsService{orderRepo: orderRepo}
}

type analyticsService struct {
	orderRepo repositories.OrderRepository
}

func (service *analyticsService) GetWeeklyVendorDelaysReport() (any, *common.ServiceError) {
	results, err := service.orderRepo.GetVendorDelaysWeeklyReport()
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return results, nil
}
