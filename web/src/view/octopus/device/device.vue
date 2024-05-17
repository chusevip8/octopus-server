<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
                @keyup.enter="onSubmit">
                <el-form-item label="编号" prop="number">
                    <el-input v-model="searchInfo.number" placeholder="搜索条件" />
                </el-form-item>
                <el-form-item label="用户" prop="nickName">
                    <el-input v-model="searchInfo.nickName" placeholder="搜索条件" />
                </el-form-item>
                <el-form-item label="备注" prop="note">
                    <el-input v-model="searchInfo.note" placeholder="搜索条件" />
                </el-form-item>
                <el-form-item label="状态" prop="status">
                    <el-select v-model="searchInfo.status" clearable placeholder="请选择">
                        <el-option v-for="item in statusOptions" :key="item.value" :label="`${item.label}`"
                            :value="item.value" />
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                    <el-button icon="refresh" @click="onReset">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="gva-table-box">
            <div class="gva-btn-list">
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
                    @click="onDelete">删除</el-button>
            </div>
            <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
                @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column align="center" label="编号" prop="number" width="120" />
                <el-table-column align="center" label="备注" prop="note" min-width="240" />
                <el-table-column align="center" label="状态" fixed="right" width="120">
                    <template #default="scope">
                        <span>{{ statusFilter(scope.row.status) }}</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" label="用户" prop="user.nickName" width="240" />

                <el-table-column align="center" label="操作" fixed="right" width="240">
                    <template #default="scope">
                        <el-button type="primary" link icon="stop" class="table-button"
                            @click="updateScriptFunc(scope.row)">停止</el-button>
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
    </div>
</template>

<script setup>
import {

    deleteDevice,
    deleteDeviceByIds,
    getDeviceList,

} from '@/api/octopus/device'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

defineOptions({
    name: 'Device'
})


const elSearchFormRef = ref()
const statusOptions = ref([
    {
        value: 1,
        label: '运行'
    },
    {
        value: 2,
        label: '就绪'
    },
    {
        value: 3,
        label: '离线'
    },
    {
        value: 4,
        label: '禁用'
    }
])

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
    searchInfo.value = {}
    getTableData()
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


const statusFilter = (value) => {
    const target = statusOptions.value.filter(item => item.value === value)[0]
    return target && `${target.label}`
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
    const table = await getDeviceList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize
    }
}

getTableData()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        deleteDeviceFunc(row)
    })
}

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
        const res = await deleteDeviceByIds({ IDs })
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

// 删除行
const deleteDeviceFunc = async (row) => {
    const res = await deleteDevice({ ID: row.ID })
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
</script>

<style></style>