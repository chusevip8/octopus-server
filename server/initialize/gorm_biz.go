package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	"gorm.io/gorm"
)

func bizModel(db *gorm.DB) error {
	return db.AutoMigrate(octopus.Script{}, octopus.Device{}, octopus.CommentTask{}, octopus.ExecTask{}, octopus.Conversation{}, octopus.Thread{}, octopus.Message{}, octopus.CmtTaskOption{}, octopus.CmtThread{}, octopus.CmtConversation{}, octopus.Comment{}, octopus.Task{})
}
