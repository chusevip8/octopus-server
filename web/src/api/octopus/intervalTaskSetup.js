import service from '@/utils/request'

// @Tags IntervalTaskSetup
// @Summary 创建间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntervalTaskSetup true "创建间隔任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /intervalTaskSetup/createIntervalTaskSetup [post]
export const createIntervalTaskSetup = (data) => {
  return service({
    url: '/intervalTaskSetup/createIntervalTaskSetup',
    method: 'post',
    data
  })
}

// @Tags IntervalTaskSetup
// @Summary 删除间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntervalTaskSetup true "删除间隔任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intervalTaskSetup/deleteIntervalTaskSetup [delete]
export const deleteIntervalTaskSetup = (params) => {
  return service({
    url: '/intervalTaskSetup/deleteIntervalTaskSetup',
    method: 'delete',
    params
  })
}

// @Tags IntervalTaskSetup
// @Summary 批量删除间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除间隔任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /intervalTaskSetup/deleteIntervalTaskSetup [delete]
export const deleteIntervalTaskSetupByIds = (params) => {
  return service({
    url: '/intervalTaskSetup/deleteIntervalTaskSetupByIds',
    method: 'delete',
    params
  })
}

// @Tags IntervalTaskSetup
// @Summary 更新间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IntervalTaskSetup true "更新间隔任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /intervalTaskSetup/updateIntervalTaskSetup [put]
export const updateIntervalTaskSetup = (data) => {
  return service({
    url: '/intervalTaskSetup/updateIntervalTaskSetup',
    method: 'put',
    data
  })
}

// @Tags IntervalTaskSetup
// @Summary 用id查询间隔任务设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IntervalTaskSetup true "用id查询间隔任务设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /intervalTaskSetup/findIntervalTaskSetup [get]
export const findIntervalTaskSetup = (params) => {
  return service({
    url: '/intervalTaskSetup/findIntervalTaskSetup',
    method: 'get',
    params
  })
}

// @Tags IntervalTaskSetup
// @Summary 分页获取间隔任务设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取间隔任务设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /intervalTaskSetup/getIntervalTaskSetupList [get]
export const getIntervalTaskSetupList = (params) => {
  return service({
    url: '/intervalTaskSetup/getIntervalTaskSetupList',
    method: 'get',
    params
  })
}
