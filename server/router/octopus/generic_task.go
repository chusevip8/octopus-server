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
		genericTaskRouter.POST("createGenericTask", genericTaskApi.CreateGenericTask) // 新建任务
	}
}