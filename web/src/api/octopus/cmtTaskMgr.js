import service from '@/utils/request'

// @Tags CmtTaskMgr
// @Summary 创建评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskMgr true "创建评论任务管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmtTaskMgr/createCmtTaskMgr [post]
export const createCmtTaskMgr = (data) => {
  return service({
    url: '/cmtTaskMgr/createCmtTaskMgr',
    method: 'post',
    data
  })
}

// @Tags CmtTaskMgr
// @Summary 删除评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskMgr true "删除评论任务管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtTaskMgr/deleteCmtTaskMgr [delete]
export const deleteCmtTaskMgr = (params) => {
  return service({
    url: '/cmtTaskMgr/deleteCmtTaskMgr',
    method: 'delete',
    params
  })
}

// @Tags CmtTaskMgr
// @Summary 批量删除评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论任务管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtTaskMgr/deleteCmtTaskMgr [delete]
export const deleteCmtTaskMgrByIds = (params) => {
  return service({
    url: '/cmtTaskMgr/deleteCmtTaskMgrByIds',
    method: 'delete',
    params
  })
}

// @Tags CmtTaskMgr
// @Summary 更新评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskMgr true "更新评论任务管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmtTaskMgr/updateCmtTaskMgr [put]
export const updateCmtTaskMgr = (data) => {
  return service({
    url: '/cmtTaskMgr/updateCmtTaskMgr',
    method: 'put',
    data
  })
}

// @Tags CmtTaskMgr
// @Summary 用id查询评论任务管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmtTaskMgr true "用id查询评论任务管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmtTaskMgr/findCmtTaskMgr [get]
export const findCmtTaskMgr = (params) => {
  return service({
    url: '/cmtTaskMgr/findCmtTaskMgr',
    method: 'get',
    params
  })
}

// @Tags CmtTaskMgr
// @Summary 分页获取评论任务管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论任务管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmtTaskMgr/getCmtTaskMgrList [get]
export const getCmtTaskMgrList = (params) => {
  return service({
    url: '/cmtTaskMgr/getCmtTaskMgrList',
    method: 'get',
    params
  })
}
