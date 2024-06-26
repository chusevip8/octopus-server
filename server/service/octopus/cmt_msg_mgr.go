package octopus

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
)

type CmtMsgMgrService struct{}

func (cmtMsgMgrService *CmtMsgMgrService) SaveCmtMeg(data string) (err error) {
	var cmtMsg request.CmtMsg
	err = json.Unmarshal([]byte(data), &cmtMsg)
	if err == nil {
		err = global.GVA_DB.Model(&octopus.Conversation{}).Where("taskID = ?", cmtMsg.TaskId).Updates(&conversation).Error
	}
}
