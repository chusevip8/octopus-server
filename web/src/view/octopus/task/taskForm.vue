<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="App名称:" prop="appName">
          <el-input v-model="formData.appName" :clearable="false"  placeholder="请输入App名称" />
       </el-form-item>
        <el-form-item label="任务类型:" prop="type">
          <el-input v-model.number="formData.type" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务参数Id:" prop="taskOptionId">
          <el-input v-model.number="formData.taskOptionId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备Id:" prop="deviceId">
          <el-input v-model.number="formData.deviceId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="任务错误信息:" prop="error">
          <el-input v-model="formData.error" :clearable="false"  placeholder="请输入任务错误信息" />
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
  createTask,
  updateTask,
  findTask
} from '@/api/octopus/task'

defineOptions({
    name: 'TaskForm'
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
            appName: '',
            type: undefined,
            taskOptionId: undefined,
            deviceId: undefined,
            status: undefined,
            error: '',
        })
// 验证规则
const rule = reactive({
               appName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               taskOptionId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               deviceId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
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
      const res = await findTask({ ID: route.query.id })
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
               res = await createTask(formData.value)
               break
             case 'update':
               res = await updateTask(formData.value)
               break
             default:
               res = await createTask(formData.value)
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
