package request

type CmtMsg struct {
	UserName   string `json:"userName"`
	AppName    string `json:"appName"`
	TaskId     uint   `json:"taskId"`
	Sender     string `json:"sender"`
	Receiver   string `json:"receiver"`
	SenderId   uint   `json:"senderId"`
	ReceiverId uint   `json:"receiverId"`
	Content    string `json:"content"`
}
