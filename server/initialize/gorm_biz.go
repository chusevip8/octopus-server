package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	"gorm.io/gorm"
)

func bizModel(db *gorm.DB) error {
	return db.AutoMigrate(octopus.Script{}, octopus.Device{}, octopus.CmtThread{}, octopus.CmtConversation{}, octopus.Comment{}, octopus.Task{}, octopus.CmtTaskSetup{}, octopus.TaskParams{}, octopus.IntervalTaskSetup{}, octopus.GenericTaskSetup{}, octopus.TaskBindData{})
}
