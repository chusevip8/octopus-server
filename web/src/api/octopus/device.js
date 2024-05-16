import service from '@/utils/request'


export const deleteDevice = (params) => {
    return service({
        url: '/device/deleteDevice',
        method: 'delete',
        params
    })
}

export const deleteDeviceByIds = (params) => {
    return service({
        url: '/device/deleteDeviceByIds',
        method: 'delete',
        params
    })
}

export const getDeviceList = (params) => {
    return service({
        url: '/device/getDeviceList',
        method: 'get',
        params
    })
}