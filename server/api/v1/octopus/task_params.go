package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type TaskParamsApi struct {}

var taskParamsService = service.ServiceGroupApp.OctopusServiceGroup.TaskParamsService


// CreateTaskParams 创建任务参数
// @Tags TaskParams
// @Summary 创建任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.TaskParams true "创建任务参数"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /taskParams/createTaskParams [post]
func (taskParamsApi *TaskParamsApi) CreateTaskParams(c *gin.Context) {
	var taskParams octopus.TaskParams
	err := c.ShouldBindJSON(&taskParams)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    taskParams.CreatedBy = utils.GetUserID(c)

	if err := taskParamsService.CreateTaskParams(&taskParams); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTaskParams 删除任务参数
// @Tags TaskParams
// @Summary 删除任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.TaskParams true "删除任务参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /taskParams/deleteTaskParams [delete]
func (taskParamsApi *TaskParamsApi) DeleteTaskParams(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := taskParamsService.DeleteTaskParams(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskParamsByIds 批量删除任务参数
// @Tags TaskParams
// @Summary 批量删除任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /taskParams/deleteTaskParamsByIds [delete]
func (taskParamsApi *TaskParamsApi) DeleteTaskParamsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := taskParamsService.DeleteTaskParamsByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTaskParams 更新任务参数
// @Tags TaskParams
// @Summary 更新任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.TaskParams true "更新任务参数"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /taskParams/updateTaskParams [put]
func (taskParamsApi *TaskParamsApi) UpdateTaskParams(c *gin.Context) {
	var taskParams octopus.TaskParams
	err := c.ShouldBindJSON(&taskParams)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    taskParams.UpdatedBy = utils.GetUserID(c)

	if err := taskParamsService.UpdateTaskParams(taskParams); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTaskParams 用id查询任务参数
// @Tags TaskParams
// @Summary 用id查询任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.TaskParams true "用id查询任务参数"
// @Success 200 {object} response.Response{data=object{retaskParams=octopus.TaskParams},msg=string} "查询成功"
// @Router /taskParams/findTaskParams [get]
func (taskParamsApi *TaskParamsApi) FindTaskParams(c *gin.Context) {
	ID := c.Query("ID")
	if retaskParams, err := taskParamsService.GetTaskParams(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(retaskParams, c)
	}
}

// GetTaskParamsList 分页获取任务参数列表
// @Tags TaskParams
// @Summary 分页获取任务参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.TaskParamsSearch true "分页获取任务参数列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /taskParams/getTaskParamsList [get]
func (taskParamsApi *TaskParamsApi) GetTaskParamsList(c *gin.Context) {
	var pageInfo octopusReq.TaskParamsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskParamsService.GetTaskParamsInfoList(pageInfo); err != nil {
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

// GetTaskParamsPublic 不需要鉴权的任务参数接口
// @Tags TaskParams
// @Summary 不需要鉴权的任务参数接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.TaskParamsSearch true "分页获取任务参数列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /taskParams/getTaskParamsPublic [get]
func (taskParamsApi *TaskParamsApi) GetTaskParamsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的任务参数接口信息",
    }, "获取成功", c)
}
