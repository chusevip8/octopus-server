package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CommentRouter struct {}

// InitCommentRouter 初始化 评论 路由信息
func (s *CommentRouter) InitCommentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	commentRouter := Router.Group("comment").Use(middleware.OperationRecord())
	commentRouterWithoutRecord := Router.Group("comment")
	commentRouterWithoutAuth := PublicRouter.Group("comment")

	var commentApi = v1.ApiGroupApp.OctopusApiGroup.CommentApi
	{
		commentRouter.POST("createComment", commentApi.CreateComment)   // 新建评论
		commentRouter.DELETE("deleteComment", commentApi.DeleteComment) // 删除评论
		commentRouter.DELETE("deleteCommentByIds", commentApi.DeleteCommentByIds) // 批量删除评论
		commentRouter.PUT("updateComment", commentApi.UpdateComment)    // 更新评论
	}
	{
		commentRouterWithoutRecord.GET("findComment", commentApi.FindComment)        // 根据ID获取评论
		commentRouterWithoutRecord.GET("getCommentList", commentApi.GetCommentList)  // 获取评论列表
	}
	{
	    commentRouterWithoutAuth.GET("getCommentPublic", commentApi.GetCommentPublic)  // 获取评论列表
	}
}
