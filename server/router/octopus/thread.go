package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ThreadRouter struct {}

// InitThreadRouter 初始化 消息组 路由信息
func (s *ThreadRouter) InitThreadRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	threadRouter := Router.Group("thread").Use(middleware.OperationRecord())
	threadRouterWithoutRecord := Router.Group("thread")
	threadRouterWithoutAuth := PublicRouter.Group("thread")

	var threadApi = v1.ApiGroupApp.OctopusApiGroup.ThreadApi
	{
		threadRouter.POST("createThread", threadApi.CreateThread)   // 新建消息组
		threadRouter.DELETE("deleteThread", threadApi.DeleteThread) // 删除消息组
		threadRouter.DELETE("deleteThreadByIds", threadApi.DeleteThreadByIds) // 批量删除消息组
		threadRouter.PUT("updateThread", threadApi.UpdateThread)    // 更新消息组
	}
	{
		threadRouterWithoutRecord.GET("findThread", threadApi.FindThread)        // 根据ID获取消息组
		threadRouterWithoutRecord.GET("getThreadList", threadApi.GetThreadList)  // 获取消息组列表
	}
	{
	    threadRouterWithoutAuth.GET("getThreadPublic", threadApi.GetThreadPublic)  // 获取消息组列表
	}
}
