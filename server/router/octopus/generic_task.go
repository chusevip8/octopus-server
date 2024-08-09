package octopus

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GenericTaskRouter struct{}

func (s *GenericTaskRouter) InitGenericTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	genericTaskRouter := Router.Group("genericTask").Use(middleware.OperationRecord())
	var genericTaskApi = v1.ApiGroupApp.OctopusApiGroup.GenericTaskApi

	{
		genericTaskRouter.POST("createGenericTask", genericTaskApi.CreateGenericTask)   // 新建任务
		genericTaskRouter.POST("bindTaskData", genericTaskApi.BindTaskData)             // 绑定任务数据
		genericTaskRouter.POST("startGenericTasks", genericTaskApi.StartGenericTasks)   // 运行所有任务
		genericTaskRouter.GET("stopGenericTask", genericTaskApi.StopGenericTask)        // 停止单个任务
		genericTaskRouter.POST("stopGenericTasks", genericTaskApi.StopGenericTasks)     // 停止所有任务
		genericTaskRouter.POST("deleteGenericTasks", genericTaskApi.DeleteGenericTasks) // 删除所有任务
		genericTaskRouter.DELETE("deleteGenericTask", genericTaskApi.DeleteGenericTask) // 删除任务
	}
}
