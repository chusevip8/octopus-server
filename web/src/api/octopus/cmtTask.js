import service from '@/utils/request'

export const createFindCmtTask = (data) => {
    return service({
        url: '/cmtTask/createFindCmtTask',
        method: 'post',
        data
    })
}

export const createWriteCmtTask = (data) => {
    return service({
        url: '/cmtTask/createWriteCmtTask',
        method: 'post',
        data
    })
}

export const stopCmtTask = (params) => {
    return service({
        url: '/cmtTask/stopCmtTask',
        method: 'get',
        params
    })
}

export const stopCmtTasks = (data) => {
    return service({
        url: '/cmtTask/stopCmtTasks',
        method: 'post',
        data
    })
}
export const deleteCmtTasks = (data) => {
    return service({
        url: '/cmtTask/deleteCmtTasks',
        method: 'post',
        data
    })
}
export const deleteCmtTask = (params) => {
    return service({
        url: '/cmtTask/deleteCmtTask',
        method: 'delete',
        params
    })
}