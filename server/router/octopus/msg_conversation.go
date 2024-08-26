package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MsgConversationRouter struct{}

// InitMsgConversationRouter 初始化 私信会话纪录 路由信息
func (s *MsgConversationRouter) InitMsgConversationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	msgConversationRouter := Router.Group("msgConversation").Use(middleware.OperationRecord())
	msgConversationRouterWithoutRecord := Router.Group("msgConversation")
	msgConversationRouterWithoutAuth := PublicRouter.Group("msgConversation")

	var msgConversationApi = v1.ApiGroupApp.OctopusApiGroup.MsgConversationApi
	{
		msgConversationRouter.POST("createMsgConversation", msgConversationApi.CreateMsgConversation)             // 新建私信会话纪录
		msgConversationRouter.DELETE("deleteMsgConversation", msgConversationApi.DeleteMsgConversation)           // 删除私信会话纪录
		msgConversationRouter.DELETE("deleteMsgConversationByIds", msgConversationApi.DeleteMsgConversationByIds) // 批量删除私信会话纪录
		msgConversationRouter.PUT("updateMsgConversation", msgConversationApi.UpdateMsgConversation)              // 更新私信会话纪录
	}
	{
		msgConversationRouterWithoutRecord.GET("findMsgConversation", msgConversationApi.FindMsgConversation)       // 根据ID获取私信会话纪录
		msgConversationRouterWithoutRecord.GET("getMsgConversationList", msgConversationApi.GetMsgConversationList) // 获取私信会话纪录列表
	}
	{
		msgConversationRouterWithoutAuth.GET("getMsgConversationPublic", msgConversationApi.GetMsgConversationPublic) // 获取私信会话纪录列表
	}
}
