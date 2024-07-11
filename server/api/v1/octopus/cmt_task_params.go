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

type CmtTaskParamsApi struct {}

var cmtTaskParamsService = service.ServiceGroupApp.OctopusServiceGroup.CmtTaskParamsService


// CreateCmtTaskParams 创建评论任务参数
// @Tags CmtTaskParams
// @Summary 创建评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskParams true "创建评论任务参数"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cmtTaskParams/createCmtTaskParams [post]
func (cmtTaskParamsApi *CmtTaskParamsApi) CreateCmtTaskParams(c *gin.Context) {
	var cmtTaskParams octopus.CmtTaskParams
	err := c.ShouldBindJSON(&cmtTaskParams)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cmtTaskParams.CreatedBy = utils.GetUserID(c)

	if err := cmtTaskParamsService.CreateCmtTaskParams(&cmtTaskParams); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmtTaskParams 删除评论任务参数
// @Tags CmtTaskParams
// @Summary 删除评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskParams true "删除评论任务参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cmtTaskParams/deleteCmtTaskParams [delete]
func (cmtTaskParamsApi *CmtTaskParamsApi) DeleteCmtTaskParams(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := cmtTaskParamsService.DeleteCmtTaskParams(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmtTaskParamsByIds 批量删除评论任务参数
// @Tags CmtTaskParams
// @Summary 批量删除评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cmtTaskParams/deleteCmtTaskParamsByIds [delete]
func (cmtTaskParamsApi *CmtTaskParamsApi) DeleteCmtTaskParamsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := cmtTaskParamsService.DeleteCmtTaskParamsByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmtTaskParams 更新评论任务参数
// @Tags CmtTaskParams
// @Summary 更新评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskParams true "更新评论任务参数"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cmtTaskParams/updateCmtTaskParams [put]
func (cmtTaskParamsApi *CmtTaskParamsApi) UpdateCmtTaskParams(c *gin.Context) {
	var cmtTaskParams octopus.CmtTaskParams
	err := c.ShouldBindJSON(&cmtTaskParams)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cmtTaskParams.UpdatedBy = utils.GetUserID(c)

	if err := cmtTaskParamsService.UpdateCmtTaskParams(cmtTaskParams); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmtTaskParams 用id查询评论任务参数
// @Tags CmtTaskParams
// @Summary 用id查询评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.CmtTaskParams true "用id查询评论任务参数"
// @Success 200 {object} response.Response{data=object{recmtTaskParams=octopus.CmtTaskParams},msg=string} "查询成功"
// @Router /cmtTaskParams/findCmtTaskParams [get]
func (cmtTaskParamsApi *CmtTaskParamsApi) FindCmtTaskParams(c *gin.Context) {
	ID := c.Query("ID")
	if recmtTaskParams, err := cmtTaskParamsService.GetCmtTaskParams(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recmtTaskParams, c)
	}
}

// GetCmtTaskParamsList 分页获取评论任务参数列表
// @Tags CmtTaskParams
// @Summary 分页获取评论任务参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtTaskParamsSearch true "分页获取评论任务参数列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cmtTaskParams/getCmtTaskParamsList [get]
func (cmtTaskParamsApi *CmtTaskParamsApi) GetCmtTaskParamsList(c *gin.Context) {
	var pageInfo octopusReq.CmtTaskParamsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmtTaskParamsService.GetCmtTaskParamsInfoList(pageInfo); err != nil {
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

// GetCmtTaskParamsPublic 不需要鉴权的评论任务参数接口
// @Tags CmtTaskParams
// @Summary 不需要鉴权的评论任务参数接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtTaskParamsSearch true "分页获取评论任务参数列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cmtTaskParams/getCmtTaskParamsPublic [get]
func (cmtTaskParamsApi *CmtTaskParamsApi) GetCmtTaskParamsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的评论任务参数接口信息",
    }, "获取成功", c)
}
