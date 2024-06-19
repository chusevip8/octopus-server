package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ExecTaskRouter struct {
}

// InitExecTaskRouter 初始化 执行任务 路由信息
func (s *ExecTaskRouter) InitExecTaskRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	execTaskRouter := Router.Group("execTask").Use(middleware.OperationRecord())
	execTaskRouterWithoutRecord := Router.Group("execTask")
	execTaskRouterWithoutAuth := PublicRouter.Group("execTask")

	var execTaskApi = v1.ApiGroupApp.OctopusApiGroup.ExecTaskApi
	{
		execTaskRouter.POST("createExecTask", execTaskApi.CreateExecTask)   // 新建执行任务
		execTaskRouter.DELETE("deleteExecTask", execTaskApi.DeleteExecTask) // 删除执行任务
		execTaskRouter.DELETE("deleteExecTaskByIds", execTaskApi.DeleteExecTaskByIds) // 批量删除执行任务
		execTaskRouter.PUT("updateExecTask", execTaskApi.UpdateExecTask)    // 更新执行任务
	}
	{
		execTaskRouterWithoutRecord.GET("findExecTask", execTaskApi.FindExecTask)        // 根据ID获取执行任务
		execTaskRouterWithoutRecord.GET("getExecTaskList", execTaskApi.GetExecTaskList)  // 获取执行任务列表
	}
	{
	    execTaskRouterWithoutAuth.GET("getExecTaskPublic", execTaskApi.GetExecTaskPublic)  // 获取执行任务列表
	}
}
