import service from '@/utils/request'

// @Tags GenericTaskSetup
// @Summary 创建通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GenericTaskSetup true "创建通用任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /genericTaskSetup/createGenericTaskSetup [post]
export const createGenericTaskSetup = (data) => {
  return service({
    url: '/genericTaskSetup/createGenericTaskSetup',
    method: 'post',
    data
  })
}

// @Tags GenericTaskSetup
// @Summary 删除通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GenericTaskSetup true "删除通用任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /genericTaskSetup/deleteGenericTaskSetup [delete]
export const deleteGenericTaskSetup = (params) => {
  return service({
    url: '/genericTaskSetup/deleteGenericTaskSetup',
    method: 'delete',
    params
  })
}

// @Tags GenericTaskSetup
// @Summary 批量删除通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除通用任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /genericTaskSetup/deleteGenericTaskSetup [delete]
export const deleteGenericTaskSetupByIds = (params) => {
  return service({
    url: '/genericTaskSetup/deleteGenericTaskSetupByIds',
    method: 'delete',
    params
  })
}

// @Tags GenericTaskSetup
// @Summary 更新通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GenericTaskSetup true "更新通用任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /genericTaskSetup/updateGenericTaskSetup [put]
export const updateGenericTaskSetup = (data) => {
  return service({
    url: '/genericTaskSetup/updateGenericTaskSetup',
    method: 'put',
    data
  })
}

// @Tags GenericTaskSetup
// @Summary 用id查询通用任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GenericTaskSetup true "用id查询通用任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /genericTaskSetup/findGenericTaskSetup [get]
export const findGenericTaskSetup = (params) => {
  return service({
    url: '/genericTaskSetup/findGenericTaskSetup',
    method: 'get',
    params
  })
}

// @Tags GenericTaskSetup
// @Summary 分页获取通用任务设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取通用任务设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /genericTaskSetup/getGenericTaskSetupList [get]
export const getGenericTaskSetupList = (params) => {
  return service({
    url: '/genericTaskSetup/getGenericTaskSetupList',
    method: 'get',
    params
  })
}

export const deleteBindData = (params) => {
  return service({
    url: '/genericTaskSetup/deleteBindData',
    method: 'delete',
    params
  })
}