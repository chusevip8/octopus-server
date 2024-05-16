package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type DeviceService struct {
}

func (deviceService *DeviceService) CreateDevice(device *octopus.Device) (err error) {
	err = global.GVA_DB.Create(device).Error
	return err
}

func (deviceService *DeviceService) DeleteDevice(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.Device{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.Device{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (deviceService *DeviceService) DeleteDeviceByIds(IDs []string, deletedBy uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.Device{}).Where("id in ?", IDs).Update("deletedBy", deletedBy).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.Device{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (deviceService *DeviceService) UpdateDevice(device octopus.Device) (err error) {
	err = global.GVA_DB.Model(&octopus.Device{}).Where("id = ?", device.ID).Updates(&device).Error
	return err
}

func (deviceService *DeviceService) GetDevice(ID string) (device octopus.Device, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&device).Error
	return
}

func (deviceService *DeviceService) GetDeviceInfoList(info octopusReq.DeviceSearch) (list []octopus.Device, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.Device{})
	var devices []octopus.Device

	if info.Number != "" {
		db = db.Where("number = ?", info.Number)
	}
	if info.Note != "" {
		db = db.Where("note LIKE ?", "%"+info.Note+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	if info.NickName != "" {
		subQuery := global.GVA_DB.Model(&system.SysUser{}).Select("username").Where("nick_name = ?", info.NickName)
		db = db.Where("username IN (?)", subQuery)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&devices).Error
	return devices, total, err
}
