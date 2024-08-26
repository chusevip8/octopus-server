package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsgTaskSetupRouter struct {}

// InitMsgTaskSetupRouter 初始化 私信任务设置 路由信息
func (s *MsgTaskSetupRouter) InitMsgTaskSetupRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	msgTaskSetupRouter := Router.Group("msgTaskSetup").Use(middleware.OperationRecord())
	msgTaskSetupRouterWithoutRecord := Router.Group("msgTaskSetup")
	msgTaskSetupRouterWithoutAuth := PublicRouter.Group("msgTaskSetup")

	var msgTaskSetupApi = v1.ApiGroupApp.OctopusApiGroup.MsgTaskSetupApi
	{
		msgTaskSetupRouter.POST("createMsgTaskSetup", msgTaskSetupApi.CreateMsgTaskSetup)   // 新建私信任务设置
		msgTaskSetupRouter.DELETE("deleteMsgTaskSetup", msgTaskSetupApi.DeleteMsgTaskSetup) // 删除私信任务设置
		msgTaskSetupRouter.DELETE("deleteMsgTaskSetupByIds", msgTaskSetupApi.DeleteMsgTaskSetupByIds) // 批量删除私信任务设置
		msgTaskSetupRouter.PUT("updateMsgTaskSetup", msgTaskSetupApi.UpdateMsgTaskSetup)    // 更新私信任务设置
	}
	{
		msgTaskSetupRouterWithoutRecord.GET("findMsgTaskSetup", msgTaskSetupApi.FindMsgTaskSetup)        // 根据ID获取私信任务设置
		msgTaskSetupRouterWithoutRecord.GET("getMsgTaskSetupList", msgTaskSetupApi.GetMsgTaskSetupList)  // 获取私信任务设置列表
	}
	{
	    msgTaskSetupRouterWithoutAuth.GET("getMsgTaskSetupPublic", msgTaskSetupApi.GetMsgTaskSetupPublic)  // 获取私信任务设置列表
	}
}
