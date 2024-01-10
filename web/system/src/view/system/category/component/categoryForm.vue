<template>
  <div>
    <el-dialog v-model="dialogVisible" :before-close="handleClose" :title="title">
      <el-form
        ref="form"
        :inline="true"
        :model="formData"
      >
        <el-form-item prop="name" label="名称：" style="width: 45%;" label-width="80px">
          <el-input v-model="formData.name" placeholder="请输入分类名称"/>
        </el-form-item>
        <el-form-item prop="parentId" label="父分类ID：" style="width: 35%;" label-width="100px">
          <el-input v-model="formData.parentId" disabled/>
        </el-form-item>
        <el-form-item prop="sorted" label="排序标记：" style="width: 35%;" label-width="100px">
          <el-input-number v-model="formData.sorted" :min="0" :max="100" class="mx-4" />
        </el-form-item>


        <el-form-item prop="parentId" label="类型：" style="width: 45%;" label-width="80px">
          <el-select v-model="formData.type"
                     disabled
                     placeholder="选择分类类型">
            <el-option
              v-for="item in typeOption"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item prop="status" label="状态：" style="width: 40%;" label-width="80px">
          <el-switch
            v-model="formData.status"
            active-text="激活"
            inactive-text="暂存"
            :active-value="1"
            :inactive-value="0"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { categoryCreateApi, categoryFindApi, categoryUpdateApi } from '@/api/system/category.js'
import { ElMessage } from 'element-plus'
import Form from '@/view/system/menu/component/form.vue'

// 驱动父类方法
const emit = defineEmits(['success'])

// 类型 option
const typeOption = ref([
  { label: '系统分类', value: 0 }
])
// 是否为编辑状态
const isEdit = ref(false)
// 弹出层 title
const title = ref('')
// 表单结构
const formData = ref({
  ID: 0,
  parentId: '',
  status: 0,
  type: 0,
  businessId: '',
  sorted: 0
})
// 弹出层开关
const dialogVisible = ref(false)


const initFormData = () => {
  formData.value = {
    ID: 0,
    categoryId: '',
    parentId: '',
    status: 0,
    type: 0,
    businessId: '',
    sorted: 0
  }
}
const openDialog = async(val) => {
  title.value = val
  if (isEdit.value) { // 编辑状态 数据回显
    const res = await categoryFindApi({categoryId: formData.value.categoryId})
    formData.value = res.data
  }
  dialogVisible.value = true
}
const closeDialog = () => {
  dialogVisible.value = false
  initFormData()
  isEdit.value = false
}

// 设置 form ID
const setFormParentId = (id) => {
  formData.value.parentId = id
}

// handleClose 默认关闭的回调
const handleClose = () => {
  dialogVisible.value = false
  isEdit.value = false
  initFormData()
}

// 弹出层提交
const enterDialog = async() => {
  let res
  formData.value.sorted = parseInt(formData.value.sorted)
  if (isEdit.value) { // 编辑菜单
    res = await categoryUpdateApi(formData.value)
  } else { // 添加菜单
    res = await categoryCreateApi(formData.value)
  }
  if (res.code === 0) {
    ElMessage({
      message: isEdit.value ? '更新分类成功' : '添加分类成功',
      type: 'success',
    })
    emit('success')
    closeDialog()
  }
}

// 设置 ID
const setFormCategoryId = (id) => {
  formData.value.categoryId = id
}

// 暴露给父组件的操作
defineExpose({
  openDialog,
  isEdit,
  setFormParentId,
  setFormCategoryId
})
</script>

<style scoped lang="scss">

</style>
