package octopus

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type MsgTaskRouter struct{}

func (s *MsgTaskRouter) InitMsgTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	//msgTaskRouter := Router.Group("msgTask").Use(middleware.OperationRecord())
	msgTaskRouterWithoutAuth := PublicRouter.Group("msgTask")
	var msgTaskApi = v1.ApiGroupApp.OctopusApiGroup.MsgTaskApi
	{
		msgTaskRouterWithoutAuth.POST("uploadMessage", msgTaskApi.UploadMessage)
	}
}
