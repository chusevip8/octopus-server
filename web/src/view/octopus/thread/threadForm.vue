<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="会话ID:" prop="conversationID">
          <el-input v-model.number="formData.conversationID" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="消息发送者:" prop="sender">
          <el-input v-model="formData.sender" :clearable="false"  placeholder="请输入消息发送者" />
       </el-form-item>
        <el-form-item label="消息发送者ID:" prop="senderId">
          <el-input v-model="formData.senderId" :clearable="false"  placeholder="请输入消息发送者ID" />
       </el-form-item>
        <el-form-item label="消息接收者:" prop="receiver">
          <el-input v-model="formData.receiver" :clearable="false"  placeholder="请输入消息接收者" />
       </el-form-item>
        <el-form-item label="消息接收者ID:" prop="receiverId">
          <el-input v-model="formData.receiverId" :clearable="false"  placeholder="请输入消息接收者ID" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createThread,
  updateThread,
  findThread
} from '@/api/octopus/thread'

defineOptions({
    name: 'ThreadForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            conversationID: undefined,
            sender: '',
            senderId: '',
            receiver: '',
            receiverId: '',
        })
// 验证规则
const rule = reactive({
               conversationID : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sender : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               senderId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               receiver : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               receiverId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findThread({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createThread(formData.value)
               break
             case 'update':
               res = await updateThread(formData.value)
               break
             default:
               res = await createThread(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
