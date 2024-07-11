package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TaskRouter struct{}

// InitTaskRouter 初始化 任务 路由信息
func (s *TaskRouter) InitTaskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	taskRouter := Router.Group("task").Use(middleware.OperationRecord())
	taskRouterWithoutRecord := Router.Group("task")
	taskRouterWithoutAuth := PublicRouter.Group("task")

	var taskApi = v1.ApiGroupApp.OctopusApiGroup.TaskApi
	{
		taskRouter.POST("createTask", taskApi.CreateTask)             // 新建任务
		taskRouter.DELETE("deleteTask", taskApi.DeleteTask)           // 删除任务
		taskRouter.DELETE("deleteTaskByIds", taskApi.DeleteTaskByIds) // 批量删除任务
		taskRouter.PUT("updateTask", taskApi.UpdateTask)              // 更新任务
	}
	{
		taskRouterWithoutRecord.GET("findTaskByDeviceId", taskApi.FindTaskByDeviceId) // 根据Device ID获取任务
		taskRouterWithoutRecord.GET("findTask", taskApi.FindTask)                     // 根据ID获取任务
		taskRouterWithoutRecord.GET("getTaskList", taskApi.GetTaskList)               // 获取任务列表
	}
	{
		taskRouterWithoutAuth.GET("getTaskPublic", taskApi.GetTaskPublic) // 获取任务列表
	}
}
