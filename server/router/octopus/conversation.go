package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConversationRouter struct {}

// InitConversationRouter 初始化 消息会话 路由信息
func (s *ConversationRouter) InitConversationRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	conversationRouter := Router.Group("conversation").Use(middleware.OperationRecord())
	conversationRouterWithoutRecord := Router.Group("conversation")
	conversationRouterWithoutAuth := PublicRouter.Group("conversation")

	var conversationApi = v1.ApiGroupApp.OctopusApiGroup.ConversationApi
	{
		conversationRouter.POST("createConversation", conversationApi.CreateConversation)   // 新建消息会话
		conversationRouter.DELETE("deleteConversation", conversationApi.DeleteConversation) // 删除消息会话
		conversationRouter.DELETE("deleteConversationByIds", conversationApi.DeleteConversationByIds) // 批量删除消息会话
		conversationRouter.PUT("updateConversation", conversationApi.UpdateConversation)    // 更新消息会话
	}
	{
		conversationRouterWithoutRecord.GET("findConversation", conversationApi.FindConversation)        // 根据ID获取消息会话
		conversationRouterWithoutRecord.GET("getConversationList", conversationApi.GetConversationList)  // 获取消息会话列表
	}
	{
	    conversationRouterWithoutAuth.GET("getConversationPublic", conversationApi.GetConversationPublic)  // 获取消息会话列表
	}
}
