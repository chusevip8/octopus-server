package request

type GenericTask struct {
	Batch        bool   `json:"batch" form:"batch"`
	DeviceGroup  string `json:"deviceGroup" form:"deviceGroup"`
	TaskSetupId  uint   `json:"taskSetupId" form:"taskSetupId"`
	MainTaskType string `json:"mainTaskType" form:"mainTaskType"`
	SubTaskType  string `json:"subTaskType" form:"subTaskType"`
	DeviceId     uint   `json:"deviceId" form:"deviceId"`
	Status       uint   `json:"status" form:"status"`
	Error        string `json:"error" form:"error"`
	CreatedBy    uint   `json:"createdBy" form:"createdBy"`
}

type BindTaskData struct {
	TaskSetupId  string `json:"taskSetupId" form:"taskSetupId"`
	MainTaskType string `json:"mainTaskType" form:"mainTaskType"`
	SubTaskType  string `json:"subTaskType" form:"subTaskType"`
}
