<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="App名称:" prop="appName">
          <el-input v-model="formData.appName" :clearable="false"  placeholder="请输入App名称" />
       </el-form-item>
        <el-form-item label="任务标题:" prop="title">
          <el-input v-model="formData.title" :clearable="false"  placeholder="请输入任务标题" />
       </el-form-item>
        <el-form-item label="文章ID:" prop="articleID">
          <el-input v-model="formData.articleID" :clearable="false"  placeholder="请输入文章ID" />
       </el-form-item>
        <el-form-item label="评论关键字:" prop="keyword">
          <el-input v-model="formData.keyword" :clearable="false"  placeholder="请输入评论关键字" />
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
  createCommentTask,
  updateCommentTask,
  findCommentTask
} from '@/api/octopus/commentTask'

defineOptions({
    name: 'CommentTaskForm'
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
            title: '',
            articleID: '',
            keyword: '',
        })
// 验证规则
const rule = reactive({
               appName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '任务标题不能为空',
                   trigger: ['input','blur'],
               }],
               articleID : [{
                   required: true,
                   message: '文章ID不能为空',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCommentTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recommentTask
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
               res = await createCommentTask(formData.value)
               break
             case 'update':
               res = await updateCommentTask(formData.value)
               break
             default:
               res = await createCommentTask(formData.value)
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
