package octopus

import "github.com/flipped-aurora/gin-vue-admin/server/service"

var (
	DeviceService = service.ServiceGroupApp.OctopusServiceGroup.DeviceService
	TaskService   = service.ServiceGroupApp.OctopusServiceGroup.TaskService
)
