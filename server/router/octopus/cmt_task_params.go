package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CmtTaskParamsRouter struct {}

// InitCmtTaskParamsRouter 初始化 评论任务参数 路由信息
func (s *CmtTaskParamsRouter) InitCmtTaskParamsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	cmtTaskParamsRouter := Router.Group("cmtTaskParams").Use(middleware.OperationRecord())
	cmtTaskParamsRouterWithoutRecord := Router.Group("cmtTaskParams")
	cmtTaskParamsRouterWithoutAuth := PublicRouter.Group("cmtTaskParams")

	var cmtTaskParamsApi = v1.ApiGroupApp.OctopusApiGroup.CmtTaskParamsApi
	{
		cmtTaskParamsRouter.POST("createCmtTaskParams", cmtTaskParamsApi.CreateCmtTaskParams)   // 新建评论任务参数
		cmtTaskParamsRouter.DELETE("deleteCmtTaskParams", cmtTaskParamsApi.DeleteCmtTaskParams) // 删除评论任务参数
		cmtTaskParamsRouter.DELETE("deleteCmtTaskParamsByIds", cmtTaskParamsApi.DeleteCmtTaskParamsByIds) // 批量删除评论任务参数
		cmtTaskParamsRouter.PUT("updateCmtTaskParams", cmtTaskParamsApi.UpdateCmtTaskParams)    // 更新评论任务参数
	}
	{
		cmtTaskParamsRouterWithoutRecord.GET("findCmtTaskParams", cmtTaskParamsApi.FindCmtTaskParams)        // 根据ID获取评论任务参数
		cmtTaskParamsRouterWithoutRecord.GET("getCmtTaskParamsList", cmtTaskParamsApi.GetCmtTaskParamsList)  // 获取评论任务参数列表
	}
	{
	    cmtTaskParamsRouterWithoutAuth.GET("getCmtTaskParamsPublic", cmtTaskParamsApi.GetCmtTaskParamsPublic)  // 获取评论任务参数列表
	}
}
