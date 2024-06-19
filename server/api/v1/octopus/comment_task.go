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

type CommentTaskApi struct {
}

var commentTaskService = service.ServiceGroupApp.OctopusServiceGroup.CommentTaskService


// CreateCommentTask 创建评论任务
// @Tags CommentTask
// @Summary 创建评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CommentTask true "创建评论任务"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /commentTask/createCommentTask [post]
func (commentTaskApi *CommentTaskApi) CreateCommentTask(c *gin.Context) {
	var commentTask octopus.CommentTask
	err := c.ShouldBindJSON(&commentTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    commentTask.CreatedBy = utils.GetUserID(c)

	if err := commentTaskService.CreateCommentTask(&commentTask); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCommentTask 删除评论任务
// @Tags CommentTask
// @Summary 删除评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CommentTask true "删除评论任务"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /commentTask/deleteCommentTask [delete]
func (commentTaskApi *CommentTaskApi) DeleteCommentTask(c *gin.Context) {
	ID := c.Query("ID")
    	userID := utils.GetUserID(c)
	if err := commentTaskService.DeleteCommentTask(ID,userID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommentTaskByIds 批量删除评论任务
// @Tags CommentTask
// @Summary 批量删除评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /commentTask/deleteCommentTaskByIds [delete]
func (commentTaskApi *CommentTaskApi) DeleteCommentTaskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	if err := commentTaskService.DeleteCommentTaskByIds(IDs,userID); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCommentTask 更新评论任务
// @Tags CommentTask
// @Summary 更新评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.CommentTask true "更新评论任务"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /commentTask/updateCommentTask [put]
func (commentTaskApi *CommentTaskApi) UpdateCommentTask(c *gin.Context) {
	var commentTask octopus.CommentTask
	err := c.ShouldBindJSON(&commentTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    commentTask.UpdatedBy = utils.GetUserID(c)

	if err := commentTaskService.UpdateCommentTask(commentTask); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCommentTask 用id查询评论任务
// @Tags CommentTask
// @Summary 用id查询评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.CommentTask true "用id查询评论任务"
// @Success 200 {object} response.Response{data=object{recommentTask=octopus.CommentTask},msg=string} "查询成功"
// @Router /commentTask/findCommentTask [get]
func (commentTaskApi *CommentTaskApi) FindCommentTask(c *gin.Context) {
	ID := c.Query("ID")
	if recommentTask, err := commentTaskService.GetCommentTask(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recommentTask, c)
	}
}

// GetCommentTaskList 分页获取评论任务列表
// @Tags CommentTask
// @Summary 分页获取评论任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CommentTaskSearch true "分页获取评论任务列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /commentTask/getCommentTaskList [get]
func (commentTaskApi *CommentTaskApi) GetCommentTaskList(c *gin.Context) {
	var pageInfo octopusReq.CommentTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := commentTaskService.GetCommentTaskInfoList(pageInfo); err != nil {
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

// GetCommentTaskPublic 不需要鉴权的评论任务接口
// @Tags CommentTask
// @Summary 不需要鉴权的评论任务接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CommentTaskSearch true "分页获取评论任务列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /commentTask/getCommentTaskPublic [get]
func (commentTaskApi *CommentTaskApi) GetCommentTaskPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的评论任务接口信息",
    }, "获取成功", c)
}
