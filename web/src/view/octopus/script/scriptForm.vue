<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
       </el-form-item>
        <el-form-item label="内容:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item label="备注:" prop="note">
          <el-input v-model="formData.note" :clearable="true"  placeholder="请输入备注" />
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
  createScript,
  updateScript,
  findScript
} from '@/api/octopus/script'

defineOptions({
    name: 'ScriptForm'
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
            title: '',
            content: '',
            note: '',
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '脚本标题不能为空',
                   trigger: ['input','blur'],
               }],
               content : [{
                   required: true,
                   message: '脚本内容不能为空',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findScript({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rescript
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
               res = await createScript(formData.value)
               break
             case 'update':
               res = await updateScript(formData.value)
               break
             default:
               res = await createScript(formData.value)
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
