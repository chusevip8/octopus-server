package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
)

type ThreadService struct {}

// CreateThread 创建消息组记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService) CreateThread(thread *octopus.Thread) (err error) {
	err = global.GVA_DB.Create(thread).Error
	return err
}

// DeleteThread 删除消息组记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService)DeleteThread(ID string) (err error) {
	err = global.GVA_DB.Delete(&octopus.Thread{},"id = ?",ID).Error
	return err
}

// DeleteThreadByIds 批量删除消息组记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService)DeleteThreadByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]octopus.Thread{},"id in ?",IDs).Error
	return err
}

// UpdateThread 更新消息组记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService)UpdateThread(thread octopus.Thread) (err error) {
	err = global.GVA_DB.Model(&octopus.Thread{}).Where("id = ?",thread.ID).Updates(&thread).Error
	return err
}

// GetThread 根据ID获取消息组记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService)GetThread(ID string) (thread octopus.Thread, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&thread).Error
	return
}

// GetThreadInfoList 分页获取消息组记录
// Author [piexlmax](https://github.com/piexlmax)
func (threadService *ThreadService)GetThreadInfoList(info octopusReq.ThreadSearch) (list []octopus.Thread, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&octopus.Thread{})
    var threads []octopus.Thread
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Sender != "" {
        db = db.Where("sender LIKE ?","%"+ info.Sender+"%")
    }
    if info.Receiver != "" {
        db = db.Where("receiver LIKE ?","%"+ info.Receiver+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&threads).Error
	return  threads, total, err
}