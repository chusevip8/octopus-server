package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DYCmtTaskRouter struct {
}

// InitDYCmtTaskRouter 初始化 抖音评论任务 路由信息
func (s *DYCmtTaskRouter) InitDYCmtTaskRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	dyCmtTaskRouter := Router.Group("dyCmtTask").Use(middleware.OperationRecord())
	dyCmtTaskRouterWithoutRecord := Router.Group("dyCmtTask")
	dyCmtTaskRouterWithoutAuth := PublicRouter.Group("dyCmtTask")

	var dyCmtTaskApi = v1.ApiGroupApp.OctopusApiGroup.DYCmtTaskApi
	{
		dyCmtTaskRouter.POST("createDYCmtTask", dyCmtTaskApi.CreateDYCmtTask)   // 新建抖音评论任务
		dyCmtTaskRouter.DELETE("deleteDYCmtTask", dyCmtTaskApi.DeleteDYCmtTask) // 删除抖音评论任务
		dyCmtTaskRouter.DELETE("deleteDYCmtTaskByIds", dyCmtTaskApi.DeleteDYCmtTaskByIds) // 批量删除抖音评论任务
		dyCmtTaskRouter.PUT("updateDYCmtTask", dyCmtTaskApi.UpdateDYCmtTask)    // 更新抖音评论任务
	}
	{
		dyCmtTaskRouterWithoutRecord.GET("findDYCmtTask", dyCmtTaskApi.FindDYCmtTask)        // 根据ID获取抖音评论任务
		dyCmtTaskRouterWithoutRecord.GET("getDYCmtTaskList", dyCmtTaskApi.GetDYCmtTaskList)  // 获取抖音评论任务列表
	}
	{
	    dyCmtTaskRouterWithoutAuth.GET("getDYCmtTaskPublic", dyCmtTaskApi.GetDYCmtTaskPublic)  // 获取抖音评论任务列表
	}
}
