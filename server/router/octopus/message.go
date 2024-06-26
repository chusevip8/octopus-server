package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MessageRouter struct {}

// InitMessageRouter 初始化 消息 路由信息
func (s *MessageRouter) InitMessageRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	messageRouter := Router.Group("message").Use(middleware.OperationRecord())
	messageRouterWithoutRecord := Router.Group("message")
	messageRouterWithoutAuth := PublicRouter.Group("message")

	var messageApi = v1.ApiGroupApp.OctopusApiGroup.MessageApi
	{
		messageRouter.POST("createMessage", messageApi.CreateMessage)   // 新建消息
		messageRouter.DELETE("deleteMessage", messageApi.DeleteMessage) // 删除消息
		messageRouter.DELETE("deleteMessageByIds", messageApi.DeleteMessageByIds) // 批量删除消息
		messageRouter.PUT("updateMessage", messageApi.UpdateMessage)    // 更新消息
	}
	{
		messageRouterWithoutRecord.GET("findMessage", messageApi.FindMessage)        // 根据ID获取消息
		messageRouterWithoutRecord.GET("getMessageList", messageApi.GetMessageList)  // 获取消息列表
	}
	{
	    messageRouterWithoutAuth.GET("getMessagePublic", messageApi.GetMessagePublic)  // 获取消息列表
	}
}
