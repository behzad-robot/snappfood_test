package repositories

import (
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"

	"gorm.io/gorm"
)

type AgentRepository interface {
	Migrate() error
	FindByID(uint) (*models.Agent, error)
	FindOne(...interface{}) (*models.Agent, error)
	Find(...interface{}) ([]*models.Agent, error)
	Insert(*models.Agent) error
	Edit(*models.Agent) error
	Delete(*models.Agent) error
}
type agentRepository struct {
	*repo.CrudRepository[models.Agent]
}

func NewAgentRepository(db *gorm.DB) AgentRepository {
	return &agentRepository{CrudRepository: &repo.CrudRepository[models.Agent]{DB: db}}
}
