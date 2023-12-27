<template>
  <div id="userLayout">
    <div class="login_panel">
      <div class="login_panel_form">
        <div class="login_panel_form_title"
             style="display: flex;align-items: center;align-content: center;flex-direction: column">
          <!--          <img-->
          <!--            class="login_panel_form_title_logo"-->
          <!--            src=""-->
          <!--            alt-->
          <!--          >-->
          <p class="login_panel_form_title_p" style="font-family: 'fm'">zero-mall&nbsp;系统后台</p>
        </div>
        <el-form ref="form" :model="formData">
          <el-form-item prop="username">
            <el-input
              v-model="formData.username"
              placeholder="请输入用户名"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <user/>
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="formData.password"
              placeholder="请输入密码"
              :type="isLock ? 'password' : 'text'"
            >
              <template #suffix>
                <span class="input-icon" @click="isLock = !isLock">
                  <el-icon>
                    <lock v-if="isLock"/>
                    <unlock v-if="!isLock"/>
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="captcha">
            <div class="vPicBox">
              <el-input
                v-model="formData.captcha"
                placeholder="请输入验证码"
                style="width: 60%"
              />
              <div class="vPic" style="margin-top: 2px">
                <img
                  v-if="picPath"
                  :src="picPath"
                  alt="验证码"
                  @click="getCaptcha"
                >
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              size="large"
              style="width: 100%;"
              @click="submitForm"
            >登 录
            </el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="login_panel_right"/>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { captchaApi } from '@/api/system/base.js'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/store/model/user.js'

// 获取路由信息
const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
// 获取路径参数
const path = route.query.path || '/'
// form ref
const form = ref(null)
// form 表单数据
const formData = ref({
  captchaId: '',
  username: 'zhengjunxing',
  password: '123456',
  captcha: ''
})
// 验证码图片地址
const picPath = ref('')
// 锁图标
const isLock = ref(true)
// 获取验证码
const getCaptcha = async() => {
  const res = await captchaApi()
  if (res.code === 0) {
    picPath.value = res.data.picPath
    formData.value.captchaId = res.data.captchaId
  }
}
getCaptcha()

// 提交表单
const submitForm = async() => {
  const flag = await userStore.login(formData.value)
  if (flag) {
    await router.push({ path: path })
  } else {
    // 重新获取验证码
    await getCaptcha()
  }
}
</script>

<style lang="scss" scoped>
@import "@/style/newLogin";
</style>
