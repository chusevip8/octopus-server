import service from '@/utils/request'

export const createGenericTask = (data) => {
    return service({
        url: '/genericTask/createGenericTask',
        method: 'post',
        data
    })
}

export const bindTaskData = (data) => {
    return service({
        url: '/genericTask/bindTaskData',
        method: 'post',
        data
    })
}

export const startGenericTasks = (data) => {
    return service({
        url: '/genericTask/startGenericTasks',
        method: 'post',
        data
    })
}

export const stopGenericTask = (params) => {
    return service({
        url: '/genericTask/stopGenericTask',
        method: 'get',
        params
    })
}

export const stopGenericTasks = (data) => {
    return service({
        url: '/genericTask/stopGenericTasks',
        method: 'post',
        data
    })
}
export const deleteGenericTasks = (data) => {
    return service({
        url: '/genericTask/deleteGenericTasks',
        method: 'post',
        data
    })
}