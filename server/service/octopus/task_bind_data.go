package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type TaskBindDataService struct{}

var taskBindDataServiceApp = new(TaskBindDataService)

// CreateTaskBindData 创建任务数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskBindDataService *TaskBindDataService) CreateTaskBindData(taskBindData *octopus.TaskBindData) (err error) {
	err = global.GVA_DB.Create(taskBindData).Error
	return err
}

// DeleteTaskBindData 删除任务数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskBindDataService *TaskBindDataService) DeleteTaskBindData(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.TaskBindData{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.TaskBindData{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTaskBindDataByIds 批量删除任务数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskBindDataService *TaskBindDataService) DeleteTaskBindDataByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.TaskBindData{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.TaskBindData{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTaskBindData 更新任务数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskBindDataService *TaskBindDataService) UpdateTaskBindData(taskBindData octopus.TaskBindData) (err error) {
	err = global.GVA_DB.Model(&octopus.TaskBindData{}).Where("id = ?", taskBindData.ID).Updates(&taskBindData).Error
	return err
}

// GetTaskBindData 根据ID获取任务数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskBindDataService *TaskBindDataService) GetTaskBindData(ID string) (taskBindData octopus.TaskBindData, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&taskBindData).Error
	return
}

// GetTaskBindDataInfoList 分页获取任务数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskBindDataService *TaskBindDataService) GetTaskBindDataInfoList(info octopusReq.TaskBindDataSearch) (list []octopus.TaskBindData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.TaskBindData{})
	var taskBindDatas []octopus.TaskBindData
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&taskBindDatas).Error
	return taskBindDatas, total, err
}
