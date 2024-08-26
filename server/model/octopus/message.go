// 自动生成模板Message
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 私信 结构体  Message
type Message struct {
	global.GVA_MODEL
	ConversationId uint   `json:"conversationId" form:"conversationId" gorm:"column:conversation_id;comment:私信会话Id;"` //私信会话Id
	TaskId         uint   `json:"taskId" form:"taskId" gorm:"column:task_id;comment:任务Id;" binding:"required"`        //任务Id
	Task           Task   `json:"task" gorm:"foreignKey:TaskId;references:ID;comment:关联任务;"`
	Sender         string `json:"sender" form:"sender" gorm:"column:sender;comment:发送者;size:64;"`                 //发送者
	SenderId       string `json:"senderId" form:"senderId" gorm:"column:sender_id;comment:发送者Id;size:128;"`       //发送者Id
	Receiver       string `json:"receiver" form:"receiver" gorm:"column:receiver;comment:接收者;size:64;"`           //接收者
	ReceiverId     string `json:"receiverId" form:"receiverId" gorm:"column:receiver_id;comment:接收者Id;size:128;"` //接收者Id
	Content        string `json:"content" form:"content" gorm:"column:content;comment:消息内容;" binding:"required"`  //消息内容
	SendAt         string `json:"sendAt" form:"sendAt" gorm:"column:send_at;comment:消息时间;size:20;"`               //消息时间
	Unread         bool   `json:"unread" form:"unread" gorm:"column:unread;comment:是否未读;default:false;"`
	Mine           bool   `json:"mine" form:"mine" gorm:"column:mine;comment:是否自己回复;default:false;"`
	CreatedBy      uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 私信 Message自定义表名 oct_message
func (Message) TableName() string {
	return "oct_message"
}
