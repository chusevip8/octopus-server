package octopus

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmtTaskRouter struct{}

func (s *CmtTaskRouter) InitCmtTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cmtTaskRouter := Router.Group("cmtTask").Use(middleware.OperationRecord())
	//cmtTaskRouterWithoutRecord := Router.Group("cmtTask")
	cmtTaskRouterWithoutAuth := PublicRouter.Group("cmtTask")
	var cmtTaskApi = v1.ApiGroupApp.OctopusApiGroup.CmtTaskApi

	{
		cmtTaskRouter.POST("createFindCmtTask", cmtTaskApi.CreateFindCmtTask)   // 新建任务
		cmtTaskRouter.POST("createWriteCmtTask", cmtTaskApi.CreateWriteCmtTask) // 新建任务
		cmtTaskRouter.DELETE("deleteCmtTask", cmtTaskApi.DeleteCmtTask)         // 删除单个任务
		cmtTaskRouter.GET("stopCmtTask", cmtTaskApi.StopCmtTask)                // 停止单个任务
		cmtTaskRouter.POST("stopCmtTasks", cmtTaskApi.StopCmtTasks)             // 停止所有任务
		cmtTaskRouter.POST("deleteCmtTasks", cmtTaskApi.DeleteCmtTasks)         // 删除所有任务
	}
	{
		cmtTaskRouterWithoutAuth.POST("uploadFindComment", cmtTaskApi.UploadFindComment)
	}
}
