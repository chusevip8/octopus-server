package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ConversationApi struct {}

var conversationService = service.ServiceGroupApp.OctopusServiceGroup.ConversationService


// CreateConversation 创建消息会话
// @Tags Conversation
// @Summary 创建消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Conversation true "创建消息会话"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /conversation/createConversation [post]
func (conversationApi *ConversationApi) CreateConversation(c *gin.Context) {
	var conversation octopus.Conversation
	err := c.ShouldBindJSON(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := conversationService.CreateConversation(&conversation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteConversation 删除消息会话
// @Tags Conversation
// @Summary 删除消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Conversation true "删除消息会话"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /conversation/deleteConversation [delete]
func (conversationApi *ConversationApi) DeleteConversation(c *gin.Context) {
	ID := c.Query("ID")
	if err := conversationService.DeleteConversation(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteConversationByIds 批量删除消息会话
// @Tags Conversation
// @Summary 批量删除消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /conversation/deleteConversationByIds [delete]
func (conversationApi *ConversationApi) DeleteConversationByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := conversationService.DeleteConversationByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateConversation 更新消息会话
// @Tags Conversation
// @Summary 更新消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Conversation true "更新消息会话"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /conversation/updateConversation [put]
func (conversationApi *ConversationApi) UpdateConversation(c *gin.Context) {
	var conversation octopus.Conversation
	err := c.ShouldBindJSON(&conversation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := conversationService.UpdateConversation(conversation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindConversation 用id查询消息会话
// @Tags Conversation
// @Summary 用id查询消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.Conversation true "用id查询消息会话"
// @Success 200 {object} response.Response{data=object{reconversation=octopus.Conversation},msg=string} "查询成功"
// @Router /conversation/findConversation [get]
func (conversationApi *ConversationApi) FindConversation(c *gin.Context) {
	ID := c.Query("ID")
	if reconversation, err := conversationService.GetConversation(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reconversation, c)
	}
}

// GetConversationList 分页获取消息会话列表
// @Tags Conversation
// @Summary 分页获取消息会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.ConversationSearch true "分页获取消息会话列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /conversation/getConversationList [get]
func (conversationApi *ConversationApi) GetConversationList(c *gin.Context) {
	var pageInfo octopusReq.ConversationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := conversationService.GetConversationInfoList(pageInfo); err != nil {
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

// GetConversationPublic 不需要鉴权的消息会话接口
// @Tags Conversation
// @Summary 不需要鉴权的消息会话接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.ConversationSearch true "分页获取消息会话列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /conversation/getConversationPublic [get]
func (conversationApi *ConversationApi) GetConversationPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的消息会话接口信息",
    }, "获取成功", c)
}
