package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CommentTaskRouter struct {
}

// InitCommentTaskRouter 初始化 评论任务 路由信息
func (s *CommentTaskRouter) InitCommentTaskRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	commentTaskRouter := Router.Group("commentTask").Use(middleware.OperationRecord())
	commentTaskRouterWithoutRecord := Router.Group("commentTask")
	commentTaskRouterWithoutAuth := PublicRouter.Group("commentTask")

	var commentTaskApi = v1.ApiGroupApp.OctopusApiGroup.CommentTaskApi
	{
		commentTaskRouter.POST("createCommentTask", commentTaskApi.CreateCommentTask)   // 新建评论任务
		commentTaskRouter.DELETE("deleteCommentTask", commentTaskApi.DeleteCommentTask) // 删除评论任务
		commentTaskRouter.DELETE("deleteCommentTaskByIds", commentTaskApi.DeleteCommentTaskByIds) // 批量删除评论任务
		commentTaskRouter.PUT("updateCommentTask", commentTaskApi.UpdateCommentTask)    // 更新评论任务
	}
	{
		commentTaskRouterWithoutRecord.GET("findCommentTask", commentTaskApi.FindCommentTask)        // 根据ID获取评论任务
		commentTaskRouterWithoutRecord.GET("getCommentTaskList", commentTaskApi.GetCommentTaskList)  // 获取评论任务列表
	}
	{
	    commentTaskRouterWithoutAuth.GET("getCommentTaskPublic", commentTaskApi.GetCommentTaskPublic)  // 获取评论任务列表
	}
}
