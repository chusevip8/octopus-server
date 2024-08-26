import service from '@/utils/request'

// @Tags MsgConversation
// @Summary 创建私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgConversation true "创建私信会话纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msgConversation/createMsgConversation [post]
export const createMsgConversation = (data) => {
  return service({
    url: '/msgConversation/createMsgConversation',
    method: 'post',
    data
  })
}

// @Tags MsgConversation
// @Summary 删除私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgConversation true "删除私信会话纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msgConversation/deleteMsgConversation [delete]
export const deleteMsgConversation = (params) => {
  return service({
    url: '/msgConversation/deleteMsgConversation',
    method: 'delete',
    params
  })
}

// @Tags MsgConversation
// @Summary 批量删除私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除私信会话纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msgConversation/deleteMsgConversation [delete]
export const deleteMsgConversationByIds = (params) => {
  return service({
    url: '/msgConversation/deleteMsgConversationByIds',
    method: 'delete',
    params
  })
}

// @Tags MsgConversation
// @Summary 更新私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgConversation true "更新私信会话纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msgConversation/updateMsgConversation [put]
export const updateMsgConversation = (data) => {
  return service({
    url: '/msgConversation/updateMsgConversation',
    method: 'put',
    data
  })
}

// @Tags MsgConversation
// @Summary 用id查询私信会话纪录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MsgConversation true "用id查询私信会话纪录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msgConversation/findMsgConversation [get]
export const findMsgConversation = (params) => {
  return service({
    url: '/msgConversation/findMsgConversation',
    method: 'get',
    params
  })
}

// @Tags MsgConversation
// @Summary 分页获取私信会话纪录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取私信会话纪录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msgConversation/getMsgConversationList [get]
export const getMsgConversationList = (params) => {
  return service({
    url: '/msgConversation/getMsgConversationList',
    method: 'get',
    params
  })
}
