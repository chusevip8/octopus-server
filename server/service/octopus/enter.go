package octopus

type ServiceGroup struct {
	ScriptService
	DeviceService
	CmtMsgMgrService
	CmtThreadService
	CmtConversationService
	CommentService
	TaskService
	CmtTaskParamsService
	CmtTaskSetupService
	CmtTaskService
}
