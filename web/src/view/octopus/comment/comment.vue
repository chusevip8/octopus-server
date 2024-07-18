<template>
  <div>
    <JwChat :taleList="tableData" @enter="onSubmit" v-model="formData.cmtContent" :showRightBox="false"
      scrollType="scroll" width="80%" height="750px" :toolConfig="chatTool" :config="chatConfig">
    </JwChat>
  </div>
</template>

<script setup>
import {
  getCommentList
} from '@/api/octopus/comment'

import {
  findCmtThread
} from '@/api/octopus/cmtThread'

import {
  createWriteCmtTask
} from '@/api/octopus/cmtTask'

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
  name: '',
  dept: ''
})

const getPostInfo = async () => {
  const cmtThread = await findCmtThread({ ID: route.params.threadId })
  if (cmtThread.code == 0) {
    chatConfig.value.name = cmtThread.data.poster
    chatConfig.value.dept = cmtThread.data.postDesc
  }
}
getPostInfo()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  conversationId: '',
  threadId: '',
  cmtContent: '',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(50)
const tableData = ref([])

// 查询
const getTableData = async () => {
  const table = await getCommentList({ page: page.value, pageSize: pageSize.value, conversationId: route.params.conversationId })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

const onSubmit = async () => {
  if (formData.value.cmtContent.trim() === '') {
    return
  }
  const text = formData.value.cmtContent
  ElMessageBox.confirm('确定要提交吗?提交后评论不能修改和删除。', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    formData.value.conversationId = route.params.conversationId
    formData.value.threadId = route.params.threadId
    formData.value.cmtContent = text
    const res = await createWriteCmtTask(formData.value)
    if (res.code === 0) {
      formData.value = {
        conversationId: '',
        threadId: '',
        cmtContent: '',
      }
      ElMessage({
        type: 'success',
        message: '提交成功'
      })
      getTableData()
    }
  })

}

</script>

<style></style>
