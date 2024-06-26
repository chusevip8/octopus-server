package octopus

type RouterGroup struct {
	ScriptRouter
	DeviceRouter
	CommentTaskRouter
	ExecTaskRouter
	ConversationRouter
	ThreadRouter
}
