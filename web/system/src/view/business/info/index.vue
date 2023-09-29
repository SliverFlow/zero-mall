<template>
  <div class="z-container">
    <div class="left">
      <el-image
        class="image"
        :src="formData.image[0]"
        :zoom-rate="1.2"
        close-on-press-escape
        preview-teleported
        lazy
        preview-src-list=""
        :initial-index="4"
        fit="cover"
      />
      <div class="title">
        <span v-if="!isEdit" class="name">{{ formData.name }}</span>
        <el-icon v-if="!isEdit" @click="isEdit = !isEdit" style="position: relative;top: 3px;cursor: pointer"
                 color="#4d70ff" size="18">
          <edit/>
        </el-icon>
        <div v-if="isEdit" class="in">
          <el-input v-model="formData.name" style="margin-right: 2px"/>
          <el-icon @click="enterBusinessName" color="green" class="icon">
            <Check/>
          </el-icon>
          <el-icon @click="closeBusinessName" color="red" class="icon">
            <Close/>
          </el-icon>
        </div>
      </div>
      <div class="time">
        注册时间：
        <el-text type="success">{{ formatTimestamp(formData.createdAt) }}</el-text>
      </div>
    </div>
    <div class="right">
      <el-tabs v-model="activeName" class="demo-tabs">
        <el-tab-pane label="相关信息" name="business">
          <el-form-item label="介绍：" style="margin-top: 16px">
            <el-input v-model="formData.detail" :disabled="!isEditDetail" :autosize="{ minRows: 2, maxRows: 10 }"
                      type="textarea"/>
          </el-form-item>
          <el-form-item label="评分：">
            <el-rate
              v-model="formData.score"
              disabled
              show-score
              score-template=""
            />
          </el-form-item>
          <el-form-item label="状态：">
            <el-text v-if="formData.status === 0" type="info">暂存</el-text>
            <el-text v-if="formData.status === 1" type="success">正常</el-text>
            <el-text v-if="formData.status === 2" type="danger">禁用</el-text>
          </el-form-item>
          <div style="color: #606266">
            信息更新事件：
            <el-text type="success">{{ formatTimestamp(formData.createdAt) }}</el-text>
          </div>
          <div class="btn-list">
            <el-button @click="isEditDetail = true" v-if="!isEditDetail" type="primary" icon="edit">点击修改详情
            </el-button>
            <el-button @click="enterEditDetail"  v-if="isEditDetail" type="primary" icon="check">确认</el-button>
            <el-button @click="closeEditDetail" v-if="isEditDetail" icon="close">取消</el-button>
          </div>
        </el-tab-pane>
      </el-tabs>

    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { formatTimestamp } from '@/utils/date.js'
import { businessFindApi, businessUpdateApi } from '@/api/system/business.js'
import { ElMessage } from 'element-plus'

const activeName = ref('business')
// 是否编辑名称
const isEdit = ref(false)
// 是否编辑详情
const isEditDetail = ref(false)
// 表单信息
const formData = ref({
  businessName: '',
  createdAt: 0,
  detail: '',
  status: 0,
  score: 3,
  businessId: '',
  updatedAt: 0,
  image: [],
  name: ''
})

// 加载商户数据
const loadFormData = async() => {
  const res = await businessFindApi()
  if (res.code === 0) {
    console.log(res)
    formData.value = res.data
  }
}
loadFormData()
// 提交修改的商户信息
const enterBusinessName = async() => {
  const res = await businessUpdateApi(formData.value)
  if (res.code === 0) {
    ElMessage({
      message: '更新商户信息成功',
      type: 'success',
    })
    await loadFormData()
    isEdit.value = false
  }
}
// 关闭修改商户信息
const closeBusinessName = async () => {
  // 数据回调
  await loadFormData()
  isEdit.value = false
}

// 关闭修改相亲
const closeEditDetail = () => {
  isEditDetail.value = false
}
// 提交修改详情
const enterEditDetail = async () => {
  const res = await businessUpdateApi(formData.value)
  if (res.code === 0) {
    ElMessage({
      message: '更新商户信息成功',
      type: 'success',
    })
    await loadFormData()
    isEditDetail.value = false
  }
}
</script>

<style scoped lang="scss">
.z-container {
  width: 100%;
  display: flex;

  .left {
    display: flex;
    align-items: center;
    flex-direction: column;
    background-color: white;
    padding: 20px;
    box-shadow: 0 0 37px -7px rgba(0, 0, 0, 0.18);
    border-radius: 4px;
    margin-right: 18px;

    .image {
      margin-top: 20px;
      width: 260px;
      height: 260px;
      border-radius: 4px;
    }

    .title {
      height: 24px;
      margin-top: 24px;
      display: flex;
      align-items: center;

      .name {
        font-size: 20px;
        margin-right: 10px;
        color: #374151;
      }

      .in {
        display: flex;
        align-items: center;

        .icon {
          margin-left: 10px;
          cursor: pointer;
        }

      }
    }

    .time {
      margin-top: 22px;
      color: #606266;
    }
  }

  .right {
    width: calc(100% - 328px);
    background-color: white;
    padding: 40px 20px;
    box-shadow: 0 0 37px -7px rgba(0, 0, 0, 0.18);
    border-radius: 4px;

    .btn-list {
      margin-top: 26px;
    }
  }
}

:deep(.el-textarea__inner) {
  resize: none;
}
</style>
