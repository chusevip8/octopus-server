import service from '@/utils/request'

// @Tags TaskParams
// @Summary 创建任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskParams true "创建任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskParams/createTaskParams [post]
export const createTaskParams = (data) => {
  return service({
    url: '/taskParams/createTaskParams',
    method: 'post',
    data
  })
}

// @Tags TaskParams
// @Summary 删除任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskParams true "删除任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskParams/deleteTaskParams [delete]
export const deleteTaskParams = (params) => {
  return service({
    url: '/taskParams/deleteTaskParams',
    method: 'delete',
    params
  })
}

// @Tags TaskParams
// @Summary 批量删除任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskParams/deleteTaskParams [delete]
export const deleteTaskParamsByIds = (params) => {
  return service({
    url: '/taskParams/deleteTaskParamsByIds',
    method: 'delete',
    params
  })
}

// @Tags TaskParams
// @Summary 更新任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskParams true "更新任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskParams/updateTaskParams [put]
export const updateTaskParams = (data) => {
  return service({
    url: '/taskParams/updateTaskParams',
    method: 'put',
    data
  })
}

// @Tags TaskParams
// @Summary 用id查询任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TaskParams true "用id查询任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskParams/findTaskParams [get]
export const findTaskParams = (params) => {
  return service({
    url: '/taskParams/findTaskParams',
    method: 'get',
    params
  })
}

// @Tags TaskParams
// @Summary 分页获取任务参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取任务参数列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskParams/getTaskParamsList [get]
export const getTaskParamsList = (params) => {
  return service({
    url: '/taskParams/getTaskParamsList',
    method: 'get',
    params
  })
}
