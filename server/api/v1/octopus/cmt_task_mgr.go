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

type CmtTaskMgrApi struct {
}

var cmtTaskMgrService = service.ServiceGroupApp.OctopusServiceGroup.CmtTaskMgrService


// CreateCmtTaskMgr 创建评论任务管理
// @Tags CmtTaskMgr
// @Summary 创建评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskMgr true "创建评论任务管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cmtTaskMgr/createCmtTaskMgr [post]
func (cmtTaskMgrApi *CmtTaskMgrApi) CreateCmtTaskMgr(c *gin.Context) {
	var cmtTaskMgr octopus.CmtTaskMgr
	err := c.ShouldBindJSON(&cmtTaskMgr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cmtTaskMgr.CreatedBy = utils.GetUserID(c)

	if err := cmtTaskMgrService.CreateCmtTaskMgr(&cmtTaskMgr); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmtTaskMgr 删除评论任务管理
// @Tags CmtTaskMgr
// @Summary 删除评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskMgr true "删除评论任务管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cmtTaskMgr/deleteCmtTaskMgr [delete]
func (cmtTaskMgrApi *CmtTaskMgrApi) DeleteCmtTaskMgr(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := cmtTaskMgrService.DeleteCmtTaskMgr(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmtTaskMgrByIds 批量删除评论任务管理
// @Tags CmtTaskMgr
// @Summary 批量删除评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cmtTaskMgr/deleteCmtTaskMgrByIds [delete]
func (cmtTaskMgrApi *CmtTaskMgrApi) DeleteCmtTaskMgrByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := cmtTaskMgrService.DeleteCmtTaskMgrByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmtTaskMgr 更新评论任务管理
// @Tags CmtTaskMgr
// @Summary 更新评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CmtTaskMgr true "更新评论任务管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cmtTaskMgr/updateCmtTaskMgr [put]
func (cmtTaskMgrApi *CmtTaskMgrApi) UpdateCmtTaskMgr(c *gin.Context) {
	var cmtTaskMgr octopus.CmtTaskMgr
	err := c.ShouldBindJSON(&cmtTaskMgr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cmtTaskMgr.UpdatedBy = utils.GetUserID(c)

	if err := cmtTaskMgrService.UpdateCmtTaskMgr(cmtTaskMgr); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmtTaskMgr 用id查询评论任务管理
// @Tags CmtTaskMgr
// @Summary 用id查询评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.CmtTaskMgr true "用id查询评论任务管理"
// @Success 200 {object} response.Response{data=object{recmtTaskMgr=octopus.CmtTaskMgr},msg=string} "查询成功"
// @Router /cmtTaskMgr/findCmtTaskMgr [get]
func (cmtTaskMgrApi *CmtTaskMgrApi) FindCmtTaskMgr(c *gin.Context) {
	ID := c.Query("ID")
	if recmtTaskMgr, err := cmtTaskMgrService.GetCmtTaskMgr(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recmtTaskMgr, c)
	}
}

// GetCmtTaskMgrList 分页获取评论任务管理列表
// @Tags CmtTaskMgr
// @Summary 分页获取评论任务管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtTaskMgrSearch true "分页获取评论任务管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cmtTaskMgr/getCmtTaskMgrList [get]
func (cmtTaskMgrApi *CmtTaskMgrApi) GetCmtTaskMgrList(c *gin.Context) {
	var pageInfo octopusReq.CmtTaskMgrSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cmtTaskMgrService.GetCmtTaskMgrInfoList(pageInfo); err != nil {
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

// GetCmtTaskMgrPublic 不需要鉴权的评论任务管理接口
// @Tags CmtTaskMgr
// @Summary 不需要鉴权的评论任务管理接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CmtTaskMgrSearch true "分页获取评论任务管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cmtTaskMgr/getCmtTaskMgrPublic [get]
func (cmtTaskMgrApi *CmtTaskMgrApi) GetCmtTaskMgrPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的评论任务管理接口信息",
    }, "获取成功", c)
}
