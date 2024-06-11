package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "gorm.io/gorm"
)

type DYCmtTaskService struct {
}

// CreateDYCmtTask 创建抖音评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (dyCmtTaskService *DYCmtTaskService) CreateDYCmtTask(dyCmtTask *octopus.DYCmtTask) (err error) {
	err = global.GVA_DB.Create(dyCmtTask).Error
	return err
}

// DeleteDYCmtTask 删除抖音评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (dyCmtTaskService *DYCmtTaskService)DeleteDYCmtTask(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.DYCmtTask{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&octopus.DYCmtTask{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteDYCmtTaskByIds 批量删除抖音评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (dyCmtTaskService *DYCmtTaskService)DeleteDYCmtTaskByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.DYCmtTask{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&octopus.DYCmtTask{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateDYCmtTask 更新抖音评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (dyCmtTaskService *DYCmtTaskService)UpdateDYCmtTask(dyCmtTask octopus.DYCmtTask) (err error) {
	err = global.GVA_DB.Model(&octopus.DYCmtTask{}).Where("id = ?",dyCmtTask.ID).Updates(&dyCmtTask).Error
	return err
}

// GetDYCmtTask 根据ID获取抖音评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (dyCmtTaskService *DYCmtTaskService)GetDYCmtTask(ID string) (dyCmtTask octopus.DYCmtTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&dyCmtTask).Error
	return
}

// GetDYCmtTaskInfoList 分页获取抖音评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (dyCmtTaskService *DYCmtTaskService)GetDYCmtTaskInfoList(info octopusReq.DYCmtTaskSearch) (list []octopus.DYCmtTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&octopus.DYCmtTask{})
    var dyCmtTasks []octopus.DYCmtTask
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.VideoTitle != "" {
        db = db.Where("video_title LIKE ?","%"+ info.VideoTitle+"%")
    }
    if info.Keyword != "" {
        db = db.Where("keyword LIKE ?","%"+ info.Keyword+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&dyCmtTasks).Error
	return  dyCmtTasks, total, err
}