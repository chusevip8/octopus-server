package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"gorm.io/gorm"
)

func filter(db *gorm.DB, userId uint) {
	user, err := system.UserServiceApp.FindUserById(int(userId))
	if err != nil {
		return
	}
	if user.AuthorityId != 1 && user.AuthorityId != 2 {
		db = db.Where("created_by = ?", userId)
	}
}

func isAdmin(userId uint) bool {
	user, err := system.UserServiceApp.FindUserById(int(userId))
	if err != nil {
		return false
	}
	return user.AuthorityId == 1 || user.AuthorityId == 2
}
