import service from '@/utils/request'

// @Tags CmtThread
// @Summary 创建评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtThread true "创建评论会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmtThread/createCmtThread [post]
export const createCmtThread = (data) => {
  return service({
    url: '/cmtThread/createCmtThread',
    method: 'post',
    data
  })
}

// @Tags CmtThread
// @Summary 删除评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtThread true "删除评论会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtThread/deleteCmtThread [delete]
export const deleteCmtThread = (params) => {
  return service({
    url: '/cmtThread/deleteCmtThread',
    method: 'delete',
    params
  })
}

// @Tags CmtThread
// @Summary 批量删除评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtThread/deleteCmtThread [delete]
export const deleteCmtThreadByIds = (params) => {
  return service({
    url: '/cmtThread/deleteCmtThreadByIds',
    method: 'delete',
    params
  })
}

// @Tags CmtThread
// @Summary 更新评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtThread true "更新评论会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmtThread/updateCmtThread [put]
export const updateCmtThread = (data) => {
  return service({
    url: '/cmtThread/updateCmtThread',
    method: 'put',
    data
  })
}

// @Tags CmtThread
// @Summary 用id查询评论会话
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmtThread true "用id查询评论会话"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmtThread/findCmtThread [get]
export const findCmtThread = (params) => {
  return service({
    url: '/cmtThread/findCmtThread',
    method: 'get',
    params
  })
}

// @Tags CmtThread
// @Summary 分页获取评论会话列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论会话列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmtThread/getCmtThreadList [get]
export const getCmtThreadList = (params) => {
  return service({
    url: '/cmtThread/getCmtThreadList',
    method: 'get',
    params
  })
}
