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

type GenericTaskApi struct {
}

var genericTaskService = service.ServiceGroupApp.OctopusServiceGroup.GenericTaskService

func (genericTaskApi *GenericTaskApi) CreateGenericTask(c *gin.Context) {
	var genericTask octopusReq.GenericTask
	err := c.ShouldBindJSON(&genericTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	genericTask.CreatedBy = utils.GetUserID(c)
	if err := genericTaskService.CreateGenericTask(genericTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (genericTaskApi *GenericTaskApi) BindTaskData(c *gin.Context) {
	setupId := c.Request.FormValue("setupId")
	mainTaskType := c.Request.FormValue("mainTaskType")
	subTaskType := c.Request.FormValue("subTaskType")
	if err := genericTaskService.BindTaskData(setupId, mainTaskType, subTaskType); err != nil {
		global.GVA_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithMessage("绑定失败", c)
	} else {
		response.OkWithMessage("绑定成功", c)
	}
}
