<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="视频标题:" prop="videoTitle">
          <el-input v-model="formData.videoTitle" :clearable="false"  placeholder="请输入视频标题" />
       </el-form-item>
        <el-form-item label="视频ID:" prop="videoId">
          <el-input v-model="formData.videoId" :clearable="false"  placeholder="请输入视频ID" />
       </el-form-item>
        <el-form-item label="关键字:" prop="keyword">
          <el-input v-model="formData.keyword" :clearable="false"  placeholder="请输入关键字" />
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
  createDYCmtTask,
  updateDYCmtTask,
  findDYCmtTask
} from '@/api/octopus/dyCmtTask'

defineOptions({
    name: 'DYCmtTaskForm'
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
            videoTitle: '',
            videoId: '',
            keyword: '',
        })
// 验证规则
const rule = reactive({
               videoTitle : [{
                   required: true,
                   message: '视频标题不能为空',
                   trigger: ['input','blur'],
               }],
               videoId : [{
                   required: true,
                   message: '视频ID不能为空',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDYCmtTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redyCmtTask
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
               res = await createDYCmtTask(formData.value)
               break
             case 'update':
               res = await updateDYCmtTask(formData.value)
               break
             default:
               res = await createDYCmtTask(formData.value)
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
