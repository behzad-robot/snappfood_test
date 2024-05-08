package agents_test

import (
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"
	"snappfood/ordering/services/agents"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type mockAgentRepository struct {
	*repo.CrudRepository[models.Agent]
	mock.Mock
}

func (m *mockAgentRepository) Migrate() error {
	args := m.Called()
	return args.Error(0)
}
func (m *mockAgentRepository) FindByID(ID uint) (*models.Agent, error) {
	args := m.Called(ID)
	var result *models.Agent
	var e error
	if args.Get(0) == nil {
		result = nil
	} else {
		result = args.Get(0).(*models.Agent)
	}
	if args.Get(1) == nil {
		e = nil
	} else {
		e = args.Get(1).(error)
	}
	return result, e
}
func (m *mockAgentRepository) Insert(data *models.Agent) error {
	args := m.Called(data)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}
func Test_Create(t *testing.T) {
	mockRepo := new(mockAgentRepository)
	service := agents.NewAgentService(mockRepo)
	mockRepo.On("Insert", &models.Agent{Name: "test"}).Return(nil)
	agent, err := service.Create(agents.CreateAgentCommand{Name: "test"})
	assert.Equal(t, agent.ID, uint(0))
	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}
func Test_GetOne(t *testing.T) {
	mockRepo := new(mockAgentRepository)
	service := agents.NewAgentService(mockRepo)
	mockRepo.On("FindByID", uint(1)).Return(&models.Agent{Model: gorm.Model{ID: uint(1)}, Name: "Test Agent 1"}, nil)
	mockRepo.On("FindByID", uint(2)).Return(nil, repo.ErrRecordNotFound)

	agent, err := service.GetOne(uint(1))
	assert.Equal(t, agent.ID, uint(1))
	assert.Nil(t, err)

	agent, err = service.GetOne(uint(2))
	assert.Nil(t, agent)
	assert.NotNil(t, err)

	mockRepo.AssertExpectations(t)
}
