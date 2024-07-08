import service from '@/utils/request'

// @Tags Comment
// @Summary 创建评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "创建评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /comment/createComment [post]
export const createComment = (data) => {
  return service({
    url: '/comment/createComment',
    method: 'post',
    data
  })
}

// @Tags Comment
// @Summary 删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "删除评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comment/deleteComment [delete]
export const deleteComment = (params) => {
  return service({
    url: '/comment/deleteComment',
    method: 'delete',
    params
  })
}

// @Tags Comment
// @Summary 批量删除评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comment/deleteComment [delete]
export const deleteCommentByIds = (params) => {
  return service({
    url: '/comment/deleteCommentByIds',
    method: 'delete',
    params
  })
}

// @Tags Comment
// @Summary 更新评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Comment true "更新评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comment/updateComment [put]
export const updateComment = (data) => {
  return service({
    url: '/comment/updateComment',
    method: 'put',
    data
  })
}

// @Tags Comment
// @Summary 用id查询评论
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Comment true "用id查询评论"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comment/findComment [get]
export const findComment = (params) => {
  return service({
    url: '/comment/findComment',
    method: 'get',
    params
  })
}

// @Tags Comment
// @Summary 分页获取评论列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comment/getCommentList [get]
export const getCommentList = (params) => {
  return service({
    url: '/comment/getCommentList',
    method: 'get',
    params
  })
}
