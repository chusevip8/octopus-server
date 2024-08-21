package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type CmtThreadService struct{}

var CmtThreadServiceApp = new(CmtThreadService)

// CreateCmtThread 创建评论会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtThreadService *CmtThreadService) CreateCmtThread(cmtThread *octopus.CmtThread) (err error) {
	err = global.GVA_DB.Create(cmtThread).Error
	return err
}

// DeleteCmtThread 删除评论会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtThreadService *CmtThreadService) DeleteCmtThread(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.CmtThread{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.CmtThread{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCmtThreadByIds 批量删除评论会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtThreadService *CmtThreadService) DeleteCmtThreadByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.CmtThread{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.CmtThread{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCmtThread 更新评论会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtThreadService *CmtThreadService) UpdateCmtThread(cmtThread octopus.CmtThread) (err error) {
	err = global.GVA_DB.Model(&octopus.CmtThread{}).Where("id = ?", cmtThread.ID).Updates(&cmtThread).Error
	return err
}

// GetCmtThread 根据ID获取评论会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtThreadService *CmtThreadService) GetCmtThread(ID string) (cmtThread octopus.CmtThread, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmtThread).Error
	return
}

// GetCmtThreadInfoList 分页获取评论会话记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtThreadService *CmtThreadService) GetCmtThreadInfoList(info octopusReq.CmtThreadSearch) (list []octopus.CmtThread, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.CmtThread{})

	if !isAdmin(info.CreatedBy) {
		db = db.Where("oct_cmt_thread.created_by = ?", info.CreatedBy)
	}

	var cmtThreads []octopus.CmtThread
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.AppName != "" {
		db = db.Where("app_name = ?", info.AppName)
	}
	if info.Poster != "" {
		db = db.Where("poster LIKE ?", "%"+info.Poster+"%")
	}
	if info.PostTitle != "" {
		db = db.Where("post_title LIKE ?", "%"+info.PostTitle+"%")
	}
	if info.PostDesc != "" {
		db = db.Where("post_desc LIKE ?", "%"+info.PostDesc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("oct_cmt_thread.*, COUNT(oct_comment.id) AS unread_count").
		Joins("LEFT JOIN oct_cmt_conversation ON oct_cmt_conversation.thread_id = oct_cmt_thread.id").
		Joins("LEFT JOIN oct_comment ON oct_comment.conversation_id = oct_cmt_conversation.id AND oct_comment.unread = 1").
		Group("oct_cmt_thread.id").
		Order("unread_count DESC").
		Find(&cmtThreads).Error

	return cmtThreads, total, err
}
