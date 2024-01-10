<template>
  <div>
    <el-dialog v-model="dialogVisible" :before-close="handleClose" :title="title">
      <el-form
        ref="form"
        :inline="true"
        :model="formData"
      >
        <el-form-item prop="name" label="名称：" style="width: 45%;" label-width="100px">
          <el-input v-model="formData.name" placeholder="请输入商品名称"/>
        </el-form-item>
        <el-form-item prop="category" label="分类：" style="width: 45%;" label-width="100px">
          <el-cascader
            v-model="formData.category"
            :options="categoryOptions"
            :props="{label:'name',value: 'ID',expandTrigger: 'hover'}"
            @change="handleCascaderChange"
          />
        </el-form-item>
        <el-form-item prop="name" label="图片：" label-width="100px">
          <el-upload
            list-type="picture-card"
            v-model:file-list="formData.image"
            :before-remove="beforeRemove"
            :before-upload="beforeUpload"
            :on-success="onSuccess"
            :http-request="uploadFile"
          >
            <el-icon>
              <Plus/>
            </el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item prop="subtitle" label="副标题：" style="width: 100%" label-width="100px">
          <el-input v-model="formData.subtitle"/>
        </el-form-item>
        <el-form-item prop="detail" label="详细介绍：" style="width: 100%" label-width="100px">
          <el-input v-model="formData.detail" :autosize="{ minRows: 2, maxRows: 10 }"
                    type="textarea"/>
        </el-form-item>
        <el-form-item prop="sorted" label="库存：" style="width: 35%;" label-width="100px">
          <el-input-number v-model="formData.stock" :min="1" :max="1000" class="mx-4"/>
        </el-form-item>
        <el-form-item prop="price" label="价格：" style="width: 50%;" label-width="100px">
          <el-input-number v-model="formData.price" :min="1" :max="1000000" class="mx-4"/>
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
import { ref, watch, watchEffect } from 'vue'
import { getOssClient } from '@/utils/oss/oss.js'
import { getUUID } from '@/utils/uuid.js'
import { ElMessage } from 'element-plus'
import { productFileTokenApi } from '@/api/system/file.js'
import { categoryTressListApi } from '@/api/system/category.js'
import { productCreateApi, productFindApi, productUpdateApi } from '@/api/system/product.js'
import Form from '@/view/system/menu/component/form.vue'

// 驱动父类方法
const emit = defineEmits(['success'])
// 是否为编辑状态
const isEdit = ref(false)
// 弹出层 title
const title = ref('')
// 表单结构
const formData = ref({
  ID: 0,
  productId: '',
  name: '',
  subtitle: '',
  image: [],
  detail: '',
  price: 0,
  stock: 0,
  status: 0,
  categories: [],
  category: []
})
// 分类级联选择器
const categoryOptions = ref([])
// 弹出层开关
const dialogVisible = ref(false)

// 加载
const loadCategoryOptions = async() => {
  // 获取分类列表
  const categoryReq = await categoryTressListApi()
  categoryOptions.value = categoryReq.data.list
}
loadCategoryOptions()
// 初始化表单值
const initFormData = () => {
  formData.value = {
    ID: 0,
    productId: '',
    name: '',
    subtitle: '',
    image: [],
    detail: '',
    price: 0,
    stock: 0,
    status: 0,
    categories: [],
    category: [],
  }
}
// 级联选择器发生改变的时候
const handleCascaderChange = (val) => {
  categoryOptions.value.forEach(i => {
    if (i.children.length && i.children.length > 0) {
      i.children.forEach(v => {
        if (v.ID === val[1]) {
          formData.value.categories = []
          formData.value.categories.push(v)
          formData.value.categories.push(i)
        }
      })
    }
  })
}
// 弹出层关闭
const handleClose = () => {
  dialogVisible.value = false
  initFormData()
  isEdit.value = false
  title.value = ''
}

// 打开
const openDialog = async(val) => {
  if (isEdit.value) {
    await loadCategoryOptions()
    // 数据回显
    const res = await productFindApi({ productId: formData.value.productId })
    formData.value = res.data
    formData.value.category = []
    // formData.value.price = formData.value.price.toFixed(2)

    if (res.data.categories && res.data.categories.length > 0) {
      res.data.categories.forEach(i => {
        // categoryOptions.value.forEach(v => {
        //   if (v.categoryId === i.parentId) {
        //     formData.value.category.push(v.ID)
        //   }
        // })
        formData.value.category.push(i.ID)
      })
    }
  }
  title.value = val
  dialogVisible.value = true
}
//  关闭
const closeDialog = () => {
  handleClose()
}
// 提交表单
const enterDialog = async() => {
  let res
  if (isEdit.value) {
    // 商品更新
    res = await productUpdateApi(formData.value)
  } else {
    //  商品新增
    res = await productCreateApi(formData.value)
  }
  if (res.code === 0) {
    ElMessage({
      message: isEdit.value ? '更新商品信息成功' : '添加商品成功',
      type: 'success',
    })
    emit('success')
    closeDialog()
  }
}
// 设置商品 id
const setProductID = (val) => {
  formData.value.productId = val
}

// 文件上传相关
const clientParams = ref({
  accessKeyId: '',
  accessKeySecret: '',
  stsToken: '',
  region: '',
  bucket: '',
  filePath: '',
})
const stsClient = ref(null)
// 文件上传的回调
const uploadFile = async(file) => {
  const fileName = getUUID() + file.file.name.substring(file.file.name.lastIndexOf('.')).toLowerCase()
  try {
    const res = await stsClient.value.put(clientParams.value.filePath + fileName, file.file)
    ElMessage({
      message: '上传文件成功',
      type: 'success',
    })
    formData.value.image.push({ url: res.url })
  } catch (err) {
    ElMessage({
      message: '上传文件失败',
      type: 'error',
    })
  }
}
const beforeRemove = (file, files) => {
  formData.value.image.forEach(i => {
    if (i.url === file.url) {
      formData.value.image.splice(formData.value.image.indexOf(i), 1)
    }
  })
  return false
}
const beforeUpload = (file) => {
  window.URL.revokeObjectURL(file.raw)
}
const onSuccess = (res, file, files) => {

}

watch(() => formData.value.image, () => {
  formData.value.image.forEach(i => {
    if (i.raw) {
      formData.value.image.splice(formData.value.image.indexOf(i), 1)
    }
  })
}, {
  immediate: true,
  deep: true
})

watchEffect(async() => {
  // 打开弹出层的时候初始化
  if (dialogVisible.value === true) {
    const res = await productFileTokenApi()
    if (res.code === 0) {
      clientParams.value = res.data
      stsClient.value = getOssClient(clientParams.value)
    }
  }
})

// 暴露给父组件的操作
defineExpose({
  openDialog,
  isEdit,
  setProductID
})
</script>

<style scoped lang="scss">

:deep(.el-textarea__inner) {
  resize: none;
}
</style>
