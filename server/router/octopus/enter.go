package octopus

type RouterGroup struct {
	ScriptRouter
	DeviceRouter
	CommentTaskRouter
	ExecTaskRouter
	ConversationRouter
	ThreadRouter
	MessageRouter
	CmtTaskOptionRouter
	CmtThreadRouter
	CmtConversationRouter
	CommentRouter
	TaskRouter
}
