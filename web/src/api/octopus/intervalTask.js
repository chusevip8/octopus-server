import service from '@/utils/request'

export const createIntervalTask = (data) => {
    return service({
        url: '/intervalTask/createIntervalTask',
        method: 'post',
        data
    })
}
