package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmtTaskSetupRouter struct {}

// InitCmtTaskSetupRouter 初始化 评论任务设置 路由信息
func (s *CmtTaskSetupRouter) InitCmtTaskSetupRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	cmtTaskSetupRouter := Router.Group("cmtTaskSetup").Use(middleware.OperationRecord())
	cmtTaskSetupRouterWithoutRecord := Router.Group("cmtTaskSetup")
	cmtTaskSetupRouterWithoutAuth := PublicRouter.Group("cmtTaskSetup")

	var cmtTaskSetupApi = v1.ApiGroupApp.OctopusApiGroup.CmtTaskSetupApi
	{
		cmtTaskSetupRouter.POST("createCmtTaskSetup", cmtTaskSetupApi.CreateCmtTaskSetup)   // 新建评论任务设置
		cmtTaskSetupRouter.DELETE("deleteCmtTaskSetup", cmtTaskSetupApi.DeleteCmtTaskSetup) // 删除评论任务设置
		cmtTaskSetupRouter.DELETE("deleteCmtTaskSetupByIds", cmtTaskSetupApi.DeleteCmtTaskSetupByIds) // 批量删除评论任务设置
		cmtTaskSetupRouter.PUT("updateCmtTaskSetup", cmtTaskSetupApi.UpdateCmtTaskSetup)    // 更新评论任务设置
	}
	{
		cmtTaskSetupRouterWithoutRecord.GET("findCmtTaskSetup", cmtTaskSetupApi.FindCmtTaskSetup)        // 根据ID获取评论任务设置
		cmtTaskSetupRouterWithoutRecord.GET("getCmtTaskSetupList", cmtTaskSetupApi.GetCmtTaskSetupList)  // 获取评论任务设置列表
	}
	{
	    cmtTaskSetupRouterWithoutAuth.GET("getCmtTaskSetupPublic", cmtTaskSetupApi.GetCmtTaskSetupPublic)  // 获取评论任务设置列表
	}
}
