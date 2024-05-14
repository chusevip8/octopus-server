package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScriptRouter struct {
}

// InitScriptRouter 初始化 脚本 路由信息
func (s *ScriptRouter) InitScriptRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	scriptRouter := Router.Group("script").Use(middleware.OperationRecord())
	scriptRouterWithoutRecord := Router.Group("script")
	scriptRouterWithoutAuth := PublicRouter.Group("script")

	var scriptApi = v1.ApiGroupApp.OctopusApiGroup.ScriptApi
	{
		scriptRouter.POST("createScript", scriptApi.CreateScript)   // 新建脚本
		scriptRouter.DELETE("deleteScript", scriptApi.DeleteScript) // 删除脚本
		scriptRouter.DELETE("deleteScriptByIds", scriptApi.DeleteScriptByIds) // 批量删除脚本
		scriptRouter.PUT("updateScript", scriptApi.UpdateScript)    // 更新脚本
	}
	{
		scriptRouterWithoutRecord.GET("findScript", scriptApi.FindScript)        // 根据ID获取脚本
		scriptRouterWithoutRecord.GET("getScriptList", scriptApi.GetScriptList)  // 获取脚本列表
	}
	{
	    scriptRouterWithoutAuth.GET("getScriptPublic", scriptApi.GetScriptPublic)  // 获取脚本列表
	}
}
