package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskParamsRouter struct {}

// InitTaskParamsRouter 初始化 任务参数 路由信息
func (s *TaskParamsRouter) InitTaskParamsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	taskParamsRouter := Router.Group("taskParams").Use(middleware.OperationRecord())
	taskParamsRouterWithoutRecord := Router.Group("taskParams")
	taskParamsRouterWithoutAuth := PublicRouter.Group("taskParams")

	var taskParamsApi = v1.ApiGroupApp.OctopusApiGroup.TaskParamsApi
	{
		taskParamsRouter.POST("createTaskParams", taskParamsApi.CreateTaskParams)   // 新建任务参数
		taskParamsRouter.DELETE("deleteTaskParams", taskParamsApi.DeleteTaskParams) // 删除任务参数
		taskParamsRouter.DELETE("deleteTaskParamsByIds", taskParamsApi.DeleteTaskParamsByIds) // 批量删除任务参数
		taskParamsRouter.PUT("updateTaskParams", taskParamsApi.UpdateTaskParams)    // 更新任务参数
	}
	{
		taskParamsRouterWithoutRecord.GET("findTaskParams", taskParamsApi.FindTaskParams)        // 根据ID获取任务参数
		taskParamsRouterWithoutRecord.GET("getTaskParamsList", taskParamsApi.GetTaskParamsList)  // 获取任务参数列表
	}
	{
	    taskParamsRouterWithoutAuth.GET("getTaskParamsPublic", taskParamsApi.GetTaskParamsPublic)  // 获取任务参数列表
	}
}
