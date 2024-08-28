package octopus

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	octopusRes "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/response"
)

type CommentService struct{}

// CreateComment 创建评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) CreateComment(comment *octopus.Comment) (err error) {
	err = global.GVA_DB.Create(comment).Error
	return err
}

// DeleteComment 删除评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) DeleteComment(ID string) (err error) {
	err = global.GVA_DB.Delete(&octopus.Comment{}, "id = ?", ID).Error
	return err
}

// DeleteCommentByIds 批量删除评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) DeleteCommentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]octopus.Comment{}, "id in ?", IDs).Error
	return err
}

// UpdateComment 更新评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) UpdateComment(comment octopus.Comment) (err error) {
	err = global.GVA_DB.Model(&octopus.Comment{}).Where("id = ?", comment.ID).Updates(&comment).Error
	return err
}

// GetComment 根据ID获取评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetComment(ID string) (comment octopus.Comment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&comment).Error
	return
}

// GetCommentInfoList 分页获取评论记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetCommentInfoList(info octopusReq.CommentSearch) (list []octopusRes.CommentRes, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.Comment{}).Preload("Task")
	var comments []octopus.Comment

	if info.ConversationId != "" {
		db = db.Where("conversation_id=?", info.ConversationId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&comments).Error

	if err != nil {
		return
	}

	err = commentService.markCommentsAsRead(comments)

	var commentResList []octopusRes.CommentRes
	for _, comment := range comments {

		commentRes := octopusRes.CommentRes{}
		commentRes.Name = comment.Commenter
		if comment.Mine {
			commentRes.Name = comment.CommentReplier
		}
		commentRes.Date = comment.PostAt
		commentRes.Mine = comment.Mine
		commentRes.Text = commentService.formatCommentText(comment)
		commentRes.Img = "avatar_blue.jpg"
		if comment.Mine {
			commentRes.Img = "avatar_red.jpg"
		}
		commentResList = append(commentResList, commentRes)
	}

	return commentResList, total, err
}
func (commentService *CommentService) formatCommentText(comment octopus.Comment) octopusRes.Text {
	if !comment.Mine {
		return octopusRes.Text{Text: comment.Content}
	}

	var color, status, taskErr string
	switch comment.Task.Status {
	case 0:
		color, status, taskErr = "red", "失败", "-任务已删除"
	case 1:
		color, status, taskErr = "yellow", "新建", ""
	case 2:
		color, status, taskErr = "yellow", "发布中", ""
	case 3:
		color, status, taskErr = "green", "已发布", ""
	default:
		color, status, taskErr = "red", "失败", "-"+comment.Task.Error
	}

	text := fmt.Sprintf(`%s<p style="color: %s;">%s%s</p>`, comment.Content, color, status, taskErr)
	return octopusRes.Text{Text: text}
}

func (commentService *CommentService) markCommentsAsRead(comments []octopus.Comment) error {
	if len(comments) == 0 {
		return nil
	}

	commentIds := make([]uint, len(comments))
	for i, comment := range comments {
		commentIds[i] = comment.ID
	}

	err := global.GVA_DB.Model(&octopus.Comment{}).Where("id IN (?) AND unread = ?", commentIds, true).Update("unread", false).Error
	return err
}
