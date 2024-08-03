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

type TaskBindDataApi struct {}

var taskBindDataService = service.ServiceGroupApp.OctopusServiceGroup.TaskBindDataService


// CreateTaskBindData 创建任务数据
// @Tags TaskBindData
// @Summary 创建任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.TaskBindData true "创建任务数据"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /taskBindData/createTaskBindData [post]
func (taskBindDataApi *TaskBindDataApi) CreateTaskBindData(c *gin.Context) {
	var taskBindData octopus.TaskBindData
	err := c.ShouldBindJSON(&taskBindData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    taskBindData.CreatedBy = utils.GetUserID(c)

	if err := taskBindDataService.CreateTaskBindData(&taskBindData); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTaskBindData 删除任务数据
// @Tags TaskBindData
// @Summary 删除任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.TaskBindData true "删除任务数据"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /taskBindData/deleteTaskBindData [delete]
func (taskBindDataApi *TaskBindDataApi) DeleteTaskBindData(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := taskBindDataService.DeleteTaskBindData(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTaskBindDataByIds 批量删除任务数据
// @Tags TaskBindData
// @Summary 批量删除任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /taskBindData/deleteTaskBindDataByIds [delete]
func (taskBindDataApi *TaskBindDataApi) DeleteTaskBindDataByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := taskBindDataService.DeleteTaskBindDataByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTaskBindData 更新任务数据
// @Tags TaskBindData
// @Summary 更新任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.TaskBindData true "更新任务数据"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /taskBindData/updateTaskBindData [put]
func (taskBindDataApi *TaskBindDataApi) UpdateTaskBindData(c *gin.Context) {
	var taskBindData octopus.TaskBindData
	err := c.ShouldBindJSON(&taskBindData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    taskBindData.UpdatedBy = utils.GetUserID(c)

	if err := taskBindDataService.UpdateTaskBindData(taskBindData); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTaskBindData 用id查询任务数据
// @Tags TaskBindData
// @Summary 用id查询任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.TaskBindData true "用id查询任务数据"
// @Success 200 {object} response.Response{data=object{retaskBindData=octopus.TaskBindData},msg=string} "查询成功"
// @Router /taskBindData/findTaskBindData [get]
func (taskBindDataApi *TaskBindDataApi) FindTaskBindData(c *gin.Context) {
	ID := c.Query("ID")
	if retaskBindData, err := taskBindDataService.GetTaskBindData(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(retaskBindData, c)
	}
}

// GetTaskBindDataList 分页获取任务数据列表
// @Tags TaskBindData
// @Summary 分页获取任务数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.TaskBindDataSearch true "分页获取任务数据列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /taskBindData/getTaskBindDataList [get]
func (taskBindDataApi *TaskBindDataApi) GetTaskBindDataList(c *gin.Context) {
	var pageInfo octopusReq.TaskBindDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := taskBindDataService.GetTaskBindDataInfoList(pageInfo); err != nil {
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

// GetTaskBindDataPublic 不需要鉴权的任务数据接口
// @Tags TaskBindData
// @Summary 不需要鉴权的任务数据接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.TaskBindDataSearch true "分页获取任务数据列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /taskBindData/getTaskBindDataPublic [get]
func (taskBindDataApi *TaskBindDataApi) GetTaskBindDataPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的任务数据接口信息",
    }, "获取成功", c)
}
