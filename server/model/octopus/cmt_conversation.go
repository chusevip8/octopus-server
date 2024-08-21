// 自动生成模板CmtConversation
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论会话记录 结构体  CmtConversation
type CmtConversation struct {
	global.GVA_MODEL
	ThreadId         uint   `json:"threadId" form:"threadId" gorm:"column:thread_id;comment:会话Id;" binding:"required"`                   //会话Id
	Commenter        string `json:"commenter" form:"commenter" gorm:"column:commenter;comment:发评论者;size:64;"`                            //发评论者
	CommenterId      string `json:"commenterId" form:"commenterId" gorm:"column:commenter_id;comment:发评论者Id;size:128;"`                  //发评论者Id
	CommentReplier   string `json:"commentReplier" form:"commentReplier" gorm:"column:comment_replier;comment:评论回复者;size:64;"`           //评论回复者
	CommentReplierId string `json:"commentReplierId" form:"commentReplierId" gorm:"column:comment_replier_id;comment:评论回复者Id;size:128;"` //评论回复者Id
	CreatedBy        uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint   `gorm:"column:deleted_by;comment:删除者"`
	UnreadCount      uint   `json:"unreadCount" form:"unreadCount" gorm:"column:unread_count"`
}

// TableName 评论会话记录 CmtConversation自定义表名 oct_cmt_conversation
func (CmtConversation) TableName() string {
	return "oct_cmt_conversation"
}
