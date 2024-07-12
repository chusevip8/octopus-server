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

func (cmtTaskApi *CmtTaskApi) FindTaskByDeviceId(c *gin.Context) {
	deviceId := c.Query("deviceId")
	taskSetupId := c.Query("taskSetupId")
	if retask, err := cmtTaskService.GetTaskByDeviceId(taskSetupId, deviceId); err != nil {
		response.OkWithData(nil, c)
	} else {
		response.OkWithData(retask, c)
	}
}

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
