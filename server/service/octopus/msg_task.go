package octopus

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type MsgTaskService struct{}

func (msgTaskService *MsgTaskService) CreateMessage(messageReq *octopusReq.MessageReq) (err error) {
	var task octopus.Task
	err = global.GVA_DB.First(&task, messageReq.TaskId).Error
	if err != nil {
		return
	}
	senderId := messageReq.SenderId
	receiverId := messageReq.ReceiverId

	if senderId == "" {
		senderId, err = msgTaskService.buildWriterId(messageReq.Sender)
	}
	if receiverId == "" {
		receiverId, err = msgTaskService.buildWriterId(messageReq.Receiver)
	}

	var msgConversation octopus.MsgConversation
	err = global.GVA_DB.Model(&octopus.MsgConversation{}).
		Where("app_name=?", task.AppName).
		Where("sender_id=?", senderId).
		Where("receiver_id=?", receiverId).First(&msgConversation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		msgConversation.AppName = task.AppName
		msgConversation.Sender = messageReq.Sender
		msgConversation.Receiver = messageReq.Receiver
		msgConversation.SenderId = senderId
		msgConversation.ReceiverId = receiverId
		msgConversation.CreatedBy = task.CreatedBy
		if err = global.GVA_DB.Create(&msgConversation).Error; err != nil {
			return
		}
	}
	var message octopus.Message
	global.GVA_DB.Model(&octopus.Message{}).
		Where("sender_id = ? AND receiver_id = ? AND content = ?",
			senderId, receiverId, messageReq.Messages[0].Content).First(&message)
	message.TaskId = task.ID
	message.ConversationId = msgConversation.ID
	writerId := messageReq.Messages[0].WriterId
	if writerId == "" {
		writerId, err = msgTaskService.buildWriterId(messageReq.Messages[0].Writer)
	}
	message.Sender = messageReq.Sender
	message.Receiver = messageReq.Receiver
	message.SenderId = senderId
	message.ReceiverId = receiverId
	message.Content = messageReq.Messages[0].Content
	message.SendAt = messageReq.Messages[0].SendAt
	if writerId == senderId {
		message.Unread = true
		message.Mine = false
	} else {
		message.Unread = false
		message.Mine = true
	}
	err = global.GVA_DB.Save(&message).Error
	return
}

func (msgTaskService *MsgTaskService) buildWriterId(writer string) (writerId string, err error) {
	hash := md5.New()
	_, err = hash.Write([]byte(writer))
	if err != nil {
		return "", err
	}
	writerId = hex.EncodeToString(hash.Sum(nil))
	return writerId, nil
}
func (msgTaskService *MsgTaskService) CreateReplyMsgTask(replyMsgTask *octopusReq.ReplyMsgTask) (err error) {
	var msgConversation octopus.MsgConversation
	msgConversation, err = MsgConversationServiceApp.GetMsgConversation(replyMsgTask.ConversationId)
	if err != nil {
		return
	}
	var msgTaskSetup octopus.MsgTaskSetup
	msgTaskSetup, err = MsgTaskSetupServiceApp.GetMsgTaskSetupByAppName(replyMsgTask.AppName)
	if err != nil {
		return
	}
	var params string
	params, err = msgTaskService.buildReplyMsgTaskParams(msgConversation, replyMsgTask.MsgContent)
	if err != nil {
		return err
	}
	var taskParams octopus.TaskParams
	taskParams.TaskSetupId = msgTaskSetup.ID
	taskParams.CreatedBy = msgConversation.CreatedBy
	taskParams.MainTaskType = "msg"
	taskParams.ScriptId = msgTaskSetup.ScriptId
	taskParams.SubTaskType = "replyMsg"
	taskParams.Params = params
	err = TaskParamsServiceApp.CreateTaskParams(&taskParams)
	if err != nil {
		return err
	}
	var task octopus.Task
	task.TaskParamsId = taskParams.ID
	task.AppName = msgTaskSetup.AppName
	//task.DeviceId = comment.Task.DeviceId
	task.CreatedBy = msgConversation.CreatedBy
	task.Status = 1
	task.Error = ""
	err = TaskServiceApp.CreateTask(&task)
	if err != nil {
		return err
	}
	return nil
}

func (msgTaskService *MsgTaskService) buildReplyMsgTaskParams(msgConversation octopus.MsgConversation, msgContent string) (params string, err error) {
	paramsMap := map[string]string{
		"userName":        msgConversation.Sender,
		"userId":          msgConversation.SenderId,
		"writeMsgContent": msgContent,
	}
	jsonData, err := json.Marshal(paramsMap)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
