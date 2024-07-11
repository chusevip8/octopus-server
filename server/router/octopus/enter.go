package octopus

type RouterGroup struct {
	ScriptRouter
	DeviceRouter
	CmtThreadRouter
	CmtConversationRouter
	CommentRouter
	TaskRouter
	CmtTaskParamsRouter
	CmtTaskSetupRouter
}
