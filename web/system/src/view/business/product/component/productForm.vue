<template>
  <div>
    <el-dialog v-model="dialogVisible" :before-close="handleClose" :title="title">
      <el-form
        ref="form"
        :inline="true"
        :model="formData"
      >
        <el-form-item prop="name" label="名称：" style="width: 45%;" label-width="80px">
          <el-input v-model="formData.name" placeholder="请输入商品名称"/>
        </el-form-item>
        <el-form-item prop="name" label="分类：" style="width: 45%;" label-width="80px">
          <el-input v-model="formData.categorys" placeholder="请输入商品名称"/>
        </el-form-item>
        <el-form-item prop="name" label="图片：" style="width: 100%;" label-width="80px">
          <template v-for="(v,k) in formData.image">
            <el-image
              style="width: 100px; height: 100px;z-index: 100;margin-right: 10px"
              :src="v"
              :zoom-rate="1.2"
              close-on-press-escape
              preview-teleported
              lazy
              fit="cover"
            />
          </template>
          <el-upload
            style="height: 100px;width: 100px"
            list-type="picture-card"
            :http-request="uploadFile"
            :on-success="uploadSuccess"
            :show-file-list="false"
          >
            <el-icon>
              <Plus/>
            </el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item prop="parentId" label="父分类ID：" style="width: 35%;" label-width="100px">
          <el-input v-model="formData.parentId" disabled/>
        </el-form-item>
        <el-form-item prop="sorted" label="排序标记：" style="width: 35%;" label-width="100px">
          <el-input-number v-model="formData.sorted" :min="0" :max="100" class="mx-4"/>
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
import { ref, watch, watchEffect } from 'vue'
import { getOssClient } from '@/utils/oss/oss.js'
import { getUUID } from '@/utils/uuid.js'
import { ElMessage } from 'element-plus'
import { productFileTokenApi } from '@/api/system/file.js'

// 驱动父类方法
const emit = defineEmits(['success'])
// 是否为编辑状态
const isEdit = ref(false)
// 弹出层 title
const title = ref('')
// 表单结构
const formData = ref({
  name: '',
  subtitle: '',
  image: [],
  detail: '',
  price: 0,
  stock: 0,
  status: 0,
  categories: []
})
// 弹出层开关
const dialogVisible = ref(false)
// 打开
const openDialog = (val) => {
  dialogVisible.value = true
  title.value = val
}
//  关闭
const closeDialog = () => {
  dialogVisible.value = false
  title.value = ''
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
    formData.value.image.push(res.url)
  } catch (err) {
    ElMessage({
      message: '上传文件失败',
      type: 'error',
    })
  }
}
const uploadSuccess = (response, file, fileList) => {
  console.log(response, file, fileList)
}

watchEffect( async () => {
  if (dialogVisible.value === true) {
    const res = await productFileTokenApi()
    console.log(res)
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
})
</script>

<style scoped lang="scss">
:deep(.el-upload--picture-card) {
  width: 100% !important;
  height: 100% !important;
}
</style>
