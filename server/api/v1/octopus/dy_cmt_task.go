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

type DYCmtTaskApi struct {
}

var dyCmtTaskService = service.ServiceGroupApp.OctopusServiceGroup.DYCmtTaskService


// CreateDYCmtTask 创建抖音评论任务
// @Tags DYCmtTask
// @Summary 创建抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.DYCmtTask true "创建抖音评论任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dyCmtTask/createDYCmtTask [post]
func (dyCmtTaskApi *DYCmtTaskApi) CreateDYCmtTask(c *gin.Context) {
	var dyCmtTask octopus.DYCmtTask
	err := c.ShouldBindJSON(&dyCmtTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    dyCmtTask.CreatedBy = utils.GetUserID(c)

	if err := dyCmtTaskService.CreateDYCmtTask(&dyCmtTask); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDYCmtTask 删除抖音评论任务
// @Tags DYCmtTask
// @Summary 删除抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.DYCmtTask true "删除抖音评论任务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dyCmtTask/deleteDYCmtTask [delete]
func (dyCmtTaskApi *DYCmtTaskApi) DeleteDYCmtTask(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := dyCmtTaskService.DeleteDYCmtTask(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDYCmtTaskByIds 批量删除抖音评论任务
// @Tags DYCmtTask
// @Summary 批量删除抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dyCmtTask/deleteDYCmtTaskByIds [delete]
func (dyCmtTaskApi *DYCmtTaskApi) DeleteDYCmtTaskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := dyCmtTaskService.DeleteDYCmtTaskByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDYCmtTask 更新抖音评论任务
// @Tags DYCmtTask
// @Summary 更新抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.DYCmtTask true "更新抖音评论任务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dyCmtTask/updateDYCmtTask [put]
func (dyCmtTaskApi *DYCmtTaskApi) UpdateDYCmtTask(c *gin.Context) {
	var dyCmtTask octopus.DYCmtTask
	err := c.ShouldBindJSON(&dyCmtTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    dyCmtTask.UpdatedBy = utils.GetUserID(c)

	if err := dyCmtTaskService.UpdateDYCmtTask(dyCmtTask); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDYCmtTask 用id查询抖音评论任务
// @Tags DYCmtTask
// @Summary 用id查询抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.DYCmtTask true "用id查询抖音评论任务"
// @Success 200 {object} response.Response{data=object{redyCmtTask=octopus.DYCmtTask},msg=string} "查询成功"
// @Router /dyCmtTask/findDYCmtTask [get]
func (dyCmtTaskApi *DYCmtTaskApi) FindDYCmtTask(c *gin.Context) {
	ID := c.Query("ID")
	if redyCmtTask, err := dyCmtTaskService.GetDYCmtTask(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redyCmtTask": redyCmtTask}, c)
	}
}

// GetDYCmtTaskList 分页获取抖音评论任务列表
// @Tags DYCmtTask
// @Summary 分页获取抖音评论任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.DYCmtTaskSearch true "分页获取抖音评论任务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dyCmtTask/getDYCmtTaskList [get]
func (dyCmtTaskApi *DYCmtTaskApi) GetDYCmtTaskList(c *gin.Context) {
	var pageInfo octopusReq.DYCmtTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dyCmtTaskService.GetDYCmtTaskInfoList(pageInfo); err != nil {
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

// GetDYCmtTaskPublic 不需要鉴权的抖音评论任务接口
// @Tags DYCmtTask
// @Summary 不需要鉴权的抖音评论任务接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.DYCmtTaskSearch true "分页获取抖音评论任务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dyCmtTask/getDYCmtTaskPublic [get]
func (dyCmtTaskApi *DYCmtTaskApi) GetDYCmtTaskPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的抖音评论任务接口信息",
    }, "获取成功", c)
}
