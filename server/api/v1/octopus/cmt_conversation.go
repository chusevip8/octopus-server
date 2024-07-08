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

type CmtConversationApi struct {}

var cmtConversationService = service.ServiceGroupApp.OctopusServiceGroup.CmtConversationService


// CreateCmtConversation 创建评论会话记录
// @Tags CmtConversation
// @Summary 创建评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtConversation true "创建评论会话记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cmtConversation/createCmtConversation [post]
func (cmtConversationApi *CmtConversationApi) CreateCmtConversation(c *gin.Context) {
	var cmtConversation octopus.CmtConversation
	err := c.ShouldBindJSON(&cmtConversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cmtConversation.CreatedBy = utils.GetUserID(c)

	if err := cmtConversationService.CreateCmtConversation(&cmtConversation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmtConversation 删除评论会话记录
// @Tags CmtConversation
// @Summary 删除评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtConversation true "删除评论会话记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cmtConversation/deleteCmtConversation [delete]
func (cmtConversationApi *CmtConversationApi) DeleteCmtConversation(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := cmtConversationService.DeleteCmtConversation(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmtConversationByIds 批量删除评论会话记录
// @Tags CmtConversation
// @Summary 批量删除评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cmtConversation/deleteCmtConversationByIds [delete]
func (cmtConversationApi *CmtConversationApi) DeleteCmtConversationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := cmtConversationService.DeleteCmtConversationByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmtConversation 更新评论会话记录
// @Tags CmtConversation
// @Summary 更新评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtConversation true "更新评论会话记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cmtConversation/updateCmtConversation [put]
func (cmtConversationApi *CmtConversationApi) UpdateCmtConversation(c *gin.Context) {
	var cmtConversation octopus.CmtConversation
	err := c.ShouldBindJSON(&cmtConversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cmtConversation.UpdatedBy = utils.GetUserID(c)

	if err := cmtConversationService.UpdateCmtConversation(cmtConversation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmtConversation 用id查询评论会话记录
// @Tags CmtConversation
// @Summary 用id查询评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.CmtConversation true "用id查询评论会话记录"
// @Success 200 {object} response.Response{data=object{recmtConversation=octopus.CmtConversation},msg=string} "查询成功"
// @Router /cmtConversation/findCmtConversation [get]
func (cmtConversationApi *CmtConversationApi) FindCmtConversation(c *gin.Context) {
	ID := c.Query("ID")
	if recmtConversation, err := cmtConversationService.GetCmtConversation(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recmtConversation, c)
	}
}

// GetCmtConversationList 分页获取评论会话记录列表
// @Tags CmtConversation
// @Summary 分页获取评论会话记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtConversationSearch true "分页获取评论会话记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cmtConversation/getCmtConversationList [get]
func (cmtConversationApi *CmtConversationApi) GetCmtConversationList(c *gin.Context) {
	var pageInfo octopusReq.CmtConversationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmtConversationService.GetCmtConversationInfoList(pageInfo); err != nil {
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

// GetCmtConversationPublic 不需要鉴权的评论会话记录接口
// @Tags CmtConversation
// @Summary 不需要鉴权的评论会话记录接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtConversationSearch true "分页获取评论会话记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cmtConversation/getCmtConversationPublic [get]
func (cmtConversationApi *CmtConversationApi) GetCmtConversationPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的评论会话记录接口信息",
    }, "获取成功", c)
}
