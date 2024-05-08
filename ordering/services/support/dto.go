package support

type ReportOrderDelayCommand struct {
	OrderID uint `json:"orderID"`
}
type UpdateTaskForAgentCommand struct {
	TaskID       uint   `json:"taskID"`
	AgentID      uint   `json:"agentID"`
	MarkAsDone   bool   `json:"markAsDone"`
	AgentMessage string `json:"agentMessage"`
}
