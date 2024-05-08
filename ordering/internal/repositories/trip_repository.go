package repositories

import (
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	Migrate() error
	FindByID(uint) (*models.Trip, error)
	FindOne(...interface{}) (*models.Trip, error)
	Find(...interface{}) ([]*models.Trip, error)
	Insert(*models.Trip) error
	Edit(*models.Trip) error
	Delete(*models.Trip) error
}
type tripRepository struct {
	*repo.CrudRepository[models.Trip]
}

func NewTripRepository(db *gorm.DB) TripRepository {
	return &tripRepository{CrudRepository: &repo.CrudRepository[models.Trip]{DB: db}}
}
