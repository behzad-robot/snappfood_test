package repositories

import (
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"
	"time"

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
	GetVendorDelaysWeeklyReport() ([]*models.VendorDelay, error)
}
type orderRepository struct {
	*repo.CrudRepository[models.Order]
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{CrudRepository: &repo.CrudRepository[models.Order]{DB: db}}
}
func (repo *orderRepository) GetVendorDelaysWeeklyReport() ([]*models.VendorDelay, error) {
	results := make([]*models.VendorDelay, 0)
	startOfWeek := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday())+1)
	endOfWeek := startOfWeek.AddDate(0, 0, 6).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	trans := repo.DB.Table("orders").
		Select("SUM(delivery_time-initial_delivery_time) as delay,vendor_id").
		Where("created_at BETWEEN ? AND ?", startOfWeek, endOfWeek).
		Group("vendor_id").
		Order("delay DESC").
		Scan(&results)
	if trans.Error != nil {
		return nil, trans.Error
	}
	return results, nil
}
