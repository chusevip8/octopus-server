<template>
  <div>
    <JwChat :taleList="tableData" @enter="onSubmit" v-model="formData.msgContent" :showRightBox="false"
      scrollType="scroll" width="80%" height="800px" :toolConfig="chatTool" :config="chatConfig">
    </JwChat>
  </div>
</template>

<script setup>
import {
  getMessageList
} from '@/api/octopus/message'

import {
  createReplyMsgTask
} from '@/api/octopus/msgTask'

// 全量引入格式化工具 请按需保留
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'Message'
})

const props = defineProps({

  conversationId: {
    type: String,
    required: true
  }
})


const route = useRoute()

const chatTool = ref({
  show: []
})

const chatConfig = ref({
  name: '发信息',
  dept: ''
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  conversationId: '',
  msgContent: '',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(50)
const tableData = ref([])

// 查询
const getTableData = async () => {
  const table = await getMessageList({ page: page.value, pageSize: pageSize.value, conversationId: props.conversationId })
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
  if (formData.value.msgContent.trim() === '') {
    return
  }
  const text = formData.value.msgContent
  ElMessageBox.confirm('确定要提交吗?提交后消息不能修改和删除。', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    formData.value.conversationId = props.conversationId + ''
    formData.value.msgContent = text
    const res = await createReplyMsgTask(formData.value)
    if (res.code === 0) {
      formData.value = {
        conversationId: '',
        msgContent: '',
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
