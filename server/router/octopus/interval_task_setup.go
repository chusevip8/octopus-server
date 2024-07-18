package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IntervalTaskSetupRouter struct {}

// InitIntervalTaskSetupRouter 初始化 间隔任务设置 路由信息
func (s *IntervalTaskSetupRouter) InitIntervalTaskSetupRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	intervalTaskSetupRouter := Router.Group("intervalTaskSetup").Use(middleware.OperationRecord())
	intervalTaskSetupRouterWithoutRecord := Router.Group("intervalTaskSetup")
	intervalTaskSetupRouterWithoutAuth := PublicRouter.Group("intervalTaskSetup")

	var intervalTaskSetupApi = v1.ApiGroupApp.OctopusApiGroup.IntervalTaskSetupApi
	{
		intervalTaskSetupRouter.POST("createIntervalTaskSetup", intervalTaskSetupApi.CreateIntervalTaskSetup)   // 新建间隔任务设置
		intervalTaskSetupRouter.DELETE("deleteIntervalTaskSetup", intervalTaskSetupApi.DeleteIntervalTaskSetup) // 删除间隔任务设置
		intervalTaskSetupRouter.DELETE("deleteIntervalTaskSetupByIds", intervalTaskSetupApi.DeleteIntervalTaskSetupByIds) // 批量删除间隔任务设置
		intervalTaskSetupRouter.PUT("updateIntervalTaskSetup", intervalTaskSetupApi.UpdateIntervalTaskSetup)    // 更新间隔任务设置
	}
	{
		intervalTaskSetupRouterWithoutRecord.GET("findIntervalTaskSetup", intervalTaskSetupApi.FindIntervalTaskSetup)        // 根据ID获取间隔任务设置
		intervalTaskSetupRouterWithoutRecord.GET("getIntervalTaskSetupList", intervalTaskSetupApi.GetIntervalTaskSetupList)  // 获取间隔任务设置列表
	}
	{
	    intervalTaskSetupRouterWithoutAuth.GET("getIntervalTaskSetupPublic", intervalTaskSetupApi.GetIntervalTaskSetupPublic)  // 获取间隔任务设置列表
	}
}
