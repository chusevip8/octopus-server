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

type IntervalTaskSetupApi struct{}

var intervalTaskSetupService = service.ServiceGroupApp.OctopusServiceGroup.IntervalTaskSetupService

// CreateIntervalTaskSetup 创建间隔任务设置
// @Tags IntervalTaskSetup
// @Summary 创建间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.IntervalTaskSetup true "创建间隔任务设置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /intervalTaskSetup/createIntervalTaskSetup [post]
func (intervalTaskApi *IntervalTaskSetupApi) CreateIntervalTaskSetup(c *gin.Context) {
	var intervalTaskSetup octopus.IntervalTaskSetup
	err := c.ShouldBindJSON(&intervalTaskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	intervalTaskSetup.CreatedBy = utils.GetUserID(c)

	if err := intervalTaskSetupService.CreateIntervalTaskSetup(&intervalTaskSetup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIntervalTaskSetup 删除间隔任务设置
// @Tags IntervalTaskSetup
// @Summary 删除间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.IntervalTaskSetup true "删除间隔任务设置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /intervalTaskSetup/deleteIntervalTaskSetup [delete]
func (intervalTaskApi *IntervalTaskSetupApi) DeleteIntervalTaskSetup(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := intervalTaskSetupService.DeleteIntervalTaskSetup(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIntervalTaskSetupByIds 批量删除间隔任务设置
// @Tags IntervalTaskSetup
// @Summary 批量删除间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /intervalTaskSetup/deleteIntervalTaskSetupByIds [delete]
func (intervalTaskApi *IntervalTaskSetupApi) DeleteIntervalTaskSetupByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := intervalTaskSetupService.DeleteIntervalTaskSetupByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIntervalTaskSetup 更新间隔任务设置
// @Tags IntervalTaskSetup
// @Summary 更新间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.IntervalTaskSetup true "更新间隔任务设置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /intervalTaskSetup/updateIntervalTaskSetup [put]
func (intervalTaskApi *IntervalTaskSetupApi) UpdateIntervalTaskSetup(c *gin.Context) {
	var intervalTaskSetup octopus.IntervalTaskSetup
	err := c.ShouldBindJSON(&intervalTaskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	intervalTaskSetup.UpdatedBy = utils.GetUserID(c)

	if err := intervalTaskSetupService.UpdateIntervalTaskSetup(intervalTaskSetup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIntervalTaskSetup 用id查询间隔任务设置
// @Tags IntervalTaskSetup
// @Summary 用id查询间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.IntervalTaskSetup true "用id查询间隔任务设置"
// @Success 200 {object} response.Response{data=object{reintervalTaskSetup=octopus.IntervalTaskSetup},msg=string} "查询成功"
// @Router /intervalTaskSetup/findIntervalTaskSetup [get]
func (intervalTaskApi *IntervalTaskSetupApi) FindIntervalTaskSetup(c *gin.Context) {
	ID := c.Query("ID")
	if reintervalTaskSetup, err := intervalTaskSetupService.GetIntervalTaskSetup(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(reintervalTaskSetup, c)
	}
}

// GetIntervalTaskSetupList 分页获取间隔任务设置列表
// @Tags IntervalTaskSetup
// @Summary 分页获取间隔任务设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.IntervalTaskSetupSearch true "分页获取间隔任务设置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /intervalTaskSetup/getIntervalTaskSetupList [get]
func (intervalTaskApi *IntervalTaskSetupApi) GetIntervalTaskSetupList(c *gin.Context) {
	var pageInfo octopusReq.IntervalTaskSetupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := intervalTaskSetupService.GetIntervalTaskSetupInfoList(pageInfo); err != nil {
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

// GetIntervalTaskSetupPublic 不需要鉴权的间隔任务设置接口
// @Tags IntervalTaskSetup
// @Summary 不需要鉴权的间隔任务设置接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.IntervalTaskSetupSearch true "分页获取间隔任务设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /intervalTaskSetup/getIntervalTaskSetupPublic [get]
func (intervalTaskApi *IntervalTaskSetupApi) GetIntervalTaskSetupPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的间隔任务设置接口信息",
	}, "获取成功", c)
}
