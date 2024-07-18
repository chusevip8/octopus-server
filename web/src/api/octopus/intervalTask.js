import service from '@/utils/request'

export const createIntervalTask = (data) => {
    return service({
        url: '/intervalTask/createFindCmtTask',
        method: 'post',
        data
    })
}
export const deleteIntervalTask = (params) => {
    return service({
        url: '/intervalTask/deleteCmtTask',
        method: 'delete',
        params
    })
}