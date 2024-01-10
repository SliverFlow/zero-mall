<template>
  <el-upload
    v-model:file-list="fileList"
    list-type="picture-card"
    :http-request="uploadFile"
    :on-success="uploadSuccess"
    :show-file-list="false"
  >
    <el-icon><Plus /></el-icon>
  </el-upload>
</template>

<script setup>
import { ref } from 'vue'
import { getOssClient } from '@/utils/oss/oss.js'
import { ElMessage } from 'element-plus'
import { getUUID } from '@/utils/uuid.js'

const fileList = ref([])

const clientParams = ref({
  accessKeyId: "STS.NTxPryVAvBzhkzKRZzGARjGtk",
  accessKeySecret: "72aQa5VTb1ZgNFWgoMV3upZ8vXuGhFDLE2LngA7o2Bkv",
  stsToken: "CAIS/gF1q6Ft5B2yfSjIr5fNG8jNu55X9biDaVz6tlovS85+hYLfiTz2IHhMe3BvB+0esvw1n2hY6vwYlqRoRoReREvCWpIotswOqZgX+W8E4p7b16cNrbH4M0rxYkeJ8a2/SuH9S8ynCZXJQlvYlyh17KLnfDG5JTKMOoGIjpgVBbZ+HHPPD1x8CcxROxFppeIDKHLVLozNCBPxhXfKB0ca0WgVy0EHsPrmk5bEs0WO1gallrNN9r6ceMb0M5NeW75kSMqw0eBMca7M7TVd8RAi9t0t1PQboGiZ44nBWwMKu0XaabDOiM9iNgg8Y6E+H+tbofK5j/p8t/wyb3sDUunVVxqAARTADOydoUfW1cUTaJ7s3a3PZ4rghoU+2V2sY3REgUj7LV1C4b/sLVy1cnWheZXNojoC1RlP4gysES+W8BAE7uxOGB8CyMsx5YLmsqTzkqTlVhy0TQ3f46lRXpdL5u1Tb3FOmP3E3DUsXBQcXeSCBSWEfbNu5ftYVpqBjwMl4/eOIAA=",
  region: "oss-cn-hangzhou",
  bucket: "zero-mall"
})

// 文件上传的回调
const uploadFile = async (file) => {
  let client = getOssClient(clientParams.value)
  const fileName = getUUID() + file.file.name.substring(file.file.name.lastIndexOf('.')).toLowerCase()
  try {
    const res = await client.put("/web_files/test/" + fileName, file.file)
    ElMessage({
      message: '上传文件成功',
      type: 'success',
    })
    fileList.value.push(res.url)
  }catch (err) {
    ElMessage({
      message: '上传文件失败',
      type: 'error',
    })
  }
}
const uploadSuccess = (response,file,fileList) => {

}
</script>

<style scoped lang="scss">

</style>
