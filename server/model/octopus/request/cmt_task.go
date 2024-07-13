package request

type FindCmtTask struct {
	SetupId   uint   `json:"setupId" form:"setupId"`
	ScriptId  uint   `json:"scriptId" form:"scriptId"`
	AppName   string `json:"appName" form:"appName"`
	TaskType  string `json:"taskType" form:"taskType"`
	DeviceId  uint   `json:"deviceId" form:"deviceId"`
	Status    uint   `json:"status" form:"status"`
	Error     string `json:"error" form:"error"`
	CreatedBy uint   `json:"createdBy" form:"createdBy"`
}

type WriteCmtTask struct {
}

type ReplyCmtTask struct {
}
