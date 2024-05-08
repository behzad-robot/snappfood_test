package repositories

import (
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"

	"gorm.io/gorm"
)

type VendorRepository interface {
	Migrate() error
	FindByID(uint) (*models.Vendor, error)
	FindOne(...interface{}) (*models.Vendor, error)
	Find(...interface{}) ([]*models.Vendor, error)
	Insert(*models.Vendor) error
	Edit(*models.Vendor) error
	Delete(*models.Vendor) error
}
type vendorRepository struct {
	*repo.CrudRepository[models.Vendor]
}

func NewVendorRepository(db *gorm.DB) VendorRepository {
	return &vendorRepository{CrudRepository: &repo.CrudRepository[models.Vendor]{DB: db}}
}
