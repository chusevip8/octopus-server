import service from '@/utils/request'

// @Tags CommentTask
// @Summary 创建评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentTask true "创建评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /commentTask/createCommentTask [post]
export const createCommentTask = (data) => {
  return service({
    url: '/commentTask/createCommentTask',
    method: 'post',
    data
  })
}

// @Tags CommentTask
// @Summary 删除评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentTask true "删除评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentTask/deleteCommentTask [delete]
export const deleteCommentTask = (params) => {
  return service({
    url: '/commentTask/deleteCommentTask',
    method: 'delete',
    params
  })
}

// @Tags CommentTask
// @Summary 批量删除评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /commentTask/deleteCommentTask [delete]
export const deleteCommentTaskByIds = (params) => {
  return service({
    url: '/commentTask/deleteCommentTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags CommentTask
// @Summary 更新评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommentTask true "更新评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /commentTask/updateCommentTask [put]
export const updateCommentTask = (data) => {
  return service({
    url: '/commentTask/updateCommentTask',
    method: 'put',
    data
  })
}

// @Tags CommentTask
// @Summary 用id查询评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CommentTask true "用id查询评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /commentTask/findCommentTask [get]
export const findCommentTask = (params) => {
  return service({
    url: '/commentTask/findCommentTask',
    method: 'get',
    params
  })
}

// @Tags CommentTask
// @Summary 分页获取评论任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /commentTask/getCommentTaskList [get]
export const getCommentTaskList = (params) => {
  return service({
    url: '/commentTask/getCommentTaskList',
    method: 'get',
    params
  })
}
