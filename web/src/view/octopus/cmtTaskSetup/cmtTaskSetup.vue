<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">

        <el-form-item label="任务标题" prop="taskTitle">
          <el-input v-model="searchInfo.taskTitle" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="评论关键字" prop="keyword">
          <el-input v-model="searchInfo.keyword" placeholder="搜索条件" />

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
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />

        <el-table-column align="center" label="任务标题" prop="taskTitle" width="180" />
        <el-table-column align="center" label="帖子链接" prop="postLink" width="120" />
        <el-table-column align="center" label="评论关键字" prop="keyword" min-width="240" />
        <el-table-column align="center" label="评论条数" prop="cmtCount" width="120" />
        <el-table-column align="center" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link icon="Cellphone" class="table-button"
              @click="openTaskManager(scope.row)">管理任务</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateCmtTaskSetupFunc(scope.row)">修改</el-button>
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
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="任务标题:" prop="taskTitle">
          <el-input v-model="formData.taskTitle" :clearable="false" placeholder="请输入任务标题" />
        </el-form-item>
        <el-form-item label="帖子链接:" prop="postLink">
          <el-input v-model="formData.postLink" :clearable="false" placeholder="请输入帖子链接" />
        </el-form-item>
        <el-form-item label="评论关键字:（以半角逗号分隔，空为所有评论）" prop="keyword">
          <el-input v-model="formData.keyword" :clearable="false" placeholder="请输入评论关键字" />
        </el-form-item>
        <el-form-item label="评论数量:（如果抓取所有评论，数值要大于总评论数）" prop="cmtCount">
          <el-input v-model="formData.cmtCount" :clearable="false" placeholder="请输入评论数量" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createCmtTaskSetup,
  deleteCmtTaskSetup,
  deleteCmtTaskSetupByIds,
  updateCmtTaskSetup,
  findCmtTaskSetup,
  getCmtTaskSetupList
} from '@/api/octopus/cmtTaskSetup'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'

defineOptions({
  name: 'CmtTaskSetup'
})

const router = useRouter()
const route = useRoute()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  appName: '',
  taskTitle: '',
  postLink: '',
  keyword: '',
  cmtCount: '9999',
  readPostCmtScriptId: undefined,
  replyPostCmtScriptId: undefined,
  replyMsgCmtScriptId: undefined,
})



// 验证规则
const rule = reactive({
  taskTitle: [{
    required: true,
    message: '任务标题不能为空',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  postLink: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  cmtCount: [{
    required: true,
    message: '评论数不能为空',
    trigger: ['input', 'blur'],
  },
  {
    pattern: /^\d+$/,
    message: '评论数只能输入数字',
    trigger: ['input', 'blur'],
  }]
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
  const table = await getCmtTaskSetupList({ page: page.value, pageSize: pageSize.value, appName: route.params.appName, ...searchInfo.value })
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
    deleteCmtTaskSetupFunc(row)
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
    const res = await deleteCmtTaskSetupByIds({ IDs })
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
const updateCmtTaskSetupFunc = async (row) => {
  const res = await findCmtTaskSetup({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteCmtTaskSetupFunc = async (row) => {
  const res = await deleteCmtTaskSetup({ ID: row.ID })
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

const openTaskManager = (row) => {
  router.push({ name: 'cmtTask', params: { appName: route.params.appName, mainTaskType: 'cmt', subTaskType: 'readPostCmt', taskSetupId: row.ID, scriptId: parseInt(route.params.readPostCmtScriptId, 10) } })
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
    taskTitle: '',
    postLink: '',
    keyword: '',
    cmtCount: '9999',
    readPostCmtScriptId: undefined,
    replyPostCmtScriptId: undefined,
    replyMsgCmtScriptId: undefined,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return

    formData.value.appName = route.params.appName
    formData.value.readPostCmtScriptId = parseInt(route.params.readPostCmtScriptId, 10)
    formData.value.replyPostCmtScriptId = parseInt(route.params.replyPostCmtScriptId, 10)
    formData.value.replyMsgCmtScriptId = parseInt(route.params.replyMsgCmtScriptId, 10)
    formData.value.keyword = formData.value.keyword.trim()
    formData.value.postLink = formData.value.postLink.trim()
    let res
    switch (type.value) {
      case 'create':
        res = await createCmtTaskSetup(formData.value)
        break
      case 'update':
        res = await updateCmtTaskSetup(formData.value)
        break
      default:
        res = await createCmtTaskSetup(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}
watch(() => route.path, (newPath, oldPath) => {
  if (newPath !== oldPath) {
    const title = route.meta.title;
    router.push({ name: "Reload", params: { title } });
  }
})
</script>

<style></style>
