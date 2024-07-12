package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TaskApi struct{}

var taskService = service.ServiceGroupApp.OctopusServiceGroup.TaskService

// CreateTask 创建任务
// @Tags Task
// @Summary 创建任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Task true "创建任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /task/createTask [post]
func (taskApi *TaskApi) CreateTask(c *gin.Context) {
	var task octopus.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	task.CreatedBy = utils.GetUserID(c)

	if err := taskService.CreateTask(&task); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTask 删除任务
// @Tags Task
// @Summary 删除任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Task true "删除任务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /task/deleteTask [delete]
func (taskApi *TaskApi) DeleteTask(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := taskService.DeleteTask(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskByIds 批量删除任务
// @Tags Task
// @Summary 批量删除任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /task/deleteTaskByIds [delete]
func (taskApi *TaskApi) DeleteTaskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := taskService.DeleteTaskByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTask 更新任务
// @Tags Task
// @Summary 更新任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Task true "更新任务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /task/updateTask [put]
func (taskApi *TaskApi) UpdateTask(c *gin.Context) {
	var task octopus.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	task.UpdatedBy = utils.GetUserID(c)

	if err := taskService.UpdateTask(task); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTask 用id查询任务
// @Tags Task
// @Summary 用id查询任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.Task true "用id查询任务"
// @Success 200 {object} response.Response{data=object{retask=octopus.Task},msg=string} "查询成功"
// @Router /task/findTask [get]
func (taskApi *TaskApi) FindTask(c *gin.Context) {
	ID := c.Query("ID")
	if retask, err := taskService.GetTask(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(retask, c)
	}
}

func (taskApi *TaskApi) FindTaskByDeviceId(c *gin.Context) {
	deviceId := c.Query("deviceId")
	if retask, err := taskService.GetTaskByDeviceId(deviceId); err != nil {
		response.OkWithData(nil, c)
	} else {
		response.OkWithData(retask, c)
	}
}

// GetTaskList 分页获取任务列表
// @Tags Task
// @Summary 分页获取任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.TaskSearch true "分页获取任务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /task/getTaskList [get]
func (taskApi *TaskApi) GetTaskList(c *gin.Context) {
	var pageInfo octopusReq.TaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskService.GetTaskInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetTaskPublic 不需要鉴权的任务接口
// @Tags Task
// @Summary 不需要鉴权的任务接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.TaskSearch true "分页获取任务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /task/getTaskPublic [get]
func (taskApi *TaskApi) GetTaskPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的任务接口信息",
	}, "获取成功", c)
}
