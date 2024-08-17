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

type IntervalTaskApi struct {
}

var intervalTaskService = service.ServiceGroupApp.OctopusServiceGroup.IntervalTaskService

func (intervalTaskApi *IntervalTaskApi) CreateIntervalTask(c *gin.Context) {
	var intervalTask octopusReq.IntervalTask
	err := c.ShouldBindJSON(&intervalTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	intervalTask.CreatedBy = utils.GetUserID(c)
	if err := intervalTaskService.CreateIntervalTask(&intervalTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
func (intervalTaskApi *IntervalTaskApi) DeleteIntervalTask(c *gin.Context) {
	id := c.Query("id")
	userId := utils.GetUserID(c)
	if err := intervalTaskService.DeleteIntervalTask(id, userId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
func (intervalTaskApi *IntervalTaskApi) StopIntervalTask(c *gin.Context) {
	taskId := c.Query("taskId")
	if err := intervalTaskService.StopIntervalTask(taskId); err != nil {
		global.GVA_LOG.Error("停止任务失败!", zap.Error(err))
		response.FailWithMessage("停止任务失败", c)
	} else {
		response.OkWithMessage("停止任务成功", c)
	}
}
