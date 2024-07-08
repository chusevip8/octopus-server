import service from '@/utils/request'

// @Tags CmtTaskOption
// @Summary 创建评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskOption true "创建评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cmtTaskOption/createCmtTaskOption [post]
export const createCmtTaskOption = (data) => {
  return service({
    url: '/cmtTaskOption/createCmtTaskOption',
    method: 'post',
    data
  })
}

// @Tags CmtTaskOption
// @Summary 删除评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskOption true "删除评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtTaskOption/deleteCmtTaskOption [delete]
export const deleteCmtTaskOption = (params) => {
  return service({
    url: '/cmtTaskOption/deleteCmtTaskOption',
    method: 'delete',
    params
  })
}

// @Tags CmtTaskOption
// @Summary 批量删除评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmtTaskOption/deleteCmtTaskOption [delete]
export const deleteCmtTaskOptionByIds = (params) => {
  return service({
    url: '/cmtTaskOption/deleteCmtTaskOptionByIds',
    method: 'delete',
    params
  })
}

// @Tags CmtTaskOption
// @Summary 更新评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CmtTaskOption true "更新评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmtTaskOption/updateCmtTaskOption [put]
export const updateCmtTaskOption = (data) => {
  return service({
    url: '/cmtTaskOption/updateCmtTaskOption',
    method: 'put',
    data
  })
}

// @Tags CmtTaskOption
// @Summary 用id查询评论任务参数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CmtTaskOption true "用id查询评论任务参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmtTaskOption/findCmtTaskOption [get]
export const findCmtTaskOption = (params) => {
  return service({
    url: '/cmtTaskOption/findCmtTaskOption',
    method: 'get',
    params
  })
}

// @Tags CmtTaskOption
// @Summary 分页获取评论任务参数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取评论任务参数列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmtTaskOption/getCmtTaskOptionList [get]
export const getCmtTaskOptionList = (params) => {
  return service({
    url: '/cmtTaskOption/getCmtTaskOptionList',
    method: 'get',
    params
  })
}
