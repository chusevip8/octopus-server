package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmtTaskMgrRouter struct {
}

// InitCmtTaskMgrRouter 初始化 评论任务管理 路由信息
func (s *CmtTaskMgrRouter) InitCmtTaskMgrRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	cmtTaskMgrRouter := Router.Group("cmtTaskMgr").Use(middleware.OperationRecord())
	cmtTaskMgrRouterWithoutRecord := Router.Group("cmtTaskMgr")
	cmtTaskMgrRouterWithoutAuth := PublicRouter.Group("cmtTaskMgr")

	var cmtTaskMgrApi = v1.ApiGroupApp.OctopusApiGroup.CmtTaskMgrApi
	{
		cmtTaskMgrRouter.POST("createCmtTaskMgr", cmtTaskMgrApi.CreateCmtTaskMgr)   // 新建评论任务管理
		cmtTaskMgrRouter.DELETE("deleteCmtTaskMgr", cmtTaskMgrApi.DeleteCmtTaskMgr) // 删除评论任务管理
		cmtTaskMgrRouter.DELETE("deleteCmtTaskMgrByIds", cmtTaskMgrApi.DeleteCmtTaskMgrByIds) // 批量删除评论任务管理
		cmtTaskMgrRouter.PUT("updateCmtTaskMgr", cmtTaskMgrApi.UpdateCmtTaskMgr)    // 更新评论任务管理
	}
	{
		cmtTaskMgrRouterWithoutRecord.GET("findCmtTaskMgr", cmtTaskMgrApi.FindCmtTaskMgr)        // 根据ID获取评论任务管理
		cmtTaskMgrRouterWithoutRecord.GET("getCmtTaskMgrList", cmtTaskMgrApi.GetCmtTaskMgrList)  // 获取评论任务管理列表
	}
	{
	    cmtTaskMgrRouterWithoutAuth.GET("getCmtTaskMgrPublic", cmtTaskMgrApi.GetCmtTaskMgrPublic)  // 获取评论任务管理列表
	}
}
