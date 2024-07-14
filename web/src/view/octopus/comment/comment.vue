<template>
  <div>
    <div class="gva-table-box">
      <!-- <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
</el-table-column>

<el-table-column align="left" label="评论记录Id" prop="conversationId" width="120" />
<el-table-column align="left" label="评论内容" prop="content" width="120" />
<el-table-column align="left" label="评论状态" prop="status" width="120" />
<el-table-column align="left" label="操作" fixed="right" min-width="240">
  <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateCommentFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
</el-table-column>
</el-table> -->
      <JwChat :taleList="tableData" @enter="bindEnter" v-model="inputMsg" :showRightBox="false" scrollType="scroll"
        width="80%" height="750px" :toolConfig="chatTool" :config="chatConfig">
      </JwChat>
      <!-- <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div> -->
    </div>

  </div>
</template>

<script setup>
import {
  createComment,
  deleteComment,
  deleteCommentByIds,
  updateComment,
  findComment,
  getCommentList
} from '@/api/octopus/comment'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'Comment'
})

const route = useRoute()

const chatTool = ref({
  show: []
})

const chatConfig = ref({
  name: 'JwChat',
  dept: '描述'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  conversationId: undefined,
  content: '',
  status: undefined,
})


const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(50)
const tableData = ref([])

// 查询
const getTableData = async () => {
  const table = await getCommentList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    console.log(tableData.value)
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createComment(formData.value)
        break
      case 'update':
        res = await updateComment(formData.value)
        break
      default:
        res = await createComment(formData.value)
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

</script>

<style></style>
