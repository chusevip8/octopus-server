import service from '@/utils/request'

export const createIntervalTask = (data) => {
    return service({
        url: '/intervalTask/createIntervalTask',
        method: 'post',
        data
    })
}
export const deleteIntervalTask = (params) => {
    return service({
        url: '/intervalTask/deleteIntervalTask',
        method: 'delete',
        params
    })
}