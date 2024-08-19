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

type GenericTaskSetupApi struct{}

var genericTaskSetupService = service.ServiceGroupApp.OctopusServiceGroup.GenericTaskSetupService

// CreateGenericTaskSetup 创建通用任务设置
// @Tags GenericTaskSetup
// @Summary 创建通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.GenericTaskSetup true "创建通用任务设置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /genericTaskSetup/createGenericTaskSetup [post]
func (genericTaskSetupApi *GenericTaskSetupApi) CreateGenericTaskSetup(c *gin.Context) {
	var genericTaskSetup octopus.GenericTaskSetup
	err := c.ShouldBindJSON(&genericTaskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	genericTaskSetup.CreatedBy = utils.GetUserID(c)

	if err := genericTaskSetupService.CreateGenericTaskSetup(&genericTaskSetup); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGenericTaskSetup 删除通用任务设置
// @Tags GenericTaskSetup
// @Summary 删除通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.GenericTaskSetup true "删除通用任务设置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /genericTaskSetup/deleteGenericTaskSetup [delete]
func (genericTaskSetupApi *GenericTaskSetupApi) DeleteGenericTaskSetup(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := genericTaskSetupService.DeleteGenericTaskSetup(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGenericTaskSetupByIds 批量删除通用任务设置
// @Tags GenericTaskSetup
// @Summary 批量删除通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /genericTaskSetup/deleteGenericTaskSetupByIds [delete]
func (genericTaskSetupApi *GenericTaskSetupApi) DeleteGenericTaskSetupByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := genericTaskSetupService.DeleteGenericTaskSetupByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGenericTaskSetup 更新通用任务设置
// @Tags GenericTaskSetup
// @Summary 更新通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.GenericTaskSetup true "更新通用任务设置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /genericTaskSetup/updateGenericTaskSetup [put]
func (genericTaskSetupApi *GenericTaskSetupApi) UpdateGenericTaskSetup(c *gin.Context) {
	var genericTaskSetup octopus.GenericTaskSetup
	err := c.ShouldBindJSON(&genericTaskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	genericTaskSetup.UpdatedBy = utils.GetUserID(c)

	if err := genericTaskSetupService.UpdateGenericTaskSetup(genericTaskSetup); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGenericTaskSetup 用id查询通用任务设置
// @Tags GenericTaskSetup
// @Summary 用id查询通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.GenericTaskSetup true "用id查询通用任务设置"
// @Success 200 {object} response.Response{data=object{regenericTaskSetup=octopus.GenericTaskSetup},msg=string} "查询成功"
// @Router /genericTaskSetup/findGenericTaskSetup [get]
func (genericTaskSetupApi *GenericTaskSetupApi) FindGenericTaskSetup(c *gin.Context) {
	ID := c.Query("ID")
	if regenericTaskSetup, err := genericTaskSetupService.GetGenericTaskSetup(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(regenericTaskSetup, c)
	}
}

// GetGenericTaskSetupList 分页获取通用任务设置列表
// @Tags GenericTaskSetup
// @Summary 分页获取通用任务设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.GenericTaskSetupSearch true "分页获取通用任务设置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /genericTaskSetup/getGenericTaskSetupList [get]
func (genericTaskSetupApi *GenericTaskSetupApi) GetGenericTaskSetupList(c *gin.Context) {
	var pageInfo octopusReq.GenericTaskSetupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.CreatedBy = utils.GetUserID(c)
	if list, total, err := genericTaskSetupService.GetGenericTaskSetupInfoList(pageInfo); err != nil {
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

func (genericTaskSetupApi *GenericTaskSetupApi) DeleteBindData(c *gin.Context) {
	setupId := c.Query("setupId")
	mainTaskType := c.Query("mainTaskType")

	if err := genericTaskSetupService.DeleteBindData(setupId, mainTaskType); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// GetGenericTaskSetupPublic 不需要鉴权的通用任务设置接口
// @Tags GenericTaskSetup
// @Summary 不需要鉴权的通用任务设置接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.GenericTaskSetupSearch true "分页获取通用任务设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /genericTaskSetup/getGenericTaskSetupPublic [get]
func (genericTaskSetupApi *GenericTaskSetupApi) GetGenericTaskSetupPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的通用任务设置接口信息",
	}, "获取成功", c)
}
