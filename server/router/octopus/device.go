package octopus

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

func (d *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	deviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	//deviceRouterWithoutRecord := Router.Group("device")
	//deviceRouterWithoutAuth := PublicRouter.Group("device")

	var deviceApi = v1.ApiGroupApp.OctopusApiGroup.DeviceApi
	{
		deviceRouter.POST("createScript", deviceApi.CreateDevice)
		deviceRouter.DELETE("deleteScript", deviceApi.DeleteDevice)
		deviceRouter.DELETE("deleteScriptByIds", deviceApi.DeleteDeviceByIds)
		deviceRouter.PUT("updateScript", deviceApi.UpdateDevice)
	}
	//{
	//	deviceRouterWithoutRecord.GET("findScript", scriptApi.FindScript)       // 根据ID获取脚本
	//	deviceRouterWithoutRecord.GET("getScriptList", scriptApi.GetScriptList) // 获取脚本列表
	//}
	//{
	//	deviceRouterWithoutAuth.GET("getScriptPublic", scriptApi.GetScriptPublic) // 获取脚本列表
	//}
}
