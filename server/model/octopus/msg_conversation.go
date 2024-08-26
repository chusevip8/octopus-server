// 自动生成模板MsgConversaton
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 私信会话纪录 结构体  MsgConversation
type MsgConversation struct {
	global.GVA_MODEL
	Sender      string `json:"sender" form:"sender" gorm:"column:sender;comment:发送者;size:64;"`                 //发送者
	SenderId    string `json:"senderId" form:"senderId" gorm:"column:sender_id;comment:发送者Id;size:128;"`       //发送者Id
	Receiver    string `json:"receiver" form:"receiver" gorm:"column:receiver;comment:接收者;size:64;"`           //接收者
	ReceiverId  string `json:"receiverId" form:"receiverId" gorm:"column:receiver_id;comment:接收者Id;size:128;"` //接收者Id
	UnreadCount uint   `json:"unreadCount" form:"unreadCount" gorm:"column:unread_count;comment:未读数;"`         //未读数
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 私信会话纪录 MsgConversation自定义表名 oct_msg_conversation
func (MsgConversation) TableName() string {
	return "oct_msg_conversation"
}
