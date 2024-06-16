package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "gorm.io/gorm"
)

type CmtTaskMgrService struct {
}

// CreateCmtTaskMgr 创建评论任务管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskMgrService *CmtTaskMgrService) CreateCmtTaskMgr(cmtTaskMgr *octopus.CmtTaskMgr) (err error) {
	err = global.GVA_DB.Create(cmtTaskMgr).Error
	return err
}

// DeleteCmtTaskMgr 删除评论任务管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskMgrService *CmtTaskMgrService)DeleteCmtTaskMgr(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.CmtTaskMgr{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&octopus.CmtTaskMgr{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCmtTaskMgrByIds 批量删除评论任务管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskMgrService *CmtTaskMgrService)DeleteCmtTaskMgrByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.CmtTaskMgr{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&octopus.CmtTaskMgr{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCmtTaskMgr 更新评论任务管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskMgrService *CmtTaskMgrService)UpdateCmtTaskMgr(cmtTaskMgr octopus.CmtTaskMgr) (err error) {
	err = global.GVA_DB.Model(&octopus.CmtTaskMgr{}).Where("id = ?",cmtTaskMgr.ID).Updates(&cmtTaskMgr).Error
	return err
}

// GetCmtTaskMgr 根据ID获取评论任务管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskMgrService *CmtTaskMgrService)GetCmtTaskMgr(ID string) (cmtTaskMgr octopus.CmtTaskMgr, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmtTaskMgr).Error
	return
}

// GetCmtTaskMgrInfoList 分页获取评论任务管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskMgrService *CmtTaskMgrService)GetCmtTaskMgrInfoList(info octopusReq.CmtTaskMgrSearch) (list []octopus.CmtTaskMgr, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&octopus.CmtTaskMgr{})
    var cmtTaskMgrs []octopus.CmtTaskMgr
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TaskTitle != "" {
        db = db.Where("task_title LIKE ?","%"+ info.TaskTitle+"%")
    }
    if info.ArticleID != "" {
        db = db.Where("article_id LIKE ?","%"+ info.ArticleID+"%")
    }
    if info.CmtKeyword != "" {
        db = db.Where("cmt_keyword LIKE ?","%"+ info.CmtKeyword+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&cmtTaskMgrs).Error
	return  cmtTaskMgrs, total, err
}