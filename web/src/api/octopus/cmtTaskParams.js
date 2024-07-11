import service from '@/utils/request'

// @Tags CmtTaskParams
// @Summary 创建评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskParams true "创建评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmtTaskParams/createCmtTaskParams [post]
export const createCmtTaskParams = (data) => {
  return service({
    url: '/cmtTaskParams/createCmtTaskParams',
    method: 'post',
    data
  })
}

// @Tags CmtTaskParams
// @Summary 删除评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskParams true "删除评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtTaskParams/deleteCmtTaskParams [delete]
export const deleteCmtTaskParams = (params) => {
  return service({
    url: '/cmtTaskParams/deleteCmtTaskParams',
    method: 'delete',
    params
  })
}

// @Tags CmtTaskParams
// @Summary 批量删除评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtTaskParams/deleteCmtTaskParams [delete]
export const deleteCmtTaskParamsByIds = (params) => {
  return service({
    url: '/cmtTaskParams/deleteCmtTaskParamsByIds',
    method: 'delete',
    params
  })
}

// @Tags CmtTaskParams
// @Summary 更新评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskParams true "更新评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmtTaskParams/updateCmtTaskParams [put]
export const updateCmtTaskParams = (data) => {
  return service({
    url: '/cmtTaskParams/updateCmtTaskParams',
    method: 'put',
    data
  })
}

// @Tags CmtTaskParams
// @Summary 用id查询评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmtTaskParams true "用id查询评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmtTaskParams/findCmtTaskParams [get]
export const findCmtTaskParams = (params) => {
  return service({
    url: '/cmtTaskParams/findCmtTaskParams',
    method: 'get',
    params
  })
}

// @Tags CmtTaskParams
// @Summary 分页获取评论任务参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论任务参数列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmtTaskParams/getCmtTaskParamsList [get]
export const getCmtTaskParamsList = (params) => {
  return service({
    url: '/cmtTaskParams/getCmtTaskParamsList',
    method: 'get',
    params
  })
}
