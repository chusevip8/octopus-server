import service from '@/utils/request'

// @Tags Conversation
// @Summary 创建消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Conversation true "创建消息会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /conversation/createConversation [post]
export const createConversation = (data) => {
  return service({
    url: '/conversation/createConversation',
    method: 'post',
    data
  })
}

// @Tags Conversation
// @Summary 删除消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Conversation true "删除消息会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /conversation/deleteConversation [delete]
export const deleteConversation = (params) => {
  return service({
    url: '/conversation/deleteConversation',
    method: 'delete',
    params
  })
}

// @Tags Conversation
// @Summary 批量删除消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除消息会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /conversation/deleteConversation [delete]
export const deleteConversationByIds = (params) => {
  return service({
    url: '/conversation/deleteConversationByIds',
    method: 'delete',
    params
  })
}

// @Tags Conversation
// @Summary 更新消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Conversation true "更新消息会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /conversation/updateConversation [put]
export const updateConversation = (data) => {
  return service({
    url: '/conversation/updateConversation',
    method: 'put',
    data
  })
}

// @Tags Conversation
// @Summary 用id查询消息会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Conversation true "用id查询消息会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /conversation/findConversation [get]
export const findConversation = (params) => {
  return service({
    url: '/conversation/findConversation',
    method: 'get',
    params
  })
}

// @Tags Conversation
// @Summary 分页获取消息会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取消息会话列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /conversation/getConversationList [get]
export const getConversationList = (params) => {
  return service({
    url: '/conversation/getConversationList',
    method: 'get',
    params
  })
}
