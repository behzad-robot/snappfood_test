package analytics

import (
	"snappfood/ordering/common"
	"snappfood/ordering/internal/models"
	"snappfood/ordering/internal/repositories"
)

type AnalyticsService interface {
	CreateDelayResult(CreateDelayResultCommand) *common.ServiceError
	GetWeeklyVendorDelaysReport() (any, *common.ServiceError)
}

func NewAnalyticsService(delayResultsRepo repositories.DelayResultRepository) AnalyticsService {
	return &analyticsService{delayResultsRepo: delayResultsRepo}
}

type analyticsService struct {
	delayResultsRepo repositories.DelayResultRepository
}

func (service *analyticsService) CreateDelayResult(command CreateDelayResultCommand) *common.ServiceError {
	err := service.delayResultsRepo.Insert(&models.DelayResult{
		OrderID:   command.OrderID,
		VendorID:  command.VendorID,
		DelayTime: command.DelayTime,
	})
	if err != nil {
		return common.NewServiceError(500, err)
	}
	return nil
}
func (service *analyticsService) GetWeeklyVendorDelaysReport() (any, *common.ServiceError) {
	results, err := service.delayResultsRepo.GetWeeklyVendorDelaysReport()
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return results, nil
}
