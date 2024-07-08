package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmtThreadRouter struct {}

// InitCmtThreadRouter 初始化 评论会话 路由信息
func (s *CmtThreadRouter) InitCmtThreadRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	cmtThreadRouter := Router.Group("cmtThread").Use(middleware.OperationRecord())
	cmtThreadRouterWithoutRecord := Router.Group("cmtThread")
	cmtThreadRouterWithoutAuth := PublicRouter.Group("cmtThread")

	var cmtThreadApi = v1.ApiGroupApp.OctopusApiGroup.CmtThreadApi
	{
		cmtThreadRouter.POST("createCmtThread", cmtThreadApi.CreateCmtThread)   // 新建评论会话
		cmtThreadRouter.DELETE("deleteCmtThread", cmtThreadApi.DeleteCmtThread) // 删除评论会话
		cmtThreadRouter.DELETE("deleteCmtThreadByIds", cmtThreadApi.DeleteCmtThreadByIds) // 批量删除评论会话
		cmtThreadRouter.PUT("updateCmtThread", cmtThreadApi.UpdateCmtThread)    // 更新评论会话
	}
	{
		cmtThreadRouterWithoutRecord.GET("findCmtThread", cmtThreadApi.FindCmtThread)        // 根据ID获取评论会话
		cmtThreadRouterWithoutRecord.GET("getCmtThreadList", cmtThreadApi.GetCmtThreadList)  // 获取评论会话列表
	}
	{
	    cmtThreadRouterWithoutAuth.GET("getCmtThreadPublic", cmtThreadApi.GetCmtThreadPublic)  // 获取评论会话列表
	}
}
