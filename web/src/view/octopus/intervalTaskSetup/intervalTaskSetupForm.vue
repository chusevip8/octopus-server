<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="应用名称:" prop="appName">
          <el-input v-model="formData.appName" :clearable="false"  placeholder="请输入应用名称" />
       </el-form-item>
        <el-form-item label="任务标题:" prop="taskTitle">
          <el-input v-model="formData.taskTitle" :clearable="false"  placeholder="请输入任务标题" />
       </el-form-item>
        <el-form-item label="脚本Id:" prop="scriptId">
          <el-input v-model.number="formData.scriptId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="脚本参数:" prop="params">
          <RichEdit v-model="formData.params"/>
       </el-form-item>
        <el-form-item label="间隔时间:" prop="intervalMin">
          <el-input v-model.number="formData.intervalMin" :clearable="false" placeholder="请输入" />
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
  createIntervalTaskSetup,
  updateIntervalTaskSetup,
  findIntervalTaskSetup
} from '@/api/octopus/intervalTaskSetup'

defineOptions({
    name: 'IntervalTaskSetupForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            appName: '',
            taskTitle: '',
            scriptId: undefined,
            params: '',
            intervalMin: undefined,
        })
// 验证规则
const rule = reactive({
               appName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               taskTitle : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               scriptId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               intervalMin : [{
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
      const res = await findIntervalTaskSetup({ ID: route.query.id })
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
               res = await createIntervalTaskSetup(formData.value)
               break
             case 'update':
               res = await updateIntervalTaskSetup(formData.value)
               break
             default:
               res = await createIntervalTaskSetup(formData.value)
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
