// 自动生成模板CmtTaskSetup
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论任务设置 结构体  CmtTaskSetup
type CmtTaskSetup struct {
	global.GVA_MODEL
	AppName              string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称;size:20;" binding:"required"`                                         //应用名称
	TaskTitle            string `json:"taskTitle" form:"taskTitle" gorm:"column:task_title;comment:任务标题;size:128;" binding:"required"`                                  //任务标题
	PostLink             string `json:"postLink" form:"postLink" gorm:"column:post_link;comment:帖子链接;" binding:"required"`                                              //帖子链接
	Keyword              string `json:"keyword" form:"keyword" gorm:"column:keyword;comment:评论关键字;"`                                                                    //评论关键字
	CmtCount             string `json:"cmtCount" form:"cmtCount" gorm:"default:0;column:cmt_count;comment:评论条数;"`                                                       //评论条数
	ReadPostCmtScriptId  uint   `json:"readPostCmtScriptId" form:"readPostCmtScriptId" gorm:"column:read_post_cmt_script_id;comment:读取帖子评论脚本Id;" binding:"required"`    //读取帖子评论脚本Id
	ReplyPostCmtScriptId uint   `json:"replyPostCmtScriptId" form:"replyPostCmtScriptId" gorm:"column:reply_post_cmt_script_id;comment:回复帖子评论脚本Id;" binding:"required"` //回复帖子评论脚本Id
	ReplyMsgCmtScriptId  uint   `json:"replyMsgCmtScriptId" form:"replyMsgCmtScriptId" gorm:"column:reply_msg_cmt_script_id;comment:回复消息评论脚本Id;" binding:"required"`    //回复消息评论脚本Id
	CreatedBy            uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy            uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy            uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 评论任务设置 CmtTaskSetup自定义表名 oct_cmt_task_setup
func (CmtTaskSetup) TableName() string {
	return "oct_cmt_task_setup"
}
