package response

type CmtConversationRes struct {
	ID             uint   `json:"ID"`             // 主键ID
	Commenter      string `json:"commenter"`      //发评论者
	CommentReplier string `json:"commentReplier"` //评论回复者
	Unread         int64  `json:"unread"`
}
