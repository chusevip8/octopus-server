<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="应用名称:" prop="appName">
          <el-input v-model="formData.appName" :clearable="false"  placeholder="请输入应用名称" />
       </el-form-item>
        <el-form-item label="帖子链接:" prop="postLink">
          <el-input v-model="formData.postLink" :clearable="false"  placeholder="请输入帖子链接" />
       </el-form-item>
        <el-form-item label="评论关键字:" prop="keyword">
          <el-input v-model="formData.keyword" :clearable="false"  placeholder="请输入评论关键字" />
       </el-form-item>
        <el-form-item label="查找评论脚本Id:" prop="findCommentScriptId">
          <el-input v-model.number="formData.findCommentScriptId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="第一次回复评论脚本Id:" prop="writeCommentScriptId">
          <el-input v-model.number="formData.writeCommentScriptId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="回复评论脚本id:" prop="replyCommentScriptId">
          <el-input v-model.number="formData.replyCommentScriptId" :clearable="false" placeholder="请输入" />
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
  createCmtTaskSetup,
  updateCmtTaskSetup,
  findCmtTaskSetup
} from '@/api/octopus/cmtTaskSetup'

defineOptions({
    name: 'CmtTaskSetupForm'
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
            postLink: '',
            keyword: '',
            findCommentScriptId: undefined,
            writeCommentScriptId: undefined,
            replyCommentScriptId: undefined,
        })
// 验证规则
const rule = reactive({
               appName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               postLink : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               findCommentScriptId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               writeCommentScriptId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               replyCommentScriptId : [{
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
      const res = await findCmtTaskSetup({ ID: route.query.id })
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
