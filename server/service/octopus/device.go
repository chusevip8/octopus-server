package octopus

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceService struct {
}

var DeviceServiceApp = new(DeviceService)

func (deviceService *DeviceService) RegisterDevice(d octopus.Device) (device octopus.Device, err error) {
	var user system.SysUser
	if errors.Is(global.GVA_DB.Model(&system.SysUser{}).Where("username = ?", d.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return device, errors.New("用户未注册")
	}

	device.Username = d.Username
	device.Number = d.Number
	device.Brand = d.Brand
	device.OS = d.OS
	device.Note = d.Note
	device.Status = 2 //表示离线状态
	device.Group = d.Group
	device.CreatedBy = user.ID
	device.LoginToken = uuid.New().String()

	var deviceInter octopus.Device
	if errors.Is(global.GVA_DB.Model(&octopus.Device{}).Where("number = ? AND username = ?", d.Number, d.Username).First(&deviceInter).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		err = global.GVA_DB.Create(&device).Error
		return device, err
	}
	device.ID = deviceInter.ID
	err = global.GVA_DB.Save(&device).Error
	return device, nil
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

func (deviceService *DeviceService) UpdateDeviceStatusById(id uint, status uint) (err error) {
	err = global.GVA_DB.Model(&octopus.Device{}).Where("id = ?", id).Where("status != ?", 3).Update("status", status).Error
	return
}

func (deviceService *DeviceService) GetDeviceByToken(token string) (device octopus.Device, err error) {
	err = global.GVA_DB.Where("login_token = ?", token).Where("status != ?", 3).First(&device).Error
	return
}

func (deviceService *DeviceService) GetReadyDeviceListByUserId(userid uint, deviceGroup string) (list []octopus.Device, err error) {
	var devices []octopus.Device
	db := global.GVA_DB.Joins("JOIN sys_users ON sys_users.username = oct_device.username").
		Where("sys_users.id = ? AND oct_device.status = ?", userid, 1)
	if deviceGroup != "" {
		db = db.Where("oct_device.group_name = ?", deviceGroup)
	}
	err = db.Find(&devices).Error
	return devices, err
}

func (deviceService *DeviceService) GetDeviceInfoList(info octopusReq.DeviceSearch) (list []octopus.Device, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.Device{}).Preload("User")
	var devices []octopus.Device

	if info.Number != "" {
		db = db.Where("number = ?", info.Number)
	}
	if info.Group != "" {
		db = db.Where("group_name = ?", info.Group)
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

func (deviceService *DeviceService) DeviceIsFree(deviceId string) bool {
	var task octopus.Task
	err := global.GVA_DB.Model(&octopus.Task{}).
		Where("device_id = ?", deviceId).
		Where("status = ?", 2).
		First(&task).Error
	if err != nil {
		var device *octopus.Device
		err = global.GVA_DB.Model(&octopus.Device{}).
			Where("id = ?", deviceId).
			Where("status = ?", 1).
			First(&device).Error
		if err != nil {
			return false
		}
		return true
	}
	return false
}
