package octopus

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type CmtTaskService struct{}

const (
	ErrorNoTask             = -1
	ErrorCreateThread       = -2
	ErrorCreateConversation = -3
	ErrorCreateComment      = -4
)

var CmtTaskServiceApp = new(CmtTaskService)

func (cmtTaskService *CmtTaskService) CreateWriteCmtTask(writeCmtTask *octopusReq.WriteCmtTask) (err error) {

	var cmtThread octopus.CmtThread
	cmtThread, err = CmtThreadServiceApp.GetCmtThread(writeCmtTask.ThreadId)
	if err != nil {
		return err
	}

	var cmtTaskSetup octopus.CmtTaskSetup
	cmtTaskSetup, err = CmtTaskSetupServiceApp.GetCmtTaskSetup(strconv.Itoa(int(cmtThread.TaskSetupId)))
	if err != nil {
		return err
	}
	var comment octopus.Comment
	err = global.GVA_DB.Model(&octopus.Comment{}).Preload("Task").
		Where("conversation_id=?", writeCmtTask.ConversationId).
		Where("mine=?", 0).
		Order("id desc").First(&comment).Error
	if err != nil {
		return err
	}
	params, err := cmtTaskService.buildWriteCmtTaskParams(cmtTaskSetup, cmtThread, comment, writeCmtTask.CmtContent)
	if err != nil {
		return err
	}
	var taskParams octopus.TaskParams
	taskParams.TaskSetupId = cmtTaskSetup.ID
	taskParams.CreatedBy = cmtTaskSetup.CreatedBy
	taskParams.MainTaskType = "cmt"
	scriptId := cmtTaskSetup.WriteCmtScriptId
	if comment.CmtFrom == "reply" {
		scriptId = cmtTaskSetup.ReplyCmtScriptId
	}
	taskParams.ScriptId = scriptId
	taskParams.SubTaskType = "writeCmt"
	taskParams.Params = params
	err = TaskParamsServiceApp.CreateTaskParams(&taskParams)
	if err != nil {
		return err
	}
	var task octopus.Task
	task.TaskParamsId = taskParams.ID
	task.AppName = cmtTaskSetup.AppName
	task.DeviceId = comment.Task.DeviceId
	task.CreatedBy = cmtTaskSetup.CreatedBy
	task.Status = 1
	task.Error = ""
	err = TaskServiceApp.CreateTask(&task)
	if err != nil {
		return err
	}

	commentReq := octopusReq.CommentReq{TaskId: strconv.Itoa(int(task.ID)),
		Poster:           cmtThread.Poster,
		PostTitle:        cmtThread.PostTitle,
		PostDesc:         cmtThread.PostDesc,
		Commenter:        comment.Commenter,
		CommenterId:      comment.CommenterId,
		CommentReplier:   comment.CommentReplier,
		CommentReplierId: comment.CommentReplierId,
		PostAt:           Today(),
		CmtFrom:          "write",
		Content:          writeCmtTask.CmtContent,
	}
	_, err = cmtTaskService.CreateComment(&commentReq)
	return err
}

func (cmtTaskService *CmtTaskService) CreateFindCmtTask(findCmtTask *octopusReq.FindCmtTask) (err error) {
	params, err := cmtTaskService.buildFindCmtTaskParams(findCmtTask.TaskSetupId)
	if err != nil {
		return err
	}

	var taskParams octopus.TaskParams
	taskParams.TaskSetupId = findCmtTask.TaskSetupId
	taskParams.ScriptId = findCmtTask.ScriptId
	taskParams.CreatedBy = findCmtTask.CreatedBy
	taskParams.MainTaskType = findCmtTask.MainTaskType
	taskParams.SubTaskType = findCmtTask.SubTaskType
	taskParams.Params = params
	err = TaskParamsServiceApp.CreateTaskParams(&taskParams)
	if err != nil {
		return err
	}

	var task octopus.Task
	task.TaskParamsId = taskParams.ID
	task.AppName = findCmtTask.AppName
	task.DeviceId = findCmtTask.DeviceId
	task.CreatedBy = findCmtTask.CreatedBy
	task.Status = findCmtTask.Status
	task.Error = findCmtTask.Error
	err = TaskServiceApp.CreateTask(&task)
	return err
}

func (cmtTaskService *CmtTaskService) buildWriteCmtTaskParams(cmtTaskSetup octopus.CmtTaskSetup, cmtThread octopus.CmtThread, comment octopus.Comment, cmtContent string) (params string, err error) {
	paramsMap := map[string]string{
		"postLink":         cmtTaskSetup.PostLink,
		"keyword":          cmtTaskSetup.Keyword,
		"poster":           cmtThread.Poster,
		"postTitle":        cmtThread.PostTitle,
		"postDesc":         cmtThread.PostDesc,
		"commenter":        comment.Commenter,
		"commenterId":      comment.CommenterId,
		"commentReplier":   comment.CommentReplier,
		"commentReplierId": comment.CommentReplierId,
		"inputText":        cmtContent,
	}
	jsonData, err := json.Marshal(paramsMap)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (cmtTaskService *CmtTaskService) buildFindCmtTaskParams(setupId uint) (params string, err error) {
	if cmtTaskSetup, err := CmtTaskSetupServiceApp.GetCmtTaskSetup(strconv.Itoa(int(setupId))); err != nil {
		return "", err
	} else {
		postLink := cmtTaskSetup.PostLink
		keyword := cmtTaskSetup.Keyword
		cmtCount := cmtTaskSetup.CmtCount
		params = fmt.Sprintf(`{"postLink": "%s", "keyword": "%s","cmtCount":"%s"}`, postLink, keyword, cmtCount)
		return params, nil
	}
}

func (cmtTaskService *CmtTaskService) UpdateFindCmtTaskParams(cmtTaskSetup octopus.CmtTaskSetup) (err error) {
	params := fmt.Sprintf(`{"postLink": "%s", "keyword": "%s"}`, cmtTaskSetup.PostLink, cmtTaskSetup.Keyword)
	err = global.GVA_DB.Model(&octopus.TaskParams{}).
		Where("task_setup_id = ?", cmtTaskSetup.ID).
		Where("main_task_type = ?", "cmt").
		Where("sub_task_type = ?", "findCmt").
		Update("params", params).Error
	return
}

func (cmtTaskService *CmtTaskService) buildPostId(poster string, postTitle string, postDesc string) (postId string, err error) {
	// 拼接字符串
	data := poster + postTitle + postDesc

	// 计算 MD5 哈希值
	hash := md5.New()
	_, err = hash.Write([]byte(data))
	if err != nil {
		return "", err
	}

	// 获取哈希值并编码为十六进制字符串
	postId = hex.EncodeToString(hash.Sum(nil))
	return postId, nil
}

func (cmtTaskService *CmtTaskService) buildCommenterId(commenter string) (commenterId string, err error) {
	hash := md5.New()
	_, err = hash.Write([]byte(commenter))
	if err != nil {
		return "", err
	}
	commenterId = hex.EncodeToString(hash.Sum(nil))
	return commenterId, nil
}

func (cmtTaskService *CmtTaskService) CreateComment(commentReq *octopusReq.CommentReq) (errCode int, err error) {
	var task octopus.Task
	err = global.GVA_DB.Preload("TaskParams").First(&task, commentReq.TaskId).Error
	if err != nil {
		return ErrorNoTask, err
	}
	var cmtThread octopus.CmtThread
	if errors.Is(global.GVA_DB.Model(&octopus.CmtThread{}).Where("task_setup_id=?", task.TaskParams.TaskSetupId).First(&cmtThread).Error, gorm.ErrRecordNotFound) {
		var postId string
		if postId, err = cmtTaskService.buildPostId(commentReq.Poster, commentReq.PostTitle, commentReq.PostDesc); err != nil {
			postId = strconv.Itoa(int(task.TaskParams.TaskSetupId))
		}
		cmtThread.AppName = task.AppName
		cmtThread.TaskSetupId = task.TaskParams.TaskSetupId
		cmtThread.Poster = commentReq.Poster
		cmtThread.PostTitle = commentReq.PostTitle
		cmtThread.PostDesc = commentReq.PostDesc
		cmtThread.CreatedBy = task.CreatedBy
		cmtThread.PostId = postId
		if err = global.GVA_DB.Create(&cmtThread).Error; err != nil {
			return ErrorCreateThread, err
		}

	}
	commenterId := commentReq.CommenterId
	commentReplierId := commentReq.CommentReplierId
	if commenterId == "" {
		commenterId, err = cmtTaskService.buildCommenterId(commentReq.Commenter)
	}
	if commentReplierId == "" {
		commentReplierId, err = cmtTaskService.buildCommenterId(commentReq.CommentReplier)
	}
	var cmtConversation octopus.CmtConversation
	err = global.GVA_DB.Model(&octopus.CmtConversation{}).
		Where("thread_id=?", cmtThread.ID).
		Where("commenter_id=?", commenterId).
		Where("comment_replier_id=?", commentReplierId).First(&cmtConversation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cmtConversation.ThreadId = cmtThread.ID
		cmtConversation.Commenter = commentReq.Commenter
		cmtConversation.CommenterId = commenterId
		cmtConversation.CommentReplier = commentReq.CommentReplier
		cmtConversation.CommentReplierId = commentReplierId
		cmtConversation.CreatedBy = cmtThread.CreatedBy
		if err = global.GVA_DB.Create(&cmtConversation).Error; err != nil {
			return ErrorCreateConversation, err
		}
	}
	var comment octopus.Comment
	comment.ConversationId = cmtConversation.ID
	comment.Commenter = commentReq.Commenter
	comment.CommenterId = commenterId
	comment.CommentReplier = commentReq.CommentReplier
	comment.CommentReplierId = commentReplierId
	comment.Content = commentReq.Content
	comment.PostAt = commentReq.PostAt
	comment.TaskId = task.ID
	comment.CmtFrom = commentReq.CmtFrom
	switch commentReq.CmtFrom {
	case "find", "reply":
		comment.Mine = false
		comment.Unread = true
	case "write":
		comment.Mine = true
		comment.Unread = false
	default:
		comment.Mine = false
		comment.Unread = false
	}

	if err = global.GVA_DB.Create(&comment).Error; err != nil {
		return ErrorCreateComment, err
	}
	return 0, nil
}
func (cmtTaskService *CmtTaskService) DeleteCmtTask(id string, userId uint) (err error) {
	var task octopus.Task
	err = global.GVA_DB.Preload("TaskParams").First(&task, id).Error
	if err != nil {
		return
	}
	var taskIds []uint
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ? AND oct_task.device_id = ?", task.TaskParams.TaskSetupId, task.TaskParams.MainTaskType, task.DeviceId).
		Pluck("oct_task.id", &taskIds).Error
	if err != nil {
		return
	}

	var taskParamsIds []uint
	err = global.GVA_DB.Model(&octopus.Task{}).
		Where("id in ?", taskIds).
		Pluck("oct_task.task_params_id", &taskParamsIds).Error
	if err != nil {
		return
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		if err := tx.Model(&octopus.Task{}).Where("id in ?", taskIds).Updates(map[string]interface{}{
			"deleted_by": userId,
			"deleted_at": now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&octopus.TaskParams{}).Where("id in ?", taskParamsIds).Updates(map[string]interface{}{
			"deleted_by": userId,
			"deleted_at": now,
		}).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}
func (cmtTaskService *CmtTaskService) StopCmtTask(id string) (err error) {
	var task octopus.Task
	err = global.GVA_DB.Preload("TaskParams").First(&task, id).Error
	if err != nil {
		return
	}
	tasks, err := TaskServiceApp.GetTasksByDeviceId(strconv.Itoa(int(task.TaskParams.TaskSetupId)), strconv.Itoa(int(task.DeviceId)), task.TaskParams.MainTaskType)
	if err != nil {
		return
	} else if len(tasks) == 0 {
		return fmt.Errorf("tasks not found to stop with task setup id %s", task.TaskParams.TaskSetupId)
	}
	for _, t := range tasks {
		StopTask <- &t
	}
	return nil
}
func (cmtTaskService *CmtTaskService) StopCmtTasks(taskSetup octopusReq.TaskSetup) (err error) {
	tasks, err := TaskServiceApp.GetTasksByTaskSetupId(taskSetup.TaskSetupId, taskSetup.MainTaskType, "")
	if err != nil {
		return err
	} else if len(tasks) == 0 {
		return fmt.Errorf("tasks not found to stop with task setup id %s", taskSetup.TaskSetupId)
	}
	for _, task := range tasks {
		StopTask <- &task
	}
	return nil
}
func (cmtTaskService *CmtTaskService) DeleteCmtTasks(taskSetup octopusReq.TaskSetup, userId uint) (err error) {
	var taskIds []uint
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ?", taskSetup.TaskSetupId, taskSetup.MainTaskType).
		Pluck("oct_task.id", &taskIds).Error
	if err != nil {
		return
	}
	var taskParamsIds []uint
	err = global.GVA_DB.Model(&octopus.Task{}).
		Where("id in ?", taskIds).
		Pluck("oct_task.task_params_id", &taskParamsIds).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		if err := tx.Model(&octopus.Task{}).Where("id in ?", taskIds).Updates(map[string]interface{}{
			"deleted_by": userId,
			"deleted_at": now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&octopus.TaskParams{}).Where("id in ?", taskParamsIds).Updates(map[string]interface{}{
			"deleted_by": userId,
			"deleted_at": now,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
