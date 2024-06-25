import service from '@/utils/request'

// @Tags ExecTask
// @Summary 创建执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExecTask true "创建执行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /execTask/createExecTask [post]
export const createExecTask = (data) => {
  return service({
    url: '/execTask/createExecTask',
    method: 'post',
    data
  })
}

// @Tags ExecTask
// @Summary 删除执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExecTask true "删除执行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /execTask/deleteExecTask [delete]
export const deleteExecTask = (params) => {
  return service({
    url: '/execTask/deleteExecTask',
    method: 'delete',
    params
  })
}

// @Tags ExecTask
// @Summary 批量删除执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除执行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /execTask/deleteExecTask [delete]
export const deleteExecTaskByIds = (params) => {
  return service({
    url: '/execTask/deleteExecTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags ExecTask
// @Summary 更新执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExecTask true "更新执行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /execTask/updateExecTask [put]
export const updateExecTask = (data) => {
  return service({
    url: '/execTask/updateExecTask',
    method: 'put',
    data
  })
}

// @Tags ExecTask
// @Summary 用id查询执行任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExecTask true "用id查询执行任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /execTask/findExecTask [get]
export const findExecTask = (params) => {
  return service({
    url: '/execTask/findExecTask',
    method: 'get',
    params
  })
}

// @Tags ExecTask
// @Summary 分页获取执行任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取执行任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /execTask/getExecTaskList [get]
export const getExecTaskList = (params) => {
  return service({
    url: '/execTask/getExecTaskList',
    method: 'get',
    params
  })
}

export const findExecTaskByDeviceId = (params) => {
  return service({
    url: '/execTask/findExecTaskByDeviceId',
    method: 'get',
    params
  })
}
