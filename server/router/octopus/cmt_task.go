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
		cmtTaskRouter.DELETE("deleteCmtTask", cmtTaskApi.DeleteCmtTask)         // 删除任务
		//taskRouter.DELETE("deleteTaskByIds", taskApi.DeleteTaskByIds) // 批量删除任务
	}
	{
		cmtTaskRouterWithoutAuth.POST("uploadComment", cmtTaskApi.UploadComment)
	}
}
