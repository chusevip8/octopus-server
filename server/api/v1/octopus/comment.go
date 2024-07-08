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

type CommentApi struct {}

var commentService = service.ServiceGroupApp.OctopusServiceGroup.CommentService


// CreateComment 创建评论
// @Tags Comment
// @Summary 创建评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Comment true "创建评论"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /comment/createComment [post]
func (commentApi *CommentApi) CreateComment(c *gin.Context) {
	var comment octopus.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := commentService.CreateComment(&comment); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteComment 删除评论
// @Tags Comment
// @Summary 删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Comment true "删除评论"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /comment/deleteComment [delete]
func (commentApi *CommentApi) DeleteComment(c *gin.Context) {
	ID := c.Query("ID")
	if err := commentService.DeleteComment(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommentByIds 批量删除评论
// @Tags Comment
// @Summary 批量删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /comment/deleteCommentByIds [delete]
func (commentApi *CommentApi) DeleteCommentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := commentService.DeleteCommentByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateComment 更新评论
// @Tags Comment
// @Summary 更新评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body octopus.Comment true "更新评论"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /comment/updateComment [put]
func (commentApi *CommentApi) UpdateComment(c *gin.Context) {
	var comment octopus.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := commentService.UpdateComment(comment); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindComment 用id查询评论
// @Tags Comment
// @Summary 用id查询评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopus.Comment true "用id查询评论"
// @Success 200 {object} response.Response{data=object{recomment=octopus.Comment},msg=string} "查询成功"
// @Router /comment/findComment [get]
func (commentApi *CommentApi) FindComment(c *gin.Context) {
	ID := c.Query("ID")
	if recomment, err := commentService.GetComment(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(recomment, c)
	}
}

// GetCommentList 分页获取评论列表
// @Tags Comment
// @Summary 分页获取评论列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CommentSearch true "分页获取评论列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /comment/getCommentList [get]
func (commentApi *CommentApi) GetCommentList(c *gin.Context) {
	var pageInfo octopusReq.CommentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := commentService.GetCommentInfoList(pageInfo); err != nil {
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

// GetCommentPublic 不需要鉴权的评论接口
// @Tags Comment
// @Summary 不需要鉴权的评论接口
// @accept application/json
// @Produce application/json
// @Param data query octopusReq.CommentSearch true "分页获取评论列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /comment/getCommentPublic [get]
func (commentApi *CommentApi) GetCommentPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的评论接口信息",
    }, "获取成功", c)
}
