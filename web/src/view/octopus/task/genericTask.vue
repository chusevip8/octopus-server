<template>
    <div v-loading.fullscreen.lock="fullscreenLoading">
        <div class="gva-search-box">
            <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
                :rules="searchRule" @keyup.enter="onSubmit">

                <el-form-item label="设备编号" prop="deviceNumber">
                    <el-input v-model="searchInfo.deviceNumber" placeholder="搜索条件" />
                </el-form-item>
                <el-form-item label="任务状态" prop="status">
                    <el-select v-model="searchInfo.status" clearable placeholder="请选择">
                        <el-option v-for="item in taskStatusOptions" :key="item.value" :label="`${item.label}`"
                            :value="item.value" />
                    </el-select>
                </el-form-item>

                <template v-if="showAllQuery">
                    <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
                </template>

                <el-form-item>
                    <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                    <el-button link type="primary" icon="arrow-down" @click="showAllQuery = true"
                        v-if="!showAllQuery">展开</el-button>
                    <el-button link type="primary" icon="arrow-up" @click="showAllQuery = false" v-else>收起</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button type="primary" icon="plus" @click="openDialog" style="margin-right: 20px;">添加设备</el-button>
                <el-button type="warning" icon="Switch" @click="bindData" style="margin-right: 20px;">绑定数据</el-button>
                <el-button type="success" icon="CaretRight" @click="startTasks">开始运行</el-button>
            </div>
            <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
                @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column align="center" label="设备编号" prop="device.number" width="120" />
                <el-table-column align="center" label="任务状态" width="120">
                    <template #default="scope">
                        <span>{{ taskStatusFilter(scope.row.status) }}</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" label="设备状态" width="120">
                    <template #default="scope">
                        <span>{{ deviceStatusFilter(scope.row.device.status) }}</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" label="参数" prop="taskParams.params" width="240" />
                <el-table-column align="center" label="错误信息" prop="error" min-width="120" />
                <el-table-column align="center" label="操作" fixed="right" min-width="240">
                    <template #default="scope">
                        <el-button type="primary" link icon="Close" class="table-button"
                            @click="stopTask(scope.row)">停止</el-button>
                        <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="gva-pagination">
                <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page"
                    :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
                    @current-change="handleCurrentChange" @size-change="handleSizeChange" />
            </div>
        </div>
        <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false"
            :before-close="closeDialog">
            <template #header>
                <div class="flex justify-between items-center">
                    <span class="text-lg">添加</span>
                    <div>
                        <el-button @click="closeDialog">关闭</el-button>
                    </div>
                </div>
            </template>
            <device-list @row-selected="saveTask" @group-selected="batchSaveTask" />
        </el-drawer>
    </div>
</template>

<script setup>
import {
    deleteTask,
    deleteTaskByIds,
    findTaskByDeviceId,
    getTaskList
} from '@/api/octopus/task'

import {
    createGenericTask,
    bindTaskData,
    startAllTasks
} from '@/api/octopus/genericTask'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, ReturnArrImg, onDownloadFile, getBaseUrl } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { DeviceList } from '@/view/octopus/components'
import { taskStatusOptions } from '@/view/octopus/utils/consts'
import { useRouter, useRoute } from 'vue-router'
import { deviceStatusOptions } from '@/view/octopus/utils/consts'

defineOptions({
    name: 'GenericTask'
})

const router = useRouter()
const route = useRoute()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
    batch: undefined,
    deviceGroup: undefined,
    taskSetupId: undefined,
    mainTaskType: undefined,
    subTaskType: undefined,
    deviceId: undefined,
    status: undefined,
    error: '',
})



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
    createdAt: [
        {
            validator: (rule, value, callback) => {
                if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
                    callback(new Error('请填写结束日期'))
                } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
                    callback(new Error('请填写开始日期'))
                } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
                    callback(new Error('开始日期应当早于结束日期'))
                } else {
                    callback()
                }
            }, trigger: 'change'
        }
    ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const fullscreenLoading = ref(false)
// 重置
const onReset = () => {
    searchInfo.value = {}
    getTableData()
}

const deviceStatusFilter = (value) => {
    const target = deviceStatusOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
}

const taskStatusFilter = (value) => {
    const target = taskStatusOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
}

// 搜索
const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
        if (!valid) return
        page.value = 1
        pageSize.value = 10
        getTableData()
    })
}

// 分页
const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}

// 查询
const getTableData = async () => {
    const table = await getTaskList({ page: page.value, pageSize: pageSize.value, taskSetupId: route.params.taskSetupId, mainTaskType: route.params.mainTaskType, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
// const deleteRow = (row) => {
//   ElMessageBox.confirm('确定要删除吗?', '提示', {
//     confirmButtonText: '确定',
//     cancelButtonText: '取消',
//     type: 'warning'
//   }).then(() => {
//     deleteTaskFunc(row)
//   })
// }

// 多选删除
const onDelete = async () => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(async () => {
        const IDs = []
        if (multipleSelection.value.length === 0) {
            ElMessage({
                type: 'warning',
                message: '请选择要删除的数据'
            })
            return
        }
        multipleSelection.value &&
            multipleSelection.value.map(item => {
                IDs.push(item.ID)
            })
        const res = await deleteTaskByIds({ IDs })
        if (res.code === 0) {
            ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === IDs.length && page.value > 1) {
                page.value--
            }
            getTableData()
        }
    })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        batch: undefined,
        deviceGroup: undefined,
        taskSetupId: undefined,
        mainTaskType: undefined,
        subTaskType: undefined,
        deviceId: undefined,
        status: undefined,
        error: '',
    }
}
const saveTask = async (params) => {
    formData.value.batch = false
    formData.value.deviceId = params.deviceId
    formData.value.mainTaskType = route.params.mainTaskType
    formData.value.subTaskType = route.params.subTaskType
    formData.value.taskSetupId = parseInt(route.params.taskSetupId)

    formData.value.status = 1

    const deviceStatus = params.deviceStatus

    if (deviceStatus === deviceStatusOptions.value[1].value) {
        ElMessage({
            type: 'error',
            message: '不能添加离线设备'
        })
    } else if (deviceStatus === deviceStatusOptions.value[2].value) {
        ElMessage({
            type: 'error',
            message: '不能添加禁用设备'
        })
    } else {
        let res = await findTaskByDeviceId({ taskSetupId: parseInt(route.params.taskSetupId), deviceId: params.deviceId, mainTaskType: route.params.mainTaskType })
        if (res.data !== null) {
            ElMessage({
                type: 'error',
                message: '该设备已添加'
            })
        } else {
            let res = await createGenericTask(formData.value)
            if (res.code === 0) {
                ElMessage({
                    type: 'success',
                    message: '创建/更改成功'
                })
                getTableData()
            }
        }
    }
}

const batchSaveTask = (group) => {
    const message = (typeof group === 'undefined' || group === null || group === '') ? '未选择分组，确定添加所有设备吗?' : '确定添加分组[' + group + ']所有设备吗？'
    ElMessageBox.confirm(message, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        fullscreenLoading.value = true
        batchSaveTaskFunc(group)
    })
}
const batchSaveTaskFunc = async (group) => {
    formData.value.batch = true
    formData.value.deviceGroup = group
    formData.value.deviceId = 0
    formData.value.mainTaskType = route.params.mainTaskType
    formData.value.subTaskType = route.params.subTaskType
    formData.value.taskSetupId = parseInt(route.params.taskSetupId)

    formData.value.status = 1

    let res = await createGenericTask(formData.value)
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '创建/更改成功'
        })
        getTableData()
    }
    fullscreenLoading.value = false
}

const bindData = () => {

    if (total.value == 0) {
        ElMessage({
            type: 'error',
            message: '未添加设备，不能绑定数据'
        })
        return
    }

    ElMessageBox.confirm('把所有数据绑定到已添加设备吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        fullscreenLoading.value = true
        bindTaskDataFunc()
    })
}

const bindTaskDataFunc = async () => {
    let res = await bindTaskData({ taskSetupId: route.params.taskSetupId, mainTaskType: route.params.mainTaskType, subTaskType: route.params.subTaskType })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '绑定成功'
        })
        getTableData()
    }
    fullscreenLoading.value = false
}

const startTasks = () => {
    if (total.value == 0) {
        ElMessage({
            type: 'error',
            message: '未添加设备，不能开始任务'
        })
        return
    }

    ElMessageBox.confirm('确定开始所有任务吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        startTasksFunc()
    })
}
const startTasksFunc = async () => {
    let res = await startAllTasks({ taskSetupId: route.params.taskSetupId, mainTaskType: route.params.mainTaskType, subTaskType: route.params.subTaskType })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '运行成功'
        })
        const title = route.meta.title;
        router.push({ name: "Reload", params: { title } });
    }
}

const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        deleteTaskFunc(row)
    })
}

const deleteTaskFunc = async (row) => {
    const res = await deleteTask({ id: row.ID })
    if (res.code === 0) {
        ElMessage({
            type: 'success',
            message: '删除成功'
        })
        if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}


// 停止任务
const stopTask = async (row) => {

}
</script>

<style></style>