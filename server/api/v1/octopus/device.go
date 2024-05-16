package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeviceApi struct {
}

var deviceService = service.ServiceGroupApp.OctopusServiceGroup.DeviceService

func (deviceApi *DeviceApi) CreateDevice(c *gin.Context) {
	var device octopus.Device
	err := c.ShouldBindJSON(&device)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	device.CreatedBy = utils.GetUserID(c)

	if err := deviceService.CreateDevice(&device); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (deviceApi *DeviceApi) DeleteDevice(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := deviceService.DeleteDevice(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (deviceApi *DeviceApi) DeleteDeviceByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := deviceService.DeleteDeviceByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (deviceApi *DeviceApi) UpdateDevice(c *gin.Context) {
	var device octopus.Device
	err := c.ShouldBindJSON(&device)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	device.UpdatedBy = utils.GetUserID(c)

	if err := deviceService.UpdateDevice(device); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
