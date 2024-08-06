package request

type TaskSetup struct {
	TaskSetupId  string `json:"taskSetupId" form:"taskSetupId"`
	MainTaskType string `json:"mainTaskType" form:"mainTaskType"`
	SubTaskType  string `json:"subTaskType" form:"subTaskType"`
}
