package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type CmtTaskParamsService struct{}

var CmtTaskParamsServiceApp = new(CmtTaskParamsService)

// CreateCmtTaskParams 创建评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskParamsService *CmtTaskParamsService) CreateCmtTaskParams(cmtTaskParams *octopus.CmtTaskParams) (err error) {
	err = global.GVA_DB.Create(cmtTaskParams).Error
	return err
}

// DeleteCmtTaskParams 删除评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskParamsService *CmtTaskParamsService) DeleteCmtTaskParams(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.CmtTaskParams{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.CmtTaskParams{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCmtTaskParamsByIds 批量删除评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskParamsService *CmtTaskParamsService) DeleteCmtTaskParamsByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.CmtTaskParams{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.CmtTaskParams{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCmtTaskParams 更新评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskParamsService *CmtTaskParamsService) UpdateCmtTaskParams(cmtTaskParams octopus.CmtTaskParams) (err error) {
	err = global.GVA_DB.Model(&octopus.CmtTaskParams{}).Where("id = ?", cmtTaskParams.ID).Updates(&cmtTaskParams).Error
	return err
}

// GetCmtTaskParams 根据ID获取评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskParamsService *CmtTaskParamsService) GetCmtTaskParams(ID string) (cmtTaskParams octopus.CmtTaskParams, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmtTaskParams).Error
	return
}

// GetCmtTaskParamsInfoList 分页获取评论任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskParamsService *CmtTaskParamsService) GetCmtTaskParamsInfoList(info octopusReq.CmtTaskParamsSearch) (list []octopus.CmtTaskParams, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.CmtTaskParams{})
	var cmtTaskParamss []octopus.CmtTaskParams
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Params != "" {
		db = db.Where("params LIKE ?", "%"+info.Params+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cmtTaskParamss).Error
	return cmtTaskParamss, total, err
}
