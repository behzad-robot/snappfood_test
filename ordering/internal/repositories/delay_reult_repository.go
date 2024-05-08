package repositories

import (
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"
	"time"

	"gorm.io/gorm"
)

type DelayResultRepository interface {
	Migrate() error
	FindByID(uint) (*models.DelayResult, error)
	FindOne(...interface{}) (*models.DelayResult, error)
	Find(...interface{}) ([]*models.DelayResult, error)
	Insert(*models.DelayResult) error
	Edit(*models.DelayResult) error
	Delete(*models.DelayResult) error
	GetWeeklyVendorDelaysReport() ([]*models.DelayResultByVendor, error)
}
type delayResultRepository struct {
	*repo.CrudRepository[models.DelayResult]
}

func NewDelayResultRepository(db *gorm.DB) DelayResultRepository {
	return &delayResultRepository{CrudRepository: &repo.CrudRepository[models.DelayResult]{DB: db}}
}
func (repo *delayResultRepository) GetWeeklyVendorDelaysReport() ([]*models.DelayResultByVendor, error) {
	results := make([]*models.DelayResultByVendor, 0)
	// Calculate start and end of the current week
	startOfWeek := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday())+1)
	endOfWeek := startOfWeek.AddDate(0, 0, 6).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
	trans := repo.DB.Table("delay_results").
		Select("vendor_id,SUM(delay_time) as total_delay").
		Where("created_at BETWEEN ? AND ?", startOfWeek, endOfWeek).
		Group("vendor_id").Scan(&results)
	return results, trans.Error
}
