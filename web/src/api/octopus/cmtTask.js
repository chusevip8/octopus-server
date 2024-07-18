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