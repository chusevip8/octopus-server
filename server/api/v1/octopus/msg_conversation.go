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

type MsgConversationApi struct{}

var msgConversationService = service.ServiceGroupApp.OctopusServiceGroup.MsgConversationService

// CreateMsgConversation 创建私信会话纪录
// @Tags MsgConversation
// @Summary 创建私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.MsgConversation true "创建私信会话纪录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /msgConversation/createMsgConversation [post]
func (msgConversationApi *MsgConversationApi) CreateMsgConversation(c *gin.Context) {
	var msgConversation octopus.MsgConversation
	err := c.ShouldBindJSON(&msgConversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msgConversation.CreatedBy = utils.GetUserID(c)

	if err := msgConversationService.CreateMsgConversation(&msgConversation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMsgConversation 删除私信会话纪录
// @Tags MsgConversation
// @Summary 删除私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.MsgConversation true "删除私信会话纪录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /msgConversation/deleteMsgConversation [delete]
func (msgConversationApi *MsgConversationApi) DeleteMsgConversation(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := msgConversationService.DeleteMsgConversaton(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMsgConversationByIds 批量删除私信会话纪录
// @Tags MsgConversation
// @Summary 批量删除私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /msgConversation/deleteMsgConversationByIds [delete]
func (msgConversationApi *MsgConversationApi) DeleteMsgConversationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := msgConversationService.DeleteMsgConversationByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMsgConversation 更新私信会话纪录
// @Tags MsgConversation
// @Summary 更新私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.MsgConversation true "更新私信会话纪录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /msgConversation/updateMsgConversation [put]
func (msgConversationApi *MsgConversationApi) UpdateMsgConversation(c *gin.Context) {
	var msgConversation octopus.MsgConversation
	err := c.ShouldBindJSON(&msgConversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	msgConversation.UpdatedBy = utils.GetUserID(c)

	if err := msgConversationService.UpdateMsgConversation(msgConversation); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMsgConversation 用id查询私信会话纪录
// @Tags MsgConversation
// @Summary 用id查询私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.MsgConversation true "用id查询私信会话纪录"
// @Success 200 {object} response.Response{data=object{remsgConversation=octopus.MsgConversation},msg=string} "查询成功"
// @Router /msgConversation/findMsgConversation [get]
func (msgConversationApi *MsgConversationApi) FindMsgConversation(c *gin.Context) {
	ID := c.Query("ID")
	if remsgConversation, err := msgConversationService.GetMsgConversation(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(remsgConversation, c)
	}
}

// GetMsgConversationList 分页获取私信会话纪录列表
// @Tags MsgConversation
// @Summary 分页获取私信会话纪录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.MsgConversationSearch true "分页获取私信会话纪录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /msgConversation/getMsgConversationList [get]
func (msgConversationApi *MsgConversationApi) GetMsgConversationList(c *gin.Context) {
	var pageInfo octopusReq.MsgConversationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := msgConversationService.GetMsgConversationInfoList(pageInfo); err != nil {
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

// GetMsgConversationPublic 不需要鉴权的私信会话纪录接口
// @Tags MsgConversation
// @Summary 不需要鉴权的私信会话纪录接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.MsgConversationSearch true "分页获取私信会话纪录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /msgConversation/getMsgConversationPublic [get]
func (msgConversationApi *MsgConversationApi) GetMsgConversationPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的私信会话纪录接口信息",
	}, "获取成功", c)
}
