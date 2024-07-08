package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmtConversationRouter struct {}

// InitCmtConversationRouter 初始化 评论会话记录 路由信息
func (s *CmtConversationRouter) InitCmtConversationRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	cmtConversationRouter := Router.Group("cmtConversation").Use(middleware.OperationRecord())
	cmtConversationRouterWithoutRecord := Router.Group("cmtConversation")
	cmtConversationRouterWithoutAuth := PublicRouter.Group("cmtConversation")

	var cmtConversationApi = v1.ApiGroupApp.OctopusApiGroup.CmtConversationApi
	{
		cmtConversationRouter.POST("createCmtConversation", cmtConversationApi.CreateCmtConversation)   // 新建评论会话记录
		cmtConversationRouter.DELETE("deleteCmtConversation", cmtConversationApi.DeleteCmtConversation) // 删除评论会话记录
		cmtConversationRouter.DELETE("deleteCmtConversationByIds", cmtConversationApi.DeleteCmtConversationByIds) // 批量删除评论会话记录
		cmtConversationRouter.PUT("updateCmtConversation", cmtConversationApi.UpdateCmtConversation)    // 更新评论会话记录
	}
	{
		cmtConversationRouterWithoutRecord.GET("findCmtConversation", cmtConversationApi.FindCmtConversation)        // 根据ID获取评论会话记录
		cmtConversationRouterWithoutRecord.GET("getCmtConversationList", cmtConversationApi.GetCmtConversationList)  // 获取评论会话记录列表
	}
	{
	    cmtConversationRouterWithoutAuth.GET("getCmtConversationPublic", cmtConversationApi.GetCmtConversationPublic)  // 获取评论会话记录列表
	}
}
