package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GenericTaskSetupRouter struct {}

// InitGenericTaskSetupRouter 初始化 通用任务设置 路由信息
func (s *GenericTaskSetupRouter) InitGenericTaskSetupRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	genericTaskSetupRouter := Router.Group("genericTaskSetup").Use(middleware.OperationRecord())
	genericTaskSetupRouterWithoutRecord := Router.Group("genericTaskSetup")
	genericTaskSetupRouterWithoutAuth := PublicRouter.Group("genericTaskSetup")

	var genericTaskSetupApi = v1.ApiGroupApp.OctopusApiGroup.GenericTaskSetupApi
	{
		genericTaskSetupRouter.POST("createGenericTaskSetup", genericTaskSetupApi.CreateGenericTaskSetup)   // 新建通用任务设置
		genericTaskSetupRouter.DELETE("deleteGenericTaskSetup", genericTaskSetupApi.DeleteGenericTaskSetup) // 删除通用任务设置
		genericTaskSetupRouter.DELETE("deleteGenericTaskSetupByIds", genericTaskSetupApi.DeleteGenericTaskSetupByIds) // 批量删除通用任务设置
		genericTaskSetupRouter.PUT("updateGenericTaskSetup", genericTaskSetupApi.UpdateGenericTaskSetup)    // 更新通用任务设置
	}
	{
		genericTaskSetupRouterWithoutRecord.GET("findGenericTaskSetup", genericTaskSetupApi.FindGenericTaskSetup)        // 根据ID获取通用任务设置
		genericTaskSetupRouterWithoutRecord.GET("getGenericTaskSetupList", genericTaskSetupApi.GetGenericTaskSetupList)  // 获取通用任务设置列表
	}
	{
	    genericTaskSetupRouterWithoutAuth.GET("getGenericTaskSetupPublic", genericTaskSetupApi.GetGenericTaskSetupPublic)  // 获取通用任务设置列表
	}
}
