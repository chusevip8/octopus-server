import service from '@/utils/request'

// @Tags MsgTaskSetup
// @Summary 创建私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgTaskSetup true "创建私信任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /msgTaskSetup/createMsgTaskSetup [post]
export const createMsgTaskSetup = (data) => {
  return service({
    url: '/msgTaskSetup/createMsgTaskSetup',
    method: 'post',
    data
  })
}

// @Tags MsgTaskSetup
// @Summary 删除私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgTaskSetup true "删除私信任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msgTaskSetup/deleteMsgTaskSetup [delete]
export const deleteMsgTaskSetup = (params) => {
  return service({
    url: '/msgTaskSetup/deleteMsgTaskSetup',
    method: 'delete',
    params
  })
}

// @Tags MsgTaskSetup
// @Summary 批量删除私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除私信任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /msgTaskSetup/deleteMsgTaskSetup [delete]
export const deleteMsgTaskSetupByIds = (params) => {
  return service({
    url: '/msgTaskSetup/deleteMsgTaskSetupByIds',
    method: 'delete',
    params
  })
}

// @Tags MsgTaskSetup
// @Summary 更新私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MsgTaskSetup true "更新私信任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /msgTaskSetup/updateMsgTaskSetup [put]
export const updateMsgTaskSetup = (data) => {
  return service({
    url: '/msgTaskSetup/updateMsgTaskSetup',
    method: 'put',
    data
  })
}

// @Tags MsgTaskSetup
// @Summary 用id查询私信任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MsgTaskSetup true "用id查询私信任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /msgTaskSetup/findMsgTaskSetup [get]
export const findMsgTaskSetup = (params) => {
  return service({
    url: '/msgTaskSetup/findMsgTaskSetup',
    method: 'get',
    params
  })
}

// @Tags MsgTaskSetup
// @Summary 分页获取私信任务设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取私信任务设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /msgTaskSetup/getMsgTaskSetupList [get]
export const getMsgTaskSetupList = (params) => {
  return service({
    url: '/msgTaskSetup/getMsgTaskSetupList',
    method: 'get',
    params
  })
}
