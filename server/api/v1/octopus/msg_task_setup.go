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

type MsgTaskSetupApi struct {}

var msgTaskSetupService = service.ServiceGroupApp.OctopusServiceGroup.MsgTaskSetupService


// CreateMsgTaskSetup 创建私信任务设置
// @Tags MsgTaskSetup
// @Summary 创建私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.MsgTaskSetup true "创建私信任务设置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /msgTaskSetup/createMsgTaskSetup [post]
func (msgTaskSetupApi *MsgTaskSetupApi) CreateMsgTaskSetup(c *gin.Context) {
	var msgTaskSetup octopus.MsgTaskSetup
	err := c.ShouldBindJSON(&msgTaskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    msgTaskSetup.CreatedBy = utils.GetUserID(c)

	if err := msgTaskSetupService.CreateMsgTaskSetup(&msgTaskSetup); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMsgTaskSetup 删除私信任务设置
// @Tags MsgTaskSetup
// @Summary 删除私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.MsgTaskSetup true "删除私信任务设置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /msgTaskSetup/deleteMsgTaskSetup [delete]
func (msgTaskSetupApi *MsgTaskSetupApi) DeleteMsgTaskSetup(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := msgTaskSetupService.DeleteMsgTaskSetup(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMsgTaskSetupByIds 批量删除私信任务设置
// @Tags MsgTaskSetup
// @Summary 批量删除私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /msgTaskSetup/deleteMsgTaskSetupByIds [delete]
func (msgTaskSetupApi *MsgTaskSetupApi) DeleteMsgTaskSetupByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := msgTaskSetupService.DeleteMsgTaskSetupByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMsgTaskSetup 更新私信任务设置
// @Tags MsgTaskSetup
// @Summary 更新私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.MsgTaskSetup true "更新私信任务设置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /msgTaskSetup/updateMsgTaskSetup [put]
func (msgTaskSetupApi *MsgTaskSetupApi) UpdateMsgTaskSetup(c *gin.Context) {
	var msgTaskSetup octopus.MsgTaskSetup
	err := c.ShouldBindJSON(&msgTaskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    msgTaskSetup.UpdatedBy = utils.GetUserID(c)

	if err := msgTaskSetupService.UpdateMsgTaskSetup(msgTaskSetup); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMsgTaskSetup 用id查询私信任务设置
// @Tags MsgTaskSetup
// @Summary 用id查询私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.MsgTaskSetup true "用id查询私信任务设置"
// @Success 200 {object} response.Response{data=object{remsgTaskSetup=octopus.MsgTaskSetup},msg=string} "查询成功"
// @Router /msgTaskSetup/findMsgTaskSetup [get]
func (msgTaskSetupApi *MsgTaskSetupApi) FindMsgTaskSetup(c *gin.Context) {
	ID := c.Query("ID")
	if remsgTaskSetup, err := msgTaskSetupService.GetMsgTaskSetup(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(remsgTaskSetup, c)
	}
}

// GetMsgTaskSetupList 分页获取私信任务设置列表
// @Tags MsgTaskSetup
// @Summary 分页获取私信任务设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.MsgTaskSetupSearch true "分页获取私信任务设置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /msgTaskSetup/getMsgTaskSetupList [get]
func (msgTaskSetupApi *MsgTaskSetupApi) GetMsgTaskSetupList(c *gin.Context) {
	var pageInfo octopusReq.MsgTaskSetupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := msgTaskSetupService.GetMsgTaskSetupInfoList(pageInfo); err != nil {
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

// GetMsgTaskSetupPublic 不需要鉴权的私信任务设置接口
// @Tags MsgTaskSetup
// @Summary 不需要鉴权的私信任务设置接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.MsgTaskSetupSearch true "分页获取私信任务设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /msgTaskSetup/getMsgTaskSetupPublic [get]
func (msgTaskSetupApi *MsgTaskSetupApi) GetMsgTaskSetupPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的私信任务设置接口信息",
    }, "获取成功", c)
}
