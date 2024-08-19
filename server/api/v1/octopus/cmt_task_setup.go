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

type CmtTaskSetupApi struct{}

var cmtTaskSetupService = service.ServiceGroupApp.OctopusServiceGroup.CmtTaskSetupService

// CreateCmtTaskSetup 创建评论任务设置
// @Tags CmtTaskSetup
// @Summary 创建评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskSetup true "创建评论任务设置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cmtTaskSetup/createCmtTaskSetup [post]
func (cmtTaskSetupApi *CmtTaskSetupApi) CreateCmtTaskSetup(c *gin.Context) {
	var cmtTaskSetup octopus.CmtTaskSetup
	err := c.ShouldBindJSON(&cmtTaskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cmtTaskSetup.CreatedBy = utils.GetUserID(c)

	if err := cmtTaskSetupService.CreateCmtTaskSetup(&cmtTaskSetup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmtTaskSetup 删除评论任务设置
// @Tags CmtTaskSetup
// @Summary 删除评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskSetup true "删除评论任务设置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cmtTaskSetup/deleteCmtTaskSetup [delete]
func (cmtTaskSetupApi *CmtTaskSetupApi) DeleteCmtTaskSetup(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := cmtTaskSetupService.DeleteCmtTaskSetup(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmtTaskSetupByIds 批量删除评论任务设置
// @Tags CmtTaskSetup
// @Summary 批量删除评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cmtTaskSetup/deleteCmtTaskSetupByIds [delete]
func (cmtTaskSetupApi *CmtTaskSetupApi) DeleteCmtTaskSetupByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := cmtTaskSetupService.DeleteCmtTaskSetupByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmtTaskSetup 更新评论任务设置
// @Tags CmtTaskSetup
// @Summary 更新评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskSetup true "更新评论任务设置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cmtTaskSetup/updateCmtTaskSetup [put]
func (cmtTaskSetupApi *CmtTaskSetupApi) UpdateCmtTaskSetup(c *gin.Context) {
	var cmtTaskSetup octopus.CmtTaskSetup
	err := c.ShouldBindJSON(&cmtTaskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cmtTaskSetup.UpdatedBy = utils.GetUserID(c)

	if err := cmtTaskSetupService.UpdateCmtTaskSetup(cmtTaskSetup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmtTaskSetup 用id查询评论任务设置
// @Tags CmtTaskSetup
// @Summary 用id查询评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.CmtTaskSetup true "用id查询评论任务设置"
// @Success 200 {object} response.Response{data=object{recmtTaskSetup=octopus.CmtTaskSetup},msg=string} "查询成功"
// @Router /cmtTaskSetup/findCmtTaskSetup [get]
func (cmtTaskSetupApi *CmtTaskSetupApi) FindCmtTaskSetup(c *gin.Context) {
	ID := c.Query("ID")
	if recmtTaskSetup, err := cmtTaskSetupService.GetCmtTaskSetup(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recmtTaskSetup, c)
	}
}

// GetCmtTaskSetupList 分页获取评论任务设置列表
// @Tags CmtTaskSetup
// @Summary 分页获取评论任务设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtTaskSetupSearch true "分页获取评论任务设置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cmtTaskSetup/getCmtTaskSetupList [get]
func (cmtTaskSetupApi *CmtTaskSetupApi) GetCmtTaskSetupList(c *gin.Context) {
	var pageInfo octopusReq.CmtTaskSetupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.CreatedBy = utils.GetUserID(c)
	if list, total, err := cmtTaskSetupService.GetCmtTaskSetupInfoList(pageInfo); err != nil {
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

// GetCmtTaskSetupPublic 不需要鉴权的评论任务设置接口
// @Tags CmtTaskSetup
// @Summary 不需要鉴权的评论任务设置接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtTaskSetupSearch true "分页获取评论任务设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cmtTaskSetup/getCmtTaskSetupPublic [get]
func (cmtTaskSetupApi *CmtTaskSetupApi) GetCmtTaskSetupPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的评论任务设置接口信息",
	}, "获取成功", c)
}
