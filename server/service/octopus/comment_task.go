package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "gorm.io/gorm"
)

type CommentTaskService struct {
}

// CreateCommentTask 创建评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentTaskService *CommentTaskService) CreateCommentTask(commentTask *octopus.CommentTask) (err error) {
	err = global.GVA_DB.Create(commentTask).Error
	return err
}

// DeleteCommentTask 删除评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentTaskService *CommentTaskService)DeleteCommentTask(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.CommentTask{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&octopus.CommentTask{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCommentTaskByIds 批量删除评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentTaskService *CommentTaskService)DeleteCommentTaskByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.CommentTask{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&octopus.CommentTask{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCommentTask 更新评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentTaskService *CommentTaskService)UpdateCommentTask(commentTask octopus.CommentTask) (err error) {
	err = global.GVA_DB.Model(&octopus.CommentTask{}).Where("id = ?",commentTask.ID).Updates(&commentTask).Error
	return err
}

// GetCommentTask 根据ID获取评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentTaskService *CommentTaskService)GetCommentTask(ID string) (commentTask octopus.CommentTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&commentTask).Error
	return
}

// GetCommentTaskInfoList 分页获取评论任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentTaskService *CommentTaskService)GetCommentTaskInfoList(info octopusReq.CommentTaskSearch) (list []octopus.CommentTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&octopus.CommentTask{})
    var commentTasks []octopus.CommentTask
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Keyword != "" {
        db = db.Where("keyword LIKE ?","%"+ info.Keyword+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&commentTasks).Error
	return  commentTasks, total, err
}