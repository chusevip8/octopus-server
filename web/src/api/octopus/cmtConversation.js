import service from '@/utils/request'

// @Tags CmtConversation
// @Summary 创建评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtConversation true "创建评论会话记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmtConversation/createCmtConversation [post]
export const createCmtConversation = (data) => {
  return service({
    url: '/cmtConversation/createCmtConversation',
    method: 'post',
    data
  })
}

// @Tags CmtConversation
// @Summary 删除评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtConversation true "删除评论会话记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtConversation/deleteCmtConversation [delete]
export const deleteCmtConversation = (params) => {
  return service({
    url: '/cmtConversation/deleteCmtConversation',
    method: 'delete',
    params
  })
}

// @Tags CmtConversation
// @Summary 批量删除评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论会话记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtConversation/deleteCmtConversation [delete]
export const deleteCmtConversationByIds = (params) => {
  return service({
    url: '/cmtConversation/deleteCmtConversationByIds',
    method: 'delete',
    params
  })
}

// @Tags CmtConversation
// @Summary 更新评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtConversation true "更新评论会话记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmtConversation/updateCmtConversation [put]
export const updateCmtConversation = (data) => {
  return service({
    url: '/cmtConversation/updateCmtConversation',
    method: 'put',
    data
  })
}

// @Tags CmtConversation
// @Summary 用id查询评论会话记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmtConversation true "用id查询评论会话记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmtConversation/findCmtConversation [get]
export const findCmtConversation = (params) => {
  return service({
    url: '/cmtConversation/findCmtConversation',
    method: 'get',
    params
  })
}

// @Tags CmtConversation
// @Summary 分页获取评论会话记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论会话记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmtConversation/getCmtConversationList [get]
export const getCmtConversationList = (params) => {
  return service({
    url: '/cmtConversation/getCmtConversationList',
    method: 'get',
    params
  })
}
