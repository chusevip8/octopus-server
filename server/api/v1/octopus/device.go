package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type DeviceApi struct {
}

var deviceService = service.ServiceGroupApp.OctopusServiceGroup.DeviceService

func (deviceApi *DeviceApi) CreateDevice(c *gin.Context) {

}
