// 自动生成模板Message
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 消息 结构体  Message
type Message struct {
	global.GVA_MODEL
	ThreadID   uint   `json:"threadID" form:"threadID" gorm:"column:thread_id;comment:消息组ID;" binding:"required"`                 //消息组ID
	Sender     string `json:"sender" form:"sender" gorm:"column:sender;comment:消息发送者;size:50;" binding:"required"`                //消息发送者
	SenderId   string `json:"senderId" form:"senderId" gorm:"column:sender_id;comment:消息发送者ID;size:20;" binding:"required"`       //消息发送者ID
	Receiver   string `json:"receiver" form:"receiver" gorm:"column:receiver;comment:消息接收者;size:50;" binding:"required"`          //消息接收者
	ReceiverId string `json:"receiverId" form:"receiverId" gorm:"column:receiver_id;comment:消息接收者ID;size:20;" binding:"required"` //消息接收者ID
	Status     uint   `json:"status" form:"status" gorm:"default:0;column:status;comment:消息状态;" binding:"required"`               //消息状态
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 消息 Message自定义表名 oct_message
func (Message) TableName() string {
	return "oct_message"
}
