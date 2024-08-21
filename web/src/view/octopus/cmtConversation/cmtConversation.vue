<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">

        <el-form-item label="发评论者" prop="commenter">
          <el-input v-model="searchInfo.commenter" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="评论回复者" prop="commentReplier">
          <el-input v-model="searchInfo.commentReplier" placeholder="搜索条件" />

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
      <!-- <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="onDelete">删除</el-button>
      </div> -->
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="center" label="发评论者" prop="commenter" min-width="240" />
        <el-table-column align="center" label="评论回复者" prop="commentReplier" min-width="240" />
        <el-table-column align="center" label="未读" prop="unreadCount" min-width="120">
          <template #default="scope">
            <span :style="{ color: scope.row.unreadCount !== 0 ? 'red' : '' }">
              {{ scope.row.unreadCount }}
            </span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link icon="View" class="table-button"
              @click="viewConversation(scope.row)">查看</el-button>
            <!-- <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button> -->
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
          <span class="text-lg">评论</span>
          <div>
            <el-button @click="closeDialog">关闭</el-button>
          </div>
        </div>
      </template>
      <comment :thread-id="threadId" :conversation-id="conversationId"></comment>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getCmtConversationList
} from '@/api/octopus/cmtConversation'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { Comment } from '@/view/octopus/components'

defineOptions({
  name: 'CmtConversation'
})

const route = useRoute()
// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

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
  const table = await getCmtConversationList({ page: page.value, pageSize: pageSize.value, threadId: route.params.threadId, ...searchInfo.value })
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


const threadId = ref('')
const conversationId = ref('')

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  threadId.value = ''
  conversationId.value = ''
  getTableData()
}


const viewConversation = (row) => {
  threadId.value = route.params.threadId
  conversationId.value = row.ID
  dialogFormVisible.value = true
}

</script>

<style></style>
