import service from '@/utils/request'
export const createFindCmtTask = (data) => {
    return service({
        url: '/task/createFindCmtTask',
        method: 'post',
        data
    })
}