<template>
    <div>
        <div class="gva-search-box">
            <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
                @keyup.enter="onSubmit">
                <el-form-item label="编号" prop="number">
                    <el-input v-model="searchInfo.number" placeholder="搜索条件" />
                </el-form-item>
                <el-form-item label="分组" prop="group">
                    <el-input v-model="searchInfo.group" placeholder="搜索条件" />
                </el-form-item>
                <el-form-item label="备注" prop="note">
                    <el-input v-model="searchInfo.note" placeholder="搜索条件" />
                </el-form-item>
                <el-form-item label="状态" prop="status">
                    <el-select v-model="searchInfo.status" clearable placeholder="请选择">
                        <el-option v-for="item in deviceStatusOptions" :key="item.value" :label="`${item.label}`"
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
            <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
                @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column align="center" label="编号" prop="number" width="120" />
                <el-table-column align="center" label="分组" prop="group" width="160" />
                <el-table-column align="center" label="状态" width="120">
                    <template #default="scope">
                        <span>{{ statusFilter(scope.row.status) }}</span>
                    </template>
                </el-table-column>
                <el-table-column align="center" label="备注" prop="note" min-width="200" />

                <el-table-column align="center" label="操作" fixed="right" width="120">
                    <template #default="scope">
                        <el-button type="primary" link icon="Select" @click="rowSelected(scope.row)">选择</el-button>
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
    getDeviceList,
} from '@/api/octopus/device'

import { ref } from 'vue'
import { deviceStatusOptions } from '@/view/octopus/utils/consts'

const emit = defineEmits(['row-selected'])

const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})


const rowSelected = (row) => {
    emit('row-selected', { deviceId: row.ID, deviceStatus: row.status })
}

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
    const target = deviceStatusOptions.value.filter(item => item.value === value)[0]
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

const multipleSelection = ref([])

const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

</script>

<style></style>