// 自动生成模板Comment
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论 结构体  Comment
type Comment struct {
	global.GVA_MODEL
	ConversationId   uint   `json:"conversationId" form:"conversationId" gorm:"column:conversation_id;comment:评论记录Id;" binding:"required"` //评论记录Id
	TaskId           uint   `json:"taskId" form:"taskId" gorm:"column:task_id;comment:任务Id;" binding:"required"`                           //任务Id
	Task             Task   `json:"task" gorm:"foreignKey:TaskId;references:ID;comment:关联任务;"`
	Commenter        string `json:"commenter" form:"commenter" gorm:"column:commenter;comment:发评论者;size:64;"`                            //发评论者
	CommenterId      string `json:"commenterId" form:"commenterId" gorm:"column:commenter_id;comment:发评论者Id;size:128;"`                  //发评论者Id
	CommentReplier   string `json:"commentReplier" form:"commentReplier" gorm:"column:comment_replier;comment:评论回复者;size:64;"`           //评论回复者
	CommentReplierId string `json:"commentReplierId" form:"commentReplierId" gorm:"column:comment_replier_id;comment:评论回复者Id;size:128;"` //评论回复者Id
	Content          string `json:"content" form:"content" gorm:"column:content;comment:评论内容;" binding:"required"`                       //评论内容
	PostAt           string `json:"postAt" form:"postAt" gorm:"column:post_at;comment:评论时间;size:20;"`                                    //评论时间
	Unread           bool   `json:"unread" form:"unread" gorm:"column:unread;comment:是否未读;default:false;"`
	Mine             bool   `json:"mine" form:"mine" gorm:"column:mine;comment:是否自己回复;default:false;"`
}

// TableName 评论 Comment自定义表名 oct_comment
func (Comment) TableName() string {
	return "oct_comment"
}
