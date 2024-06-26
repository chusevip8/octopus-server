package initialize

import (
	"gorm.io/gorm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
)

func bizModel(db *gorm.DB) error {
	return db.AutoMigrate(octopus.Conversation{}, octopus.Thread{})
}
