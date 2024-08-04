import { ref } from 'vue'

export const deviceStatusOptions = ref([
    {
        value: 1,
        label: '就绪'
    },
    {
        value: 2,
        label: '离线'
    },
    {
        value: 3,
        label: '禁用'
    }
])

export const taskStatusOptions = ref([
    {
        value: 1,
        label: '新建'
    },
    {
        value: 2,
        label: '执行中'
    },
    {
        value: 3,
        label: '完成'
    },
    {
        value: 4,
        label: '失败'
    }
])

export const taskBindDataOptions = ref([
    {
        value: 1,
        label: '新建'
    },
    {
        value: 2,
        label: '已绑定'
    }
    // ,
    // {
    //     value: 3,
    //     label: '执行中'
    // },
    // {
    //     value: 4,
    //     label: '完成'
    // },
    // {
    //     value: 5,
    //     label: '失败'
    // }
])