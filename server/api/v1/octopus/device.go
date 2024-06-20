package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeviceApi struct {
}

var deviceService = service.ServiceGroupApp.OctopusServiceGroup.DeviceService

func (deviceApi *DeviceApi) RegisterDevice(c *gin.Context) {
	var device octopus.Device
	err := c.ShouldBindJSON(&device)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if _, err := deviceService.RegisterDevice(device); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("设备注册失败--"+err.Error(), c)
	} else {
		response.OkWithMessage("设备注册成功", c)
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

func (deviceApi *DeviceApi) GetDeviceList(c *gin.Context) {
	var pageInfo octopusReq.DeviceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := deviceService.GetDeviceInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
