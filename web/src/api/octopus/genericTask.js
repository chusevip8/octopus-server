import service from '@/utils/request'

export const createGenericTask = (data) => {
    return service({
        url: '/genericTask/createGenericTask',
        method: 'post',
        data
    })
}
