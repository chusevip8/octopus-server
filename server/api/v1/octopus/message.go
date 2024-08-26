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

type MessageApi struct {}

var messageService = service.ServiceGroupApp.OctopusServiceGroup.MessageService


// CreateMessage 创建私信
// @Tags Message
// @Summary 创建私信
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Message true "创建私信"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /message/createMessage [post]
func (messageApi *MessageApi) CreateMessage(c *gin.Context) {
	var message octopus.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    message.CreatedBy = utils.GetUserID(c)

	if err := messageService.CreateMessage(&message); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMessage 删除私信
// @Tags Message
// @Summary 删除私信
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Message true "删除私信"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /message/deleteMessage [delete]
func (messageApi *MessageApi) DeleteMessage(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := messageService.DeleteMessage(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMessageByIds 批量删除私信
// @Tags Message
// @Summary 批量删除私信
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /message/deleteMessageByIds [delete]
func (messageApi *MessageApi) DeleteMessageByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := messageService.DeleteMessageByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMessage 更新私信
// @Tags Message
// @Summary 更新私信
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Message true "更新私信"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /message/updateMessage [put]
func (messageApi *MessageApi) UpdateMessage(c *gin.Context) {
	var message octopus.Message
	err := c.ShouldBindJSON(&message)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    message.UpdatedBy = utils.GetUserID(c)

	if err := messageService.UpdateMessage(message); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMessage 用id查询私信
// @Tags Message
// @Summary 用id查询私信
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.Message true "用id查询私信"
// @Success 200 {object} response.Response{data=object{remessage=octopus.Message},msg=string} "查询成功"
// @Router /message/findMessage [get]
func (messageApi *MessageApi) FindMessage(c *gin.Context) {
	ID := c.Query("ID")
	if remessage, err := messageService.GetMessage(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(remessage, c)
	}
}

// GetMessageList 分页获取私信列表
// @Tags Message
// @Summary 分页获取私信列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.MessageSearch true "分页获取私信列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /message/getMessageList [get]
func (messageApi *MessageApi) GetMessageList(c *gin.Context) {
	var pageInfo octopusReq.MessageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := messageService.GetMessageInfoList(pageInfo); err != nil {
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

// GetMessagePublic 不需要鉴权的私信接口
// @Tags Message
// @Summary 不需要鉴权的私信接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.MessageSearch true "分页获取私信列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /message/getMessagePublic [get]
func (messageApi *MessageApi) GetMessagePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的私信接口信息",
    }, "获取成功", c)
}
