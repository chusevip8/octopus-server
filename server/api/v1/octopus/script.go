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

type ScriptApi struct {
}

var scriptService = service.ServiceGroupApp.OctopusServiceGroup.ScriptService


// CreateScript 创建脚本
// @Tags Script
// @Summary 创建脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Script true "创建脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /script/createScript [post]
func (scriptApi *ScriptApi) CreateScript(c *gin.Context) {
	var script octopus.Script
	err := c.ShouldBindJSON(&script)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    script.CreatedBy = utils.GetUserID(c)

	if err := scriptService.CreateScript(&script); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteScript 删除脚本
// @Tags Script
// @Summary 删除脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Script true "删除脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /script/deleteScript [delete]
func (scriptApi *ScriptApi) DeleteScript(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := scriptService.DeleteScript(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteScriptByIds 批量删除脚本
// @Tags Script
// @Summary 批量删除脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /script/deleteScriptByIds [delete]
func (scriptApi *ScriptApi) DeleteScriptByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := scriptService.DeleteScriptByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateScript 更新脚本
// @Tags Script
// @Summary 更新脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Script true "更新脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /script/updateScript [put]
func (scriptApi *ScriptApi) UpdateScript(c *gin.Context) {
	var script octopus.Script
	err := c.ShouldBindJSON(&script)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    script.UpdatedBy = utils.GetUserID(c)

	if err := scriptService.UpdateScript(script); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindScript 用id查询脚本
// @Tags Script
// @Summary 用id查询脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.Script true "用id查询脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /script/findScript [get]
func (scriptApi *ScriptApi) FindScript(c *gin.Context) {
	ID := c.Query("ID")
	if rescript, err := scriptService.GetScript(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rescript": rescript}, c)
	}
}

// GetScriptList 分页获取脚本列表
// @Tags Script
// @Summary 分页获取脚本列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.ScriptSearch true "分页获取脚本列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /script/getScriptList [get]
func (scriptApi *ScriptApi) GetScriptList(c *gin.Context) {
	var pageInfo octopusReq.ScriptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := scriptService.GetScriptInfoList(pageInfo); err != nil {
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

// GetScriptPublic 不需要鉴权的脚本接口
// @Tags Script
// @Summary 不需要鉴权的脚本接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.ScriptSearch true "分页获取脚本列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /script/getScriptPublic [get]
func (scriptApi *ScriptApi) GetScriptPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的脚本接口信息",
    }, "获取成功", c)
}
