package octopus

type ServiceGroup struct {
	ScriptService
	DeviceService
	CmtThreadService
	CmtConversationService
	CommentService
	TaskService
	CmtTaskSetupService
	CmtTaskService
	TaskParamsService
	IntervalTaskSetupService
	IntervalTaskService
	GenericTaskSetupService
	DataFileService
	GenericTaskService
	TaskBindDataService
	MsgTaskSetupService
	MsgConversationService
	MessageService
}
