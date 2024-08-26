package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MsgTaskApi struct{}

var msgTaskService = service.ServiceGroupApp.OctopusServiceGroup.MsgTaskService

func (msgTaskApi *MsgTaskApi) UploadMessage(c *gin.Context) {
	var messageReq octopusReq.MessageReq
	err := c.ShouldBindJSON(&messageReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = msgTaskService.CreateMessage(&messageReq); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
