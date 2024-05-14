package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "gorm.io/gorm"
)

type ScriptService struct {
}

// CreateScript 创建脚本记录
// Author [piexlmax](https://github.com/piexlmax)
func (scriptService *ScriptService) CreateScript(script *octopus.Script) (err error) {
	err = global.GVA_DB.Create(script).Error
	return err
}

// DeleteScript 删除脚本记录
// Author [piexlmax](https://github.com/piexlmax)
func (scriptService *ScriptService)DeleteScript(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.Script{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&octopus.Script{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteScriptByIds 批量删除脚本记录
// Author [piexlmax](https://github.com/piexlmax)
func (scriptService *ScriptService)DeleteScriptByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.Script{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&octopus.Script{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateScript 更新脚本记录
// Author [piexlmax](https://github.com/piexlmax)
func (scriptService *ScriptService)UpdateScript(script octopus.Script) (err error) {
	err = global.GVA_DB.Model(&octopus.Script{}).Where("id = ?",script.ID).Updates(&script).Error
	return err
}

// GetScript 根据ID获取脚本记录
// Author [piexlmax](https://github.com/piexlmax)
func (scriptService *ScriptService)GetScript(ID string) (script octopus.Script, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&script).Error
	return
}

// GetScriptInfoList 分页获取脚本记录
// Author [piexlmax](https://github.com/piexlmax)
func (scriptService *ScriptService)GetScriptInfoList(info octopusReq.ScriptSearch) (list []octopus.Script, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&octopus.Script{})
    var scripts []octopus.Script
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Note != "" {
        db = db.Where("note LIKE ?","%"+ info.Note+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&scripts).Error
	return  scripts, total, err
}