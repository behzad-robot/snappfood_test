package repositories

import (
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"

	"gorm.io/gorm"
)

type DelayReportRepository interface {
	Migrate() error
	FindByID(uint) (*models.DelayReport, error)
	FindOne(...interface{}) (*models.DelayReport, error)
	Find(...interface{}) ([]*models.DelayReport, error)
	Insert(*models.DelayReport) error
	Edit(*models.DelayReport) error
	Delete(*models.DelayReport) error
}
type delayReportRepository struct {
	*repo.CrudRepository[models.DelayReport]
}

func NewDelayReportRepository(db *gorm.DB) DelayReportRepository {
	return &delayReportRepository{CrudRepository: &repo.CrudRepository[models.DelayReport]{DB: db}}
}
