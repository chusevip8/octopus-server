package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskBindDataRouter struct {}

// InitTaskBindDataRouter 初始化 任务数据 路由信息
func (s *TaskBindDataRouter) InitTaskBindDataRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	taskBindDataRouter := Router.Group("taskBindData").Use(middleware.OperationRecord())
	taskBindDataRouterWithoutRecord := Router.Group("taskBindData")
	taskBindDataRouterWithoutAuth := PublicRouter.Group("taskBindData")

	var taskBindDataApi = v1.ApiGroupApp.OctopusApiGroup.TaskBindDataApi
	{
		taskBindDataRouter.POST("createTaskBindData", taskBindDataApi.CreateTaskBindData)   // 新建任务数据
		taskBindDataRouter.DELETE("deleteTaskBindData", taskBindDataApi.DeleteTaskBindData) // 删除任务数据
		taskBindDataRouter.DELETE("deleteTaskBindDataByIds", taskBindDataApi.DeleteTaskBindDataByIds) // 批量删除任务数据
		taskBindDataRouter.PUT("updateTaskBindData", taskBindDataApi.UpdateTaskBindData)    // 更新任务数据
	}
	{
		taskBindDataRouterWithoutRecord.GET("findTaskBindData", taskBindDataApi.FindTaskBindData)        // 根据ID获取任务数据
		taskBindDataRouterWithoutRecord.GET("getTaskBindDataList", taskBindDataApi.GetTaskBindDataList)  // 获取任务数据列表
	}
	{
	    taskBindDataRouterWithoutAuth.GET("getTaskBindDataPublic", taskBindDataApi.GetTaskBindDataPublic)  // 获取任务数据列表
	}
}
