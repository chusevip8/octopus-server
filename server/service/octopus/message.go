package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "gorm.io/gorm"
)

type MessageService struct {}

// CreateMessage 创建私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) CreateMessage(message *octopus.Message) (err error) {
	err = global.GVA_DB.Create(message).Error
	return err
}

// DeleteMessage 删除私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)DeleteMessage(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.Message{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&octopus.Message{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteMessageByIds 批量删除私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)DeleteMessageByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.Message{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&octopus.Message{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateMessage 更新私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)UpdateMessage(message octopus.Message) (err error) {
	err = global.GVA_DB.Model(&octopus.Message{}).Where("id = ?",message.ID).Updates(&message).Error
	return err
}

// GetMessage 根据ID获取私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)GetMessage(ID string) (message octopus.Message, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&message).Error
	return
}

// GetMessageInfoList 分页获取私信记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)GetMessageInfoList(info octopusReq.MessageSearch) (list []octopus.Message, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&octopus.Message{})
    var messages []octopus.Message
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&messages).Error
	return  messages, total, err
}