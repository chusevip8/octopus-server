import service from '@/utils/request'

// @Tags TaskBindData
// @Summary 创建任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskBindData true "创建任务数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /taskBindData/createTaskBindData [post]
export const createTaskBindData = (data) => {
  return service({
    url: '/taskBindData/createTaskBindData',
    method: 'post',
    data
  })
}

// @Tags TaskBindData
// @Summary 删除任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskBindData true "删除任务数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskBindData/deleteTaskBindData [delete]
export const deleteTaskBindData = (params) => {
  return service({
    url: '/taskBindData/deleteTaskBindData',
    method: 'delete',
    params
  })
}

// @Tags TaskBindData
// @Summary 批量删除任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除任务数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /taskBindData/deleteTaskBindData [delete]
export const deleteTaskBindDataByIds = (params) => {
  return service({
    url: '/taskBindData/deleteTaskBindDataByIds',
    method: 'delete',
    params
  })
}

// @Tags TaskBindData
// @Summary 更新任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TaskBindData true "更新任务数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /taskBindData/updateTaskBindData [put]
export const updateTaskBindData = (data) => {
  return service({
    url: '/taskBindData/updateTaskBindData',
    method: 'put',
    data
  })
}

// @Tags TaskBindData
// @Summary 用id查询任务数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TaskBindData true "用id查询任务数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /taskBindData/findTaskBindData [get]
export const findTaskBindData = (params) => {
  return service({
    url: '/taskBindData/findTaskBindData',
    method: 'get',
    params
  })
}

// @Tags TaskBindData
// @Summary 分页获取任务数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取任务数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /taskBindData/getTaskBindDataList [get]
export const getTaskBindDataList = (params) => {
  return service({
    url: '/taskBindData/getTaskBindDataList',
    method: 'get',
    params
  })
}
