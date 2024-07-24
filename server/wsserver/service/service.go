package service

import "github.com/flipped-aurora/gin-vue-admin/server/service"

var (
	DeviceService = service.ServiceGroupApp.OctopusServiceGroup.DeviceService
	TaskService   = service.ServiceGroupApp.OctopusServiceGroup.TaskService
	ScriptService = service.ServiceGroupApp.OctopusServiceGroup.ScriptService
)
