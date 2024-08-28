package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type MsgConversationService struct{}

var MsgConversationServiceApp = new(MsgConversationService)

// CreateMsgConversation 创建私信会话纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgConversationService *MsgConversationService) CreateMsgConversation(msgConversation *octopus.MsgConversation) (err error) {
	err = global.GVA_DB.Create(msgConversation).Error
	return err
}

// DeleteMsgConversaton 删除私信会话纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgConversationService *MsgConversationService) DeleteMsgConversation(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.MsgConversation{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.MsgConversation{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMsgConversationByIds 批量删除私信会话纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgConversationService *MsgConversationService) DeleteMsgConversationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.MsgConversation{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.MsgConversation{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMsgConversation 更新私信会话纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgConversationService *MsgConversationService) UpdateMsgConversation(msgConversation octopus.MsgConversation) (err error) {
	err = global.GVA_DB.Model(&octopus.MsgConversation{}).Where("id = ?", msgConversation.ID).Updates(&msgConversation).Error
	return err
}

// GetMsgConversation 根据ID获取私信会话纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgConversationService *MsgConversationService) GetMsgConversation(ID string) (msgConversation octopus.MsgConversation, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&msgConversation).Error
	return
}

// GetMsgConversationInfoList 分页获取私信会话纪录记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgConversationService *MsgConversationService) GetMsgConversationInfoList(info octopusReq.MsgConversationSearch) (list []octopus.MsgConversation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.MsgConversation{})

	if !isAdmin(info.CreatedBy) {
		db = db.Where("oct_msg_conversation.created_by = ?", info.CreatedBy)
	}

	var msgConversations []octopus.MsgConversation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.AppName != "" {
		db = db.Where("app_name = ?", info.AppName)
	}
	if info.Sender != "" {
		db = db.Where("sender LIKE ?", "%"+info.Sender+"%")
	}
	if info.Receiver != "" {
		db = db.Where("receiver LIKE ?", "%"+info.Receiver+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("oct_msg_conversation.*, COUNT(oct_message.id) AS unread_count").
		Joins("LEFT JOIN oct_message ON oct_message.conversation_id = oct_msg_conversation.id AND oct_message.unread = 1").
		Group("oct_msg_conversation.id").
		Order("unread_count DESC").
		Find(&msgConversations).Error

	return msgConversations, total, err
}
