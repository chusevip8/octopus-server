package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CmtTaskApi struct{}

var cmtTaskService = service.ServiceGroupApp.OctopusServiceGroup.CmtTaskService

func (cmtTaskApi *CmtTaskApi) CreateFindCmtTask(c *gin.Context) {
	var findCmtTask octopusReq.FindCmtTask
	err := c.ShouldBindJSON(&findCmtTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	findCmtTask.CreatedBy = utils.GetUserID(c)
	if err := cmtTaskService.CreateFindCmtTask(&findCmtTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (cmtTaskApi *CmtTaskApi) CreateWriteCmtTask(c *gin.Context) {
	var writeCmtTask octopusReq.WriteCmtTask
	err := c.ShouldBindJSON(&writeCmtTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmtTaskService.CreateWriteCmtTask(&writeCmtTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (cmtTaskApi *CmtTaskApi) UploadFindComment(c *gin.Context) {
	var findCommentReq octopusReq.FindCommentReq
	err := c.ShouldBindJSON(&findCommentReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	commentReq := octopusReq.CommentReq{TaskId: findCommentReq.TaskId,
		Poster:           findCommentReq.Poster,
		PostTitle:        findCommentReq.PostTitle,
		PostDesc:         findCommentReq.PostDesc,
		CommentReplier:   findCommentReq.CommentReplier,
		CommentReplierId: findCommentReq.CommentReplierId,
		Content:          findCommentReq.Comments[0].Content,
		Commenter:        findCommentReq.Comments[0].Commenter,
		CommenterId:      findCommentReq.Comments[0].CommenterId,
		PostAt:           findCommentReq.Comments[0].PostAt,
		CmtFrom:          "find"}
	var errCode int
	if errCode, err = cmtTaskService.CreateComment(&commentReq); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCode(errCode, "创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (cmtTaskApi *CmtTaskApi) UploadMsgComment(c *gin.Context) {
	var msgCommentReq octopusReq.MsgCommentReq
	err := c.ShouldBindJSON(&msgCommentReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	postTitle := msgCommentReq.Comments[0].PostTitle
	if len(postTitle) > 150 {
		postTitle = postTitle[:150]
	}
	postDesc := msgCommentReq.Comments[0].PostDesc
	if len(postDesc) > 150 {
		postDesc = postDesc[:150]
	}
	commentReq := octopusReq.CommentReq{TaskId: msgCommentReq.TaskId,
		Poster:           msgCommentReq.Comments[0].Poster,
		PostTitle:        postTitle,
		PostDesc:         postDesc,
		CommentReplier:   msgCommentReq.CommentReplier,
		CommentReplierId: msgCommentReq.CommentReplierId,
		Content:          msgCommentReq.Comments[0].Content,
		Commenter:        msgCommentReq.Comments[0].Commenter,
		CommenterId:      msgCommentReq.Comments[0].CommenterId,
		PostAt:           msgCommentReq.Comments[0].PostAt,
		CmtFrom:          "msgCmt"}
	var errCode int
	if errCode, err = cmtTaskService.CreateComment(&commentReq); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCode(errCode, "创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (cmtTaskApi *CmtTaskApi) DeleteCmtTask(c *gin.Context) {
	id := c.Query("id")
	userId := utils.GetUserID(c)
	if err := cmtTaskService.DeleteCmtTask(id, userId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
func (cmtTaskApi *CmtTaskApi) DeleteCmtTasks(c *gin.Context) {
	var taskSetup octopusReq.TaskSetup
	err := c.ShouldBindJSON(&taskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if err := cmtTaskService.DeleteCmtTasks(taskSetup, userId); err != nil {
		global.GVA_LOG.Error("删除任务失败!", zap.Error(err))
		response.FailWithMessage("删除任务失败", c)
	} else {
		response.OkWithMessage("删除任务成功", c)
	}
}
func (cmtTaskApi *CmtTaskApi) StopCmtTask(c *gin.Context) {
	id := c.Query("taskId")
	if err := cmtTaskService.StopCmtTask(id); err != nil {
		global.GVA_LOG.Error("停止失败!", zap.Error(err))
		response.FailWithMessage("停止失败", c)
	} else {
		response.OkWithMessage("停止成功", c)
	}
}
func (cmtTaskApi *CmtTaskApi) StopCmtTasks(c *gin.Context) {
	var taskSetup octopusReq.TaskSetup
	err := c.ShouldBindJSON(&taskSetup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := cmtTaskService.StopCmtTasks(taskSetup); err != nil {
		global.GVA_LOG.Error("停止任务失败!", zap.Error(err))
		response.FailWithMessage("停止任务失败", c)
	} else {
		response.OkWithMessage("停止任务成功", c)
	}
}
