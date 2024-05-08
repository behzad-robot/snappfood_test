package support

import (
	"encoding/json"
	"fmt"
	"snappfood/ordering/common"
	"snappfood/ordering/common/repo"
	"snappfood/ordering/internal/models"
	"snappfood/ordering/internal/repositories"
	"snappfood/ordering/services/agents"
	"snappfood/ordering/services/orders"
	"snappfood/ordering/services/rabbit"
	"snappfood/ordering/services/trips"
	"time"
)

const (
	agent_tasks_queue = "agentTasks"
)

type SupportService interface {
	GetOrderDelayStatus(orderID uint) ([]*models.DelayReport, *common.ServiceError)
	ReportOrderDelay(ReportOrderDelayCommand) (*models.DelayReport, *common.ServiceError)
	GetTaskForAgent(agentID uint) (*models.DelayReport, *common.ServiceError)
	UpdateTaskForAgent(UpdateTaskForAgentCommand) *common.ServiceError
}

func NewSupportService(delayReportsRepo repositories.DelayReportRepository,
	agentService agents.AgentService, orderService orders.OrderService, tripService trips.TripService,
	rabbitHelperService rabbit.RabbitHelperService) SupportService {
	return &supportService{
		delayReportsRepo:    delayReportsRepo,
		agentService:        agentService,
		orderService:        orderService,
		tripService:         tripService,
		rabbitHelperService: rabbitHelperService,
	}
}

type supportService struct {
	delayReportsRepo    repositories.DelayReportRepository
	agentService        agents.AgentService
	orderService        orders.OrderService
	tripService         trips.TripService
	rabbitHelperService rabbit.RabbitHelperService
}

func (service *supportService) GetOrderDelayStatus(orderID uint) ([]*models.DelayReport, *common.ServiceError) {
	order, serviceErr := service.orderService.GetOne(orderID)
	if serviceErr != nil {
		return nil, serviceErr
	}
	delays, err := service.delayReportsRepo.Find(map[string]any{"order_id": order.ID})
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return delays, nil
}
func (service *supportService) ReportOrderDelay(command ReportOrderDelayCommand) (*models.DelayReport, *common.ServiceError) {
	order, serviceErr := service.orderService.GetOne(command.OrderID)
	//order is not late chill:
	if order.GetDeliveredAt().After(time.Now()) {
		return nil, common.NotAcceptable
	}
	if serviceErr != nil {
		return nil, serviceErr
	}
	trip, serviceErr := service.tripService.GetOneByOrderID(order.ID)
	fmt.Println(trip)
	fmt.Println(serviceErr)
	hasTrip := serviceErr == nil
	hasGoodTripStatus := hasTrip && (trip.Status == models.TRIP_STATUS_AT_VENDOR || trip.Status == models.TRIP_STATUS_ASSIGNED || trip.Status == models.TRIP_STATUS_PICKED)
	fmt.Println("hasTrip=", hasTrip)
	fmt.Println("hasTripGoodStatus=", hasGoodTripStatus)
	if hasGoodTripStatus {
		//edit order  with new ETA:
		service.orderService.AddToDeliveryTime(order.ID, time.Minute*10) //TODO: randomize!
		//create delay report too: dont need agent attention
		delayReport := &models.DelayReport{
			OrderID:            order.ID,
			TripStatusAtReport: trip.Status,
			AgentID:            -1,
		}
		err := service.delayReportsRepo.Insert(delayReport)
		if err != nil {
			return nil, common.NewServiceError(500, err)
		}
		return delayReport, nil
	}
	//create delay report that needs agent attention:
	statusAtReport := models.TRIP_STATUS_NONE
	if hasTrip {
		statusAtReport = trip.Status
	}
	delayReport := &models.DelayReport{
		OrderID:            order.ID,
		TripStatusAtReport: statusAtReport,
		AgentID:            0,
	}
	err := service.delayReportsRepo.Insert(delayReport)
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	//also publish task on rabbitMQ:
	js, _ := json.Marshal(delayReport)
	service.rabbitHelperService.PublishTo(agent_tasks_queue, js)
	return delayReport, nil
}
func (service *supportService) GetTaskForAgent(agentID uint) (*models.DelayReport, *common.ServiceError) {
	agent, serviceErr := service.agentService.GetOne(agentID)
	if serviceErr != nil {
		return nil, serviceErr
	}
	//check if agent has a task right now:
	currentTask, err := service.delayReportsRepo.FindOne(map[string]any{"agent_id": agent.ID, "is_replied_by_agent": false})
	if err != nil && err != repo.ErrRecordNotFound {
		return nil, common.NewServiceError(500, err)
	}
	if err == nil {
		return currentTask, nil
	}
	//check queue for a task:
	jsMessage, hasSth, err := service.rabbitHelperService.SimpleConsumeOne(agent_tasks_queue)
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	if !hasSth {
		return nil, common.NewServiceMessage(404, "no tasks available")
	}
	newTask := &models.DelayReport{}
	json.Unmarshal(jsMessage, &newTask)
	newTask.AgentID = int(agent.ID)
	err = service.delayReportsRepo.Edit(newTask)
	if err != nil {
		return nil, common.NewServiceError(500, err)
	}
	return newTask, nil
}
func (service *supportService) UpdateTaskForAgent(command UpdateTaskForAgentCommand) *common.ServiceError {
	task, err := service.delayReportsRepo.FindByID(command.TaskID)
	if err == repo.ErrRecordNotFound {
		return common.NotFound
	}
	if task.AgentID != int(command.AgentID) {
		return common.BadParameters
	}
	if task.IsRepliedByAgent {
		return common.NotAcceptable
	}
	task.AgentMessage = command.AgentMessage
	task.IsRepliedByAgent = command.MarkAsDone
	err = service.delayReportsRepo.Edit(task)
	if err != nil {
		return common.NewServiceError(500, err)
	}
	return nil
}
