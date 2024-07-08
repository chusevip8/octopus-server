// 自动生成模板CmtTaskOption
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论任务参数 结构体  CmtTaskOption
type CmtTaskOption struct {
	global.GVA_MODEL
	AppName              string `json:"appName" form:"appName" gorm:"column:app_name;comment:App名称;size:20;" binding:"required"`                        //App名称
	TaskTitle            string `json:"taskTitle" form:"taskTitle" gorm:"column:task_title;comment:任务标题;size:128;" binding:"required"`                  //任务标题
	PostLink             string `json:"postLink" form:"postLink" gorm:"column:post_link;comment:帖子链接;" binding:"required"`                              //帖子链接
	Keyword              string `json:"keyword" form:"keyword" gorm:"column:keyword;comment:评论查找关键字;"`                                                  //评论查找关键字
	Commenter            string `json:"commenter" form:"commenter" gorm:"column:commenter;comment:发评论者;size:64;"`                                       //发评论者
	CommenterId          string `json:"commenterId" form:"commenterId" gorm:"column:commenter_id;comment:发评论者Id;size:20;"`                              //发评论者Id
	CommentId            uint   `json:"commentId" form:"commentId" gorm:"column:comment_id;comment:评论消息Id;"`                                            //评论消息Id
	FindCommentScriptId  uint   `json:"findCommentScriptId" form:"findCommentScriptId" gorm:"column:find_comment_script_id;comment:查找评论脚本Id;"`          //查找评论脚本Id
	WriteCommentScriptId uint   `json:"writeCommentScriptId" form:"writeCommentScriptId" gorm:"column:write_comment_script_id;comment:第一次回复评论脚本Id;"`    //第一次回复评论脚本Id
	ReplyCommentScriptId uint   `json:"replyCommentScriptId" form:"replyCommentScriptId" gorm:"column:reply_comment_script_id;comment:回复评论脚本id;"`       //回复评论脚本id
	IsFindCommentTask    bool   `json:"isFindCommentTask" form:"isFindCommentTask" gorm:"default:false;column:is_find_comment_task;comment:是否为查找评论任务;"` //是否为查找评论任务
	CreatedBy            uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy            uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy            uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 评论任务参数 CmtTaskOption自定义表名 oct_cmt_task_option
func (CmtTaskOption) TableName() string {
	return "oct_cmt_task_option"
}
