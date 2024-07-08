package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "gorm.io/gorm"
)

type CmtTaskOptionService struct {}

// CreateCmtTaskOption 创建评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskOptionService *CmtTaskOptionService) CreateCmtTaskOption(cmtTaskOption *octopus.CmtTaskOption) (err error) {
	err = global.GVA_DB.Create(cmtTaskOption).Error
	return err
}

// DeleteCmtTaskOption 删除评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskOptionService *CmtTaskOptionService)DeleteCmtTaskOption(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.CmtTaskOption{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&octopus.CmtTaskOption{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCmtTaskOptionByIds 批量删除评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskOptionService *CmtTaskOptionService)DeleteCmtTaskOptionByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.CmtTaskOption{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&octopus.CmtTaskOption{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCmtTaskOption 更新评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskOptionService *CmtTaskOptionService)UpdateCmtTaskOption(cmtTaskOption octopus.CmtTaskOption) (err error) {
	err = global.GVA_DB.Model(&octopus.CmtTaskOption{}).Where("id = ?",cmtTaskOption.ID).Updates(&cmtTaskOption).Error
	return err
}

// GetCmtTaskOption 根据ID获取评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskOptionService *CmtTaskOptionService)GetCmtTaskOption(ID string) (cmtTaskOption octopus.CmtTaskOption, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmtTaskOption).Error
	return
}

// GetCmtTaskOptionInfoList 分页获取评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskOptionService *CmtTaskOptionService)GetCmtTaskOptionInfoList(info octopusReq.CmtTaskOptionSearch) (list []octopus.CmtTaskOption, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&octopus.CmtTaskOption{})
    var cmtTaskOptions []octopus.CmtTaskOption
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TaskTitle != "" {
        db = db.Where("task_title LIKE ?","%"+ info.TaskTitle+"%")
    }
    if info.Keyword != "" {
        db = db.Where("keyword LIKE ?","%"+ info.Keyword+"%")
    }
    if info.Commenter != "" {
        db = db.Where("commenter LIKE ?","%"+ info.Commenter+"%")
    }
    if info.CommenterId != "" {
        db = db.Where("commenter_id LIKE ?","%"+ info.CommenterId+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&cmtTaskOptions).Error
	return  cmtTaskOptions, total, err
}