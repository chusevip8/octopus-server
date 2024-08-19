package request

type ReadPostCmtTask struct {
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

type ReplyCmtTask struct {
	ThreadId       string `json:"threadId" form:"threadId"`
	ConversationId string `json:"conversationId" form:"conversationId"`
	CmtContent     string `json:"cmtContent" form:"cmtContent"`
}

type CommentReq struct {
	TaskId           string `json:"taskId" form:"taskId"`
	Poster           string `json:"poster" form:"poster" `
	PostTitle        string `json:"postTitle" form:"postTitle" `
	PostDesc         string `json:"postDesc" form:"postDesc" `
	Commenter        string `json:"commenter" form:"commenter"`                //发评论者
	CommenterId      string `json:"commenterId" form:"commenterId" `           //发评论者Id
	CommentReplier   string `json:"commentReplier" form:"commentReplier" `     //评论回复者
	CommentReplierId string `json:"commentReplierId" form:"commentReplierId" ` //评论回复者Id
	Content          string `json:"content" form:"content" `                   //评论内容
	PostAt           string `json:"postAt" form:"postAt" `
	CmtFrom          string `json:"cmtFrom" form:"cmtFrom" `
}

type PostCommentReq struct {
	TaskId           string            `json:"taskId" form:"taskId"`
	Poster           string            `json:"poster" form:"poster" `
	PostTitle        string            `json:"postTitle" form:"postTitle" `
	PostDesc         string            `json:"postDesc" form:"postDesc" `
	CommentReplier   string            `json:"commentReplier" form:"commentReplier" `     //评论回复者
	CommentReplierId string            `json:"commentReplierId" form:"commentReplierId" ` //评论回复者Id
	Comments         []PostCommentItem `json:"comments" form:"comments"`
}
type PostCommentItem struct {
	Content     string `json:"content" form:"content" ` //评论内容
	PostAt      string `json:"postAt" form:"postAt" `
	Commenter   string `json:"commenter" form:"commenter"`      //发评论者
	CommenterId string `json:"commenterId" form:"commenterId" ` //发评论者Id
}

type MsgCommentReq struct {
	TaskId           string           `json:"taskId" form:"taskId"`
	CommentReplier   string           `json:"commentReplier" form:"commentReplier" `     //评论回复者
	CommentReplierId string           `json:"commentReplierId" form:"commentReplierId" ` //评论回复者Id
	Comments         []MsgCommentItem `json:"comments" form:"comments"`
}
type MsgCommentItem struct {
	Content     string `json:"content" form:"content"`
	PostAt      string `json:"postAt" form:"postAt" `
	Commenter   string `json:"commenter" form:"commenter"`      //发评论者
	CommenterId string `json:"commenterId" form:"commenterId" ` //发评论者Id
	Poster      string `json:"poster" form:"poster" `
	PostTitle   string `json:"postTitle" form:"postTitle" `
	PostDesc    string `json:"postDesc" form:"postDesc" `
}
