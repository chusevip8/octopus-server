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

export const startAllTasks = (data) => {
    return service({
        url: '/genericTask/startAllTasks',
        method: 'post',
        data
    })
}