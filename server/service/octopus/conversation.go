package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
)

type ConversationService struct{}

// CreateConversation 创建消息会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) CreateConversation(conversation *octopus.Conversation) (err error) {
	err = global.GVA_DB.Create(conversation).Error
	return err
}

// DeleteConversation 删除消息会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) DeleteConversation(ID string) (err error) {
	err = global.GVA_DB.Delete(&octopus.Conversation{}, "id = ?", ID).Error
	return err
}

// DeleteConversationByIds 批量删除消息会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) DeleteConversationByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]octopus.Conversation{}, "id in ?", IDs).Error
	return err
}

// UpdateConversation 更新消息会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) UpdateConversation(conversation octopus.Conversation) (err error) {
	err = global.GVA_DB.Model(&octopus.Conversation{}).Where("id = ?", conversation.ID).Updates(&conversation).Error
	return err
}

// GetConversation 根据ID获取消息会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) GetConversation(ID string) (conversation octopus.Conversation, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&conversation).Error
	return
}

// GetConversationInfoList 分页获取消息会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) GetConversationInfoList(info octopusReq.ConversationSearch) (list []octopus.Conversation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.Conversation{})
	var conversations []octopus.Conversation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AppName != "" {
		db = db.Where("app_name = ?", info.AppName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&conversations).Error
	return conversations, total, err
}
