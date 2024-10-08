package octopus

type RouterGroup struct {
	ScriptRouter
	DeviceRouter
	CmtThreadRouter
	CmtConversationRouter
	CommentRouter
	TaskRouter
	CmtTaskSetupRouter
	CmtTaskRouter
	TaskParamsRouter
	IntervalTaskSetupRouter
	IntervalTaskRouter
	GenericTaskSetupRouter
	DataFileRouter
	GenericTaskRouter
	TaskBindDataRouter
	MsgTaskSetupRouter
	MsgConversationRouter
	MessageRouter
	MsgTaskRouter
}
