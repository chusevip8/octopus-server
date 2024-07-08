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

type CmtTaskOptionApi struct {}

var cmtTaskOptionService = service.ServiceGroupApp.OctopusServiceGroup.CmtTaskOptionService


// CreateCmtTaskOption 创建评论任务参数
// @Tags CmtTaskOption
// @Summary 创建评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskOption true "创建评论任务参数"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cmtTaskOption/createCmtTaskOption [post]
func (cmtTaskOptionApi *CmtTaskOptionApi) CreateCmtTaskOption(c *gin.Context) {
	var cmtTaskOption octopus.CmtTaskOption
	err := c.ShouldBindJSON(&cmtTaskOption)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cmtTaskOption.CreatedBy = utils.GetUserID(c)

	if err := cmtTaskOptionService.CreateCmtTaskOption(&cmtTaskOption); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmtTaskOption 删除评论任务参数
// @Tags CmtTaskOption
// @Summary 删除评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskOption true "删除评论任务参数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cmtTaskOption/deleteCmtTaskOption [delete]
func (cmtTaskOptionApi *CmtTaskOptionApi) DeleteCmtTaskOption(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := cmtTaskOptionService.DeleteCmtTaskOption(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmtTaskOptionByIds 批量删除评论任务参数
// @Tags CmtTaskOption
// @Summary 批量删除评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cmtTaskOption/deleteCmtTaskOptionByIds [delete]
func (cmtTaskOptionApi *CmtTaskOptionApi) DeleteCmtTaskOptionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := cmtTaskOptionService.DeleteCmtTaskOptionByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmtTaskOption 更新评论任务参数
// @Tags CmtTaskOption
// @Summary 更新评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskOption true "更新评论任务参数"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cmtTaskOption/updateCmtTaskOption [put]
func (cmtTaskOptionApi *CmtTaskOptionApi) UpdateCmtTaskOption(c *gin.Context) {
	var cmtTaskOption octopus.CmtTaskOption
	err := c.ShouldBindJSON(&cmtTaskOption)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cmtTaskOption.UpdatedBy = utils.GetUserID(c)

	if err := cmtTaskOptionService.UpdateCmtTaskOption(cmtTaskOption); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmtTaskOption 用id查询评论任务参数
// @Tags CmtTaskOption
// @Summary 用id查询评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.CmtTaskOption true "用id查询评论任务参数"
// @Success 200 {object} response.Response{data=object{recmtTaskOption=octopus.CmtTaskOption},msg=string} "查询成功"
// @Router /cmtTaskOption/findCmtTaskOption [get]
func (cmtTaskOptionApi *CmtTaskOptionApi) FindCmtTaskOption(c *gin.Context) {
	ID := c.Query("ID")
	if recmtTaskOption, err := cmtTaskOptionService.GetCmtTaskOption(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recmtTaskOption, c)
	}
}

// GetCmtTaskOptionList 分页获取评论任务参数列表
// @Tags CmtTaskOption
// @Summary 分页获取评论任务参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtTaskOptionSearch true "分页获取评论任务参数列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cmtTaskOption/getCmtTaskOptionList [get]
func (cmtTaskOptionApi *CmtTaskOptionApi) GetCmtTaskOptionList(c *gin.Context) {
	var pageInfo octopusReq.CmtTaskOptionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmtTaskOptionService.GetCmtTaskOptionInfoList(pageInfo); err != nil {
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

// GetCmtTaskOptionPublic 不需要鉴权的评论任务参数接口
// @Tags CmtTaskOption
// @Summary 不需要鉴权的评论任务参数接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtTaskOptionSearch true "分页获取评论任务参数列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cmtTaskOption/getCmtTaskOptionPublic [get]
func (cmtTaskOptionApi *CmtTaskOptionApi) GetCmtTaskOptionPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的评论任务参数接口信息",
    }, "获取成功", c)
}
