<template>
  <div class="z-container">
    <div class="left">
      <el-image
        class="image"
        :src="formData.avatar"
        :zoom-rate="1.2"
        close-on-press-escape
        preview-teleported
        lazy
        :initial-index="4"
        fit="cover"
      />
      <div class="title">
        <span v-if="!isEdit" class="name">{{ formData.nickname }}</span>
        <el-icon v-if="!isEdit" @click="isEdit = !isEdit" style="position: relative;top: 3px;cursor: pointer"
                 color="#4d70ff" size="18">
          <edit/>
        </el-icon>
        <div v-if="isEdit" class="in">
          <el-input v-model="formData.nickname" style="margin-right: 2px"/>
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
          <el-form-item label="状态：">
            <el-text v-if="formData.status === 2" type="danger">禁用</el-text>
            <el-switch
              v-else
              v-model="formData.status"
              :disabled="true"
              inline-prompt
              :active-value="1"
              :inactive-value="0"
              active-text="正常"
              inactive-text="暂存"
            />
          </el-form-item>
          <el-form-item label="邮箱：" style="margin-top: 16px">
            <el-input v-model="formData.email"/>
          </el-form-item>
          <el-form-item label="原密码：" style="margin-top: 16px">
            <el-input v-model="formData.password" type="password"/>
          </el-form-item>
          <el-form-item label="新密码：" style="margin-top: 16px">
            <el-input v-model="formData.newPassword" type="password"/>
          </el-form-item>
          <el-form-item label="电话号码：" style="margin-top: 16px">
            <el-input v-model="formData.phone"/>
          </el-form-item>
          <div style="color: #606266">
            信息更新时间：
            <el-text type="success">{{ formatTimestamp(formData.updatedAt) }}</el-text>
          </div>
          <div class="btn-list">
            <el-button type="primary" icon="edit" @click="submitForm">保存</el-button>
          </div>
        </el-tab-pane>
      </el-tabs>

    </div>
  </div>
</template>

<script setup>
import { ref, watch, watchEffect } from 'vue'
import { formatTimestamp } from '@/utils/date.js'
import { businessFindApi, businessUpdateApi } from '@/api/system/business.js'
import { ElMessage } from 'element-plus'
import { Watch } from '@element-plus/icons-vue'
import { userFindByUUIDApi, userUpdateByUUIDApi } from '@/api/system/user.js'

const activeName = ref('business')
// 是否编辑名称
const isEdit = ref(false)
// 表单信息
const formData = ref({
  nickname: '',
  createdAt: 0,
  email: '',
  status: 0,
  updatedAt: 0,
  avatar: '',
  phone: '',
  password: '',
  newPassword: ''
})

const loadFormData = async() => {
  const res = await userFindByUUIDApi()
  formData.value = res.data.user
}
loadFormData()
// 提交修改的商户信息
const enterBusinessName = async() => {
  const res = await businessUpdateApi(formData.value)
  if (res.code !== 0) return
  ElMessage({
    message: '更新商户信息成功',
    type: 'success',
  })
  await loadFormData()
  isEdit.value = false
}
// 关闭修改商户信息
const closeBusinessName = async() => {
  // 数据回调
  await loadFormData()
  isEdit.value = false
}

// 提交修改详情
const submitForm = async() => {
  formData.value.password = ''
  const res = await userUpdateByUUIDApi(formData.value)
  if (res.code !== 0) return
  ElMessage({
    message: '更新用户信息成功',
    type: 'success',
  })
  await loadFormData()
}

// 监听取消修改状态
watch(isEdit, async(val, old) => {
  if (!val) {
    await loadFormData()
  }
})
</script>

<style scoped lang="scss">
@import "@/style/info";
</style>
