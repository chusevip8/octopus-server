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

type ThreadApi struct {}

var threadService = service.ServiceGroupApp.OctopusServiceGroup.ThreadService


// CreateThread 创建消息组
// @Tags Thread
// @Summary 创建消息组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Thread true "创建消息组"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /thread/createThread [post]
func (threadApi *ThreadApi) CreateThread(c *gin.Context) {
	var thread octopus.Thread
	err := c.ShouldBindJSON(&thread)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := threadService.CreateThread(&thread); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteThread 删除消息组
// @Tags Thread
// @Summary 删除消息组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Thread true "删除消息组"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /thread/deleteThread [delete]
func (threadApi *ThreadApi) DeleteThread(c *gin.Context) {
	ID := c.Query("ID")
	if err := threadService.DeleteThread(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteThreadByIds 批量删除消息组
// @Tags Thread
// @Summary 批量删除消息组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /thread/deleteThreadByIds [delete]
func (threadApi *ThreadApi) DeleteThreadByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := threadService.DeleteThreadByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateThread 更新消息组
// @Tags Thread
// @Summary 更新消息组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Thread true "更新消息组"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /thread/updateThread [put]
func (threadApi *ThreadApi) UpdateThread(c *gin.Context) {
	var thread octopus.Thread
	err := c.ShouldBindJSON(&thread)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := threadService.UpdateThread(thread); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindThread 用id查询消息组
// @Tags Thread
// @Summary 用id查询消息组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.Thread true "用id查询消息组"
// @Success 200 {object} response.Response{data=object{rethread=octopus.Thread},msg=string} "查询成功"
// @Router /thread/findThread [get]
func (threadApi *ThreadApi) FindThread(c *gin.Context) {
	ID := c.Query("ID")
	if rethread, err := threadService.GetThread(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rethread, c)
	}
}

// GetThreadList 分页获取消息组列表
// @Tags Thread
// @Summary 分页获取消息组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.ThreadSearch true "分页获取消息组列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /thread/getThreadList [get]
func (threadApi *ThreadApi) GetThreadList(c *gin.Context) {
	var pageInfo octopusReq.ThreadSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := threadService.GetThreadInfoList(pageInfo); err != nil {
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

// GetThreadPublic 不需要鉴权的消息组接口
// @Tags Thread
// @Summary 不需要鉴权的消息组接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.ThreadSearch true "分页获取消息组列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /thread/getThreadPublic [get]
func (threadApi *ThreadApi) GetThreadPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的消息组接口信息",
    }, "获取成功", c)
}
