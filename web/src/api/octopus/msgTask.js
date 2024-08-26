import service from '@/utils/request'

export const createReplyMsgTask = (data) => {
    return service({
        url: '/cmtTask/createReplyCmtTask',
        method: 'post',
        data
    })
}