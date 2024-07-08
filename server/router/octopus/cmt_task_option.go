package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmtTaskOptionRouter struct {}

// InitCmtTaskOptionRouter 初始化 评论任务参数 路由信息
func (s *CmtTaskOptionRouter) InitCmtTaskOptionRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	cmtTaskOptionRouter := Router.Group("cmtTaskOption").Use(middleware.OperationRecord())
	cmtTaskOptionRouterWithoutRecord := Router.Group("cmtTaskOption")
	cmtTaskOptionRouterWithoutAuth := PublicRouter.Group("cmtTaskOption")

	var cmtTaskOptionApi = v1.ApiGroupApp.OctopusApiGroup.CmtTaskOptionApi
	{
		cmtTaskOptionRouter.POST("createCmtTaskOption", cmtTaskOptionApi.CreateCmtTaskOption)   // 新建评论任务参数
		cmtTaskOptionRouter.DELETE("deleteCmtTaskOption", cmtTaskOptionApi.DeleteCmtTaskOption) // 删除评论任务参数
		cmtTaskOptionRouter.DELETE("deleteCmtTaskOptionByIds", cmtTaskOptionApi.DeleteCmtTaskOptionByIds) // 批量删除评论任务参数
		cmtTaskOptionRouter.PUT("updateCmtTaskOption", cmtTaskOptionApi.UpdateCmtTaskOption)    // 更新评论任务参数
	}
	{
		cmtTaskOptionRouterWithoutRecord.GET("findCmtTaskOption", cmtTaskOptionApi.FindCmtTaskOption)        // 根据ID获取评论任务参数
		cmtTaskOptionRouterWithoutRecord.GET("getCmtTaskOptionList", cmtTaskOptionApi.GetCmtTaskOptionList)  // 获取评论任务参数列表
	}
	{
	    cmtTaskOptionRouterWithoutAuth.GET("getCmtTaskOptionPublic", cmtTaskOptionApi.GetCmtTaskOptionPublic)  // 获取评论任务参数列表
	}
}
