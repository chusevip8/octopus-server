package octopus

import (
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
	db := global.GVA_DB.Model(&octopus.Comment{})
	var comments []octopus.Comment

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&comments).Error

	var commentResList []octopusRes.CommentRes
	for _, comment := range comments {
		text := octopusRes.Text{Text: comment.Content}
		commentRes := octopusRes.CommentRes{}
		commentRes.Name = comment.Commenter
		commentRes.Text = text
		commentRes.Date = comment.PostAt
		commentResList = append(commentResList, commentRes)
	}

	return commentResList, total, err
}
