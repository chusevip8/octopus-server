package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type CmtConversationService struct{}

var CmtConversationServiceApp = new(CmtConversationService)

// CreateCmtConversation 创建评论会话记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtConversationService *CmtConversationService) CreateCmtConversation(cmtConversation *octopus.CmtConversation) (err error) {
	err = global.GVA_DB.Create(cmtConversation).Error
	return err
}

// DeleteCmtConversation 删除评论会话记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtConversationService *CmtConversationService) DeleteCmtConversation(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.CmtConversation{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.CmtConversation{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCmtConversationByIds 批量删除评论会话记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtConversationService *CmtConversationService) DeleteCmtConversationByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.CmtConversation{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.CmtConversation{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCmtConversation 更新评论会话记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtConversationService *CmtConversationService) UpdateCmtConversation(cmtConversation octopus.CmtConversation) (err error) {
	err = global.GVA_DB.Model(&octopus.CmtConversation{}).Where("id = ?", cmtConversation.ID).Updates(&cmtConversation).Error
	return err
}

// GetCmtConversation 根据ID获取评论会话记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtConversationService *CmtConversationService) GetCmtConversation(ID string) (cmtConversation octopus.CmtConversation, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmtConversation).Error
	return
}

// GetCmtConversationInfoList 分页获取评论会话记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtConversationService *CmtConversationService) GetCmtConversationInfoList(info octopusReq.CmtConversationSearch) (list []octopus.CmtConversation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.CmtConversation{})
	var cmtConversations []octopus.CmtConversation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ThreadId != "" {
		db = db.Where("thread_id = ?", info.ThreadId)
	}
	if info.Commenter != "" {
		db = db.Where("commenter LIKE ?", "%"+info.Commenter+"%")
	}
	if info.CommentReplier != "" {
		db = db.Where("comment_replier LIKE ?", "%"+info.CommentReplier+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cmtConversations).Error
	return cmtConversations, total, err
}
