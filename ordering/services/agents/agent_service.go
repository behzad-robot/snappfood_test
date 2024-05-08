package agents

import (
	"snappfood/ordering/common"
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"
	"snappfood/ordering/internal/repositories"
)

type AgentService interface {
	GetOne(ID uint) (*models.Agent, *common.ServiceError)
	Create(CreateAgentCommand) (*models.Agent, *common.ServiceError)
}

func NewAgentService(agentsRepo repositories.AgentRepository) AgentService {
	return &agentService{agentsRepo: agentsRepo}
}

type agentService struct {
	agentsRepo repositories.AgentRepository
}

func (service *agentService) GetOne(ID uint) (*models.Agent, *common.ServiceError) {
	item, err := service.agentsRepo.FindByID(ID)
	if err == repo.ErrRecordNotFound {
		return nil, common.NotFound
	}
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return item, nil
}
func (service *agentService) Create(command CreateAgentCommand) (*models.Agent, *common.ServiceError) {
	item := &models.Agent{
		Name: command.Name,
	}
	err := service.agentsRepo.Insert(item)
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return item, nil
}
