package request

type MessageReq struct {
	TaskId     string        `json:"taskId" form:"taskId"`
	Receiver   string        `json:"receiver" form:"receiver"`     //接收者
	ReceiverId string        `json:"receiverId" form:"receiverId"` //接收者Id
	Sender     string        `json:"sender" form:"sender"`         //发送者
	SenderId   string        `json:"senderId" form:"senderId"`     //发送者Id
	Messages   []MessageItem `json:"messages" form:"messages"`
}

type MessageItem struct {
	Content  string `json:"content" form:"content"`
	SendAt   string `json:"sendAt" form:"sendAt" `
	Writer   string `json:"writer" form:"writer"`     //发送者
	WriterId string `json:"writerId" form:"writerId"` //发送者Id
}
