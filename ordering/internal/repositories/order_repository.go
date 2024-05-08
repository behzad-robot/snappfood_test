package repositories

import (
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Migrate() error
	FindByID(uint) (*models.Order, error)
	FindOne(...interface{}) (*models.Order, error)
	Find(...interface{}) ([]*models.Order, error)
	Insert(*models.Order) error
	Edit(*models.Order) error
	Delete(*models.Order) error
}
type orderRepository struct {
	*repo.CrudRepository[models.Order]
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{CrudRepository: &repo.CrudRepository[models.Order]{DB: db}}
}
