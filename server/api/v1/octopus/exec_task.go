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

type ExecTaskApi struct {
}

var execTaskService = service.ServiceGroupApp.OctopusServiceGroup.ExecTaskService


// CreateExecTask 创建执行任务
// @Tags ExecTask
// @Summary 创建执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.ExecTask true "创建执行任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /execTask/createExecTask [post]
func (execTaskApi *ExecTaskApi) CreateExecTask(c *gin.Context) {
	var execTask octopus.ExecTask
	err := c.ShouldBindJSON(&execTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    execTask.CreatedBy = utils.GetUserID(c)

	if err := execTaskService.CreateExecTask(&execTask); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteExecTask 删除执行任务
// @Tags ExecTask
// @Summary 删除执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.ExecTask true "删除执行任务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /execTask/deleteExecTask [delete]
func (execTaskApi *ExecTaskApi) DeleteExecTask(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := execTaskService.DeleteExecTask(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExecTaskByIds 批量删除执行任务
// @Tags ExecTask
// @Summary 批量删除执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /execTask/deleteExecTaskByIds [delete]
func (execTaskApi *ExecTaskApi) DeleteExecTaskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := execTaskService.DeleteExecTaskByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExecTask 更新执行任务
// @Tags ExecTask
// @Summary 更新执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.ExecTask true "更新执行任务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /execTask/updateExecTask [put]
func (execTaskApi *ExecTaskApi) UpdateExecTask(c *gin.Context) {
	var execTask octopus.ExecTask
	err := c.ShouldBindJSON(&execTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    execTask.UpdatedBy = utils.GetUserID(c)

	if err := execTaskService.UpdateExecTask(execTask); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExecTask 用id查询执行任务
// @Tags ExecTask
// @Summary 用id查询执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.ExecTask true "用id查询执行任务"
// @Success 200 {object} response.Response{data=object{reexecTask=octopus.ExecTask},msg=string} "查询成功"
// @Router /execTask/findExecTask [get]
func (execTaskApi *ExecTaskApi) FindExecTask(c *gin.Context) {
	ID := c.Query("ID")
	if reexecTask, err := execTaskService.GetExecTask(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reexecTask, c)
	}
}

// GetExecTaskList 分页获取执行任务列表
// @Tags ExecTask
// @Summary 分页获取执行任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.ExecTaskSearch true "分页获取执行任务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /execTask/getExecTaskList [get]
func (execTaskApi *ExecTaskApi) GetExecTaskList(c *gin.Context) {
	var pageInfo octopusReq.ExecTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := execTaskService.GetExecTaskInfoList(pageInfo); err != nil {
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

// GetExecTaskPublic 不需要鉴权的执行任务接口
// @Tags ExecTask
// @Summary 不需要鉴权的执行任务接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.ExecTaskSearch true "分页获取执行任务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /execTask/getExecTaskPublic [get]
func (execTaskApi *ExecTaskApi) GetExecTaskPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的执行任务接口信息",
    }, "获取成功", c)
}
