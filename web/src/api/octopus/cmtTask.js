import service from '@/utils/request'

export const createFindCmtTask = (data) => {
    return service({
        url: '/cmtTask/createFindCmtTask',
        method: 'post',
        data
    })
}
export const findTaskByDeviceId = (params) => {
    return service({
        url: '/cmtTask/findTaskByDeviceId',
        method: 'get',
        params
    })
}

export const deleteCmtTask = (params) => {
    return service({
        url: '/cmtTask/deleteCmtTask',
        method: 'delete',
        params
    })
}