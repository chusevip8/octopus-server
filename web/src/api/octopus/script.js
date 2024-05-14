import service from '@/utils/request'

// @Tags Script
// @Summary 创建脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Script true "创建脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /script/createScript [post]
export const createScript = (data) => {
  return service({
    url: '/script/createScript',
    method: 'post',
    data
  })
}

// @Tags Script
// @Summary 删除脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Script true "删除脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /script/deleteScript [delete]
export const deleteScript = (params) => {
  return service({
    url: '/script/deleteScript',
    method: 'delete',
    params
  })
}

// @Tags Script
// @Summary 批量删除脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /script/deleteScript [delete]
export const deleteScriptByIds = (params) => {
  return service({
    url: '/script/deleteScriptByIds',
    method: 'delete',
    params
  })
}

// @Tags Script
// @Summary 更新脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Script true "更新脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /script/updateScript [put]
export const updateScript = (data) => {
  return service({
    url: '/script/updateScript',
    method: 'put',
    data
  })
}

// @Tags Script
// @Summary 用id查询脚本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Script true "用id查询脚本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /script/findScript [get]
export const findScript = (params) => {
  return service({
    url: '/script/findScript',
    method: 'get',
    params
  })
}

// @Tags Script
// @Summary 分页获取脚本列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取脚本列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /script/getScriptList [get]
export const getScriptList = (params) => {
  return service({
    url: '/script/getScriptList',
    method: 'get',
    params
  })
}
