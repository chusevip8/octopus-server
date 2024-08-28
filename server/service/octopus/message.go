package octopus

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	octopusRes "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/response"
	"gorm.io/gorm"
)

type MessageService struct{}

// CreateMessage 创建私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) CreateMessage(message *octopus.Message) (err error) {
	err = global.GVA_DB.Create(message).Error
	return err
}

// DeleteMessage 删除私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) DeleteMessage(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.Message{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.Message{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMessageByIds 批量删除私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) DeleteMessageByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.Message{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.Message{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMessage 更新私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) UpdateMessage(message octopus.Message) (err error) {
	err = global.GVA_DB.Model(&octopus.Message{}).Where("id = ?", message.ID).Updates(&message).Error
	return err
}

// GetMessage 根据ID获取私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) GetMessage(ID string) (message octopus.Message, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&message).Error
	return
}

// GetMessageInfoList 分页获取私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) GetMessageInfoList(info octopusReq.MessageSearch) (list []octopusRes.MessageRes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.Message{})
	var messages []octopus.Message
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ConversationId != "" {
		db = db.Where("conversation_id=?", info.ConversationId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&messages).Error

	if err != nil {
		return
	}
	err = messageService.markMessagesAsRead(messages)

	var messageResList []octopusRes.MessageRes
	for _, message := range messages {

		messageRes := octopusRes.MessageRes{}
		messageRes.Name = message.Sender
		if message.Mine {
			messageRes.Name = message.Receiver
		}
		messageRes.Date = message.SendAt
		messageRes.Mine = message.Mine
		messageRes.Text = messageService.formatMessageText(message)
		messageRes.Img = "avatar_blue.jpg"
		if message.Mine {
			messageRes.Img = "avatar_red.jpg"
		}
		messageResList = append(messageResList, messageRes)
	}
	return messageResList, total, err
}

func (messageService *MessageService) formatMessageText(message octopus.Message) octopusRes.Text {
	if !message.Mine {
		return octopusRes.Text{Text: message.Content}
	}

	var color, status, taskErr string
	switch message.Task.Status {
	case 0:
		color, status, taskErr = "red", "失败", "-任务已删除"
	case 1:
		color, status, taskErr = "yellow", "新建", ""
	case 2:
		color, status, taskErr = "yellow", "发送中", ""
	case 3:
		color, status, taskErr = "green", "已发送", ""
	default:
		color, status, taskErr = "red", "失败", "-"+message.Task.Error
	}

	text := fmt.Sprintf(`%s<p style="color: %s;">%s%s</p>`, message.Content, color, status, taskErr)
	return octopusRes.Text{Text: text}
}

func (messageService *MessageService) markMessagesAsRead(messages []octopus.Message) error {
	if len(messages) == 0 {
		return nil
	}

	messageIds := make([]uint, len(messages))
	for i, message := range messages {
		messageIds[i] = message.ID
	}

	err := global.GVA_DB.Model(&octopus.Message{}).Where("id IN (?) AND unread = ?", messageIds, true).Update("unread", false).Error
	return err
}
