package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GenericTaskApi struct {
}

var genericTaskService = service.ServiceGroupApp.OctopusServiceGroup.GenericTaskService

func (genericTaskApi *GenericTaskApi) CreateGenericTask(c *gin.Context) {
	var genericTask octopusReq.GenericTask
	err := c.ShouldBindJSON(&genericTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	genericTask.CreatedBy = utils.GetUserID(c)
	if err := genericTaskService.CreateGenericTask(genericTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (genericTaskApi *GenericTaskApi) BindTaskData(c *gin.Context) {
	var bindTaskData octopusReq.BindTaskData
	err := c.ShouldBindJSON(&bindTaskData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := genericTaskService.BindTaskData(bindTaskData); err != nil {
		global.GVA_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithMessage("绑定失败", c)
	} else {
		response.OkWithMessage("绑定成功", c)
	}
}

func (genericTaskApi *GenericTaskApi) StopGenericTask(c *gin.Context) {
	taskId := c.Query("taskId")
	if err := genericTaskService.StopGenericTask(taskId); err != nil {
		global.GVA_LOG.Error("停止任务失败!", zap.Error(err))
		response.FailWithMessage("停止任务失败", c)
	} else {
		response.OkWithMessage("停止任务成功", c)
	}
}
func (genericTaskApi *GenericTaskApi) StopGenericTasks(c *gin.Context) {
	var taskSetup octopusReq.TaskSetup
	err := c.ShouldBindJSON(&taskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := genericTaskService.StopGenericTasks(taskSetup); err != nil {
		global.GVA_LOG.Error("停止任务失败!", zap.Error(err))
		response.FailWithMessage("停止任务失败", c)
	} else {
		response.OkWithMessage("停止任务成功", c)
	}
}
func (genericTaskApi *GenericTaskApi) StartGenericTasks(c *gin.Context) {
	var taskSetup octopusReq.TaskSetup
	err := c.ShouldBindJSON(&taskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := genericTaskService.StartGenericTasks(taskSetup); err != nil {
		global.GVA_LOG.Error("运行任务失败!", zap.Error(err))
		response.FailWithMessage("运行任务失败", c)
	} else {
		response.OkWithMessage("运行任务成功", c)
	}
}
func (genericTaskApi *GenericTaskApi) DeleteGenericTasks(c *gin.Context) {
	var taskSetup octopusReq.TaskSetup
	err := c.ShouldBindJSON(&taskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := genericTaskService.DeleteGenericTasks(taskSetup, userId); err != nil {
		global.GVA_LOG.Error("删除任务失败!", zap.Error(err))
		response.FailWithMessage("删除任务失败", c)
	} else {
		response.OkWithMessage("删除任务成功", c)
	}
}

func (genericTaskApi *GenericTaskApi) DeleteGenericTask(c *gin.Context) {
	id := c.Query("id")
	userId := utils.GetUserID(c)
	if err := genericTaskService.DeleteGenericTask(id, userId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
