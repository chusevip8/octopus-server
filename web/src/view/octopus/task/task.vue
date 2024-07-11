<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">

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
        <el-button type="primary" icon="plus" @click="openDialog">添加</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="onDelete">删除</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="center" label="设备编号" prop="device.number" width="120" />
        <el-table-column align="center" label="设备备注" prop="device.note" width="360" />
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
        <el-table-column align="center" label="错误信息" prop="error" min-width="120" />
        <el-table-column align="center" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="stopTask(scope.row)">停止</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">添加</span>
          <div>
            <el-button @click="closeDialog">关闭</el-button>
          </div>
        </div>
      </template>
      <device-list @row-selected="saveTask" />
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createTask,
  deleteTask,
  deleteTaskByIds,
  updateTask,
  findTask,
  getTaskList
} from '@/api/octopus/task'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { DeviceList } from '@/view/octopus/components'
import { taskStatusOptions } from '@/view/octopus/utils/consts'
import { useRoute } from 'vue-router'
import { deviceStatusOptions } from '@/view/octopus/utils/consts'

defineOptions({
  name: 'Task'
})

const route = useRoute()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  appName: '',
  type: undefined,
  taskOptionId: undefined,
  deviceId: undefined,
  status: 1,
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
  const table = await getTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteTaskFunc(row)
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

// 更新行
const updateTaskFunc = async (row) => {
  const res = await findTask({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

// 停止任务
const stopTask = async (row) => {

}

// 删除行
const deleteTaskFunc = async (row) => {
  const res = await deleteTask({ ID: row.ID })
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
    appName: '',
    type: undefined,
    taskOptionId: undefined,
    deviceId: undefined,
    status: undefined,
    error: '',
  }
}
const saveTask = async (params) => {
  formData.value.appName = route.params.appName
  formData.value.type = route.params.taskType
  formData.value.deviceId = params.deviceId

  const deviceStatus = params.deviceStatus

  if (deviceStatus === deviceStatusOptions.value[2].value) {
    ElMessage({
      type: 'error',
      message: '不能添加离线设备'
    })
  } else if (deviceStatus === deviceStatusOptions.value[3].value) {
    ElMessage({
      type: 'error',
      message: '不能添加禁用设备'
    })
  } else {
    let res = await findExecTaskByDeviceId({ deviceID: params.deviceID })
    if (res.data !== null) {
      ElMessage({
        type: 'error',
        message: '该设备已添加'
      })
    } else {
      let res = await createExecTask(formData.value)
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

</script>

<style></style>
