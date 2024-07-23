package socket

import "github.com/flipped-aurora/gin-vue-admin/server/service"

var (
	deviceService = service.ServiceGroupApp.OctopusServiceGroup.DeviceService
	taskService   = service.ServiceGroupApp.OctopusServiceGroup.TaskService
)
