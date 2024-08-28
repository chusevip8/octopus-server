package octopus

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsgTaskRouter struct{}

func (s *MsgTaskRouter) InitMsgTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	msgTaskRouter := Router.Group("msgTask").Use(middleware.OperationRecord())
	msgTaskRouterWithoutAuth := PublicRouter.Group("msgTask")
	var msgTaskApi = v1.ApiGroupApp.OctopusApiGroup.MsgTaskApi
	{
		msgTaskRouter.POST("createReplyMsgTask", msgTaskApi.CreateReplyMsgTask) // 新建任务
	}

	{
		msgTaskRouterWithoutAuth.POST("uploadMessage", msgTaskApi.UploadMessage)
	}
}
