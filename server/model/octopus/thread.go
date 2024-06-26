// 自动生成模板Thread
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 消息组 结构体  Thread
type Thread struct {
	global.GVA_MODEL
	ConversationID uint   `json:"conversationID" form:"conversationID" gorm:"column:conversation_id;comment:会话ID;" binding:"required"` //会话ID
	Sender         string `json:"sender" form:"sender" gorm:"column:sender;comment:消息发送者;size:50;" binding:"required"`                 //消息发送者
	SenderId       string `json:"senderId" form:"senderId" gorm:"column:sender_id;comment:消息发送者ID;size:20;" binding:"required"`        //消息发送者ID
	Receiver       string `json:"receiver" form:"receiver" gorm:"column:receiver;comment:消息接收者;size:50;" binding:"required"`           //消息接收者
	ReceiverId     string `json:"receiverId" form:"receiverId" gorm:"column:receiver_id;comment:消息接收者ID;size:20;" binding:"required"`  //消息接收者ID
	CreatedBy      uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 消息组 Thread自定义表名 oct_thread
func (Thread) TableName() string {
	return "oct_thread"
}
