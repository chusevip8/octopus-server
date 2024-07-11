import service from '@/utils/request'

// @Tags CmtTaskSetup
// @Summary 创建评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskSetup true "创建评论任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmtTaskSetup/createCmtTaskSetup [post]
export const createCmtTaskSetup = (data) => {
  return service({
    url: '/cmtTaskSetup/createCmtTaskSetup',
    method: 'post',
    data
  })
}

// @Tags CmtTaskSetup
// @Summary 删除评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskSetup true "删除评论任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtTaskSetup/deleteCmtTaskSetup [delete]
export const deleteCmtTaskSetup = (params) => {
  return service({
    url: '/cmtTaskSetup/deleteCmtTaskSetup',
    method: 'delete',
    params
  })
}

// @Tags CmtTaskSetup
// @Summary 批量删除评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtTaskSetup/deleteCmtTaskSetup [delete]
export const deleteCmtTaskSetupByIds = (params) => {
  return service({
    url: '/cmtTaskSetup/deleteCmtTaskSetupByIds',
    method: 'delete',
    params
  })
}

// @Tags CmtTaskSetup
// @Summary 更新评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskSetup true "更新评论任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmtTaskSetup/updateCmtTaskSetup [put]
export const updateCmtTaskSetup = (data) => {
  return service({
    url: '/cmtTaskSetup/updateCmtTaskSetup',
    method: 'put',
    data
  })
}

// @Tags CmtTaskSetup
// @Summary 用id查询评论任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmtTaskSetup true "用id查询评论任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmtTaskSetup/findCmtTaskSetup [get]
export const findCmtTaskSetup = (params) => {
  return service({
    url: '/cmtTaskSetup/findCmtTaskSetup',
    method: 'get',
    params
  })
}

// @Tags CmtTaskSetup
// @Summary 分页获取评论任务设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论任务设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmtTaskSetup/getCmtTaskSetupList [get]
export const getCmtTaskSetupList = (params) => {
  return service({
    url: '/cmtTaskSetup/getCmtTaskSetupList',
    method: 'get',
    params
  })
}
