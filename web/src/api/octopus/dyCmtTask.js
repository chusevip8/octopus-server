import service from '@/utils/request'

// @Tags DYCmtTask
// @Summary 创建抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DYCmtTask true "创建抖音评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dyCmtTask/createDYCmtTask [post]
export const createDYCmtTask = (data) => {
  return service({
    url: '/dyCmtTask/createDYCmtTask',
    method: 'post',
    data
  })
}

// @Tags DYCmtTask
// @Summary 删除抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DYCmtTask true "删除抖音评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dyCmtTask/deleteDYCmtTask [delete]
export const deleteDYCmtTask = (params) => {
  return service({
    url: '/dyCmtTask/deleteDYCmtTask',
    method: 'delete',
    params
  })
}

// @Tags DYCmtTask
// @Summary 批量删除抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除抖音评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dyCmtTask/deleteDYCmtTask [delete]
export const deleteDYCmtTaskByIds = (params) => {
  return service({
    url: '/dyCmtTask/deleteDYCmtTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags DYCmtTask
// @Summary 更新抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DYCmtTask true "更新抖音评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dyCmtTask/updateDYCmtTask [put]
export const updateDYCmtTask = (data) => {
  return service({
    url: '/dyCmtTask/updateDYCmtTask',
    method: 'put',
    data
  })
}

// @Tags DYCmtTask
// @Summary 用id查询抖音评论任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DYCmtTask true "用id查询抖音评论任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dyCmtTask/findDYCmtTask [get]
export const findDYCmtTask = (params) => {
  return service({
    url: '/dyCmtTask/findDYCmtTask',
    method: 'get',
    params
  })
}

// @Tags DYCmtTask
// @Summary 分页获取抖音评论任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取抖音评论任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dyCmtTask/getDYCmtTaskList [get]
export const getDYCmtTaskList = (params) => {
  return service({
    url: '/dyCmtTask/getDYCmtTaskList',
    method: 'get',
    params
  })
}
