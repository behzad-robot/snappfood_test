package vendors

import (
	"snappfood/ordering/common"
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"
	"snappfood/ordering/internal/repositories"
)

type VendorService interface {
	GetOne(ID uint) (*models.Vendor, *common.ServiceError)
	Create(CreateVendorCommand) (*models.Vendor, *common.ServiceError)
}

func NewVendorService(vendorsRepo repositories.VendorRepository) VendorService {
	return &vendorService{vendorsRepo: vendorsRepo}
}

type vendorService struct {
	vendorsRepo repositories.VendorRepository
}

func (service *vendorService) GetOne(ID uint) (*models.Vendor, *common.ServiceError) {
	item, err := service.vendorsRepo.FindByID(ID)
	if err == repo.ErrRecordNotFound {
		return nil, common.NotFound
	}
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return item, nil
}
func (service *vendorService) Create(command CreateVendorCommand) (*models.Vendor, *common.ServiceError) {
	item := &models.Vendor{
		Name: command.Name,
	}
	err := service.vendorsRepo.Insert(item)
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return item, nil
}
