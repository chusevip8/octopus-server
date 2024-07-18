package octopus

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IntervalTaskRouter struct{}

func (s *IntervalTaskRouter) InitIntervalTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	intervalTaskRouter := Router.Group("intervalTask").Use(middleware.OperationRecord())
	var intervalTaskApi = v1.ApiGroupApp.OctopusApiGroup.IntervalTaskApi

	{
		intervalTaskRouter.POST("createIntervalTask", intervalTaskApi.CreateIntervalTask) // 新建任务
	}

}
