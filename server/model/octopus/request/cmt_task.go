package request

type FindCmtTask struct {
	TaskSetupId  uint   `json:"taskSetupId" form:"taskSetupId"`
	ScriptId     uint   `json:"scriptId" form:"scriptId"`
	AppName      string `json:"appName" form:"appName"`
	MainTaskType string `json:"mainTaskType" form:"mainTaskType"`
	SubTaskType  string `json:"subTaskType" form:"subTaskType"`
	DeviceId     uint   `json:"deviceId" form:"deviceId"`
	Status       uint   `json:"status" form:"status"`
	Error        string `json:"error" form:"error"`
	CreatedBy    uint   `json:"createdBy" form:"createdBy"`
}

type WriteCmtTask struct {
}

type ReplyCmtTask struct {
}
