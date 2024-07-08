// 自动生成模板Comment
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论 结构体  Comment
type Comment struct {
	global.GVA_MODEL
	ConversationId   uint   `json:"conversationId" form:"conversationId" gorm:"column:conversation_id;comment:评论记录Id;" binding:"required"` //评论记录Id
	Commenter        string `json:"commenter" form:"commenter" gorm:"column:commenter;comment:发评论者;size:64;"`                              //发评论者
	CommenterId      string `json:"commenterId" form:"commenterId" gorm:"column:commenter_id;comment:发评论者Id;size:20;"`                     //发评论者Id
	CommentReplier   string `json:"commentReplier" form:"commentReplier" gorm:"column:comment_replier;comment:评论回复者;size:64;"`             //评论回复者
	CommentReplierId string `json:"commentReplierId" form:"commentReplierId" gorm:"column:comment_replier_id;comment:评论回复者Id;size:20;"`    //评论回复者Id
	Content          string `json:"content" form:"content" gorm:"column:content;comment:评论内容;" binding:"required"`                         //评论内容
	Status           uint   `json:"status" form:"status" gorm:"default:0;column:status;comment:评论状态;" binding:"required"`                  //评论状态
}

// TableName 评论 Comment自定义表名 oct_comment
func (Comment) TableName() string {
	return "oct_comment"
}
