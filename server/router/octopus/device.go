package octopus

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

func (d *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	deviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	deviceRouterWithoutRecord := Router.Group("device")
	deviceRouterWithoutAuth := PublicRouter.Group("device")

	var deviceApi = v1.ApiGroupApp.OctopusApiGroup.DeviceApi
	{
		deviceRouter.DELETE("deleteDevice", deviceApi.DeleteDevice)
		deviceRouter.DELETE("deleteDeviceByIds", deviceApi.DeleteDeviceByIds)
	}
	{
		deviceRouterWithoutRecord.GET("getDeviceList", deviceApi.GetDeviceList)
	}
	{
		deviceRouterWithoutAuth.POST("registerDevice", deviceApi.RegisterDevice)
	}
}
