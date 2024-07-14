package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
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
func (commentService *CommentService) GetCommentInfoList(info octopusReq.CommentSearch) (list []octopus.Comment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.Comment{})
	var comments []octopus.Comment
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&comments).Error
	return comments, total, err
}
