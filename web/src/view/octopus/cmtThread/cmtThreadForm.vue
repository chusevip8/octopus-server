<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="应用名称:" prop="appName">
          <el-input v-model="formData.appName" :clearable="false"  placeholder="请输入应用名称" />
       </el-form-item>
        <el-form-item label="任务参数Id:" prop="taskOptionId">
          <el-input v-model.number="formData.taskOptionId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="帖子Id:" prop="postId">
          <el-input v-model="formData.postId" :clearable="false"  placeholder="请输入帖子Id" />
       </el-form-item>
        <el-form-item label="发帖者:" prop="poster">
          <el-input v-model="formData.poster" :clearable="false"  placeholder="请输入发帖者" />
       </el-form-item>
        <el-form-item label="帖子标题:" prop="postTitle">
          <el-input v-model="formData.postTitle" :clearable="false"  placeholder="请输入帖子标题" />
       </el-form-item>
        <el-form-item label="帖子描述:" prop="postDesc">
          <el-input v-model="formData.postDesc" :clearable="false"  placeholder="请输入帖子描述" />
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
  createCmtThread,
  updateCmtThread,
  findCmtThread
} from '@/api/octopus/cmtThread'

defineOptions({
    name: 'CmtThreadForm'
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
            taskOptionId: undefined,
            postId: '',
            poster: '',
            postTitle: '',
            postDesc: '',
        })
// 验证规则
const rule = reactive({
               appName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               taskOptionId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               postId : [{
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
      const res = await findCmtThread({ ID: route.query.id })
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
               res = await createCmtThread(formData.value)
               break
             case 'update':
               res = await updateCmtThread(formData.value)
               break
             default:
               res = await createCmtThread(formData.value)
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
