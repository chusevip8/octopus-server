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

type CmtThreadApi struct{}

var cmtThreadService = service.ServiceGroupApp.OctopusServiceGroup.CmtThreadService

// CreateCmtThread 创建评论会话
// @Tags CmtThread
// @Summary 创建评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtThread true "创建评论会话"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cmtThread/createCmtThread [post]
func (cmtThreadApi *CmtThreadApi) CreateCmtThread(c *gin.Context) {
	var cmtThread octopus.CmtThread
	err := c.ShouldBindJSON(&cmtThread)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cmtThread.CreatedBy = utils.GetUserID(c)

	if err := cmtThreadService.CreateCmtThread(&cmtThread); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmtThread 删除评论会话
// @Tags CmtThread
// @Summary 删除评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtThread true "删除评论会话"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cmtThread/deleteCmtThread [delete]
func (cmtThreadApi *CmtThreadApi) DeleteCmtThread(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := cmtThreadService.DeleteCmtThread(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmtThreadByIds 批量删除评论会话
// @Tags CmtThread
// @Summary 批量删除评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cmtThread/deleteCmtThreadByIds [delete]
func (cmtThreadApi *CmtThreadApi) DeleteCmtThreadByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := cmtThreadService.DeleteCmtThreadByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmtThread 更新评论会话
// @Tags CmtThread
// @Summary 更新评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtThread true "更新评论会话"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cmtThread/updateCmtThread [put]
func (cmtThreadApi *CmtThreadApi) UpdateCmtThread(c *gin.Context) {
	var cmtThread octopus.CmtThread
	err := c.ShouldBindJSON(&cmtThread)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cmtThread.UpdatedBy = utils.GetUserID(c)

	if err := cmtThreadService.UpdateCmtThread(cmtThread); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmtThread 用id查询评论会话
// @Tags CmtThread
// @Summary 用id查询评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.CmtThread true "用id查询评论会话"
// @Success 200 {object} response.Response{data=object{recmtThread=octopus.CmtThread},msg=string} "查询成功"
// @Router /cmtThread/findCmtThread [get]
func (cmtThreadApi *CmtThreadApi) FindCmtThread(c *gin.Context) {
	ID := c.Query("ID")
	if recmtThread, err := cmtThreadService.GetCmtThread(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recmtThread, c)
	}
}

// GetCmtThreadList 分页获取评论会话列表
// @Tags CmtThread
// @Summary 分页获取评论会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtThreadSearch true "分页获取评论会话列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cmtThread/getCmtThreadList [get]
func (cmtThreadApi *CmtThreadApi) GetCmtThreadList(c *gin.Context) {
	var pageInfo octopusReq.CmtThreadSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.CreatedBy = utils.GetUserID(c)
	if list, total, err := cmtThreadService.GetCmtThreadInfoList(pageInfo); err != nil {
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

// GetCmtThreadPublic 不需要鉴权的评论会话接口
// @Tags CmtThread
// @Summary 不需要鉴权的评论会话接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtThreadSearch true "分页获取评论会话列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cmtThread/getCmtThreadPublic [get]
func (cmtThreadApi *CmtThreadApi) GetCmtThreadPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的评论会话接口信息",
	}, "获取成功", c)
}
