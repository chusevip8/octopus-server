import service from '@/utils/request'

export const createReadPostCmtTask = (data) => {
    return service({
        url: '/cmtTask/createReadPostCmtTask',
        method: 'post',
        data
    })
}

export const createReplyCmtTask = (data) => {
    return service({
        url: '/cmtTask/createReplyCmtTask',
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