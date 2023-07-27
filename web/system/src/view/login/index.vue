<template>
  <div id="userLayout">
    <div class="login_panel">
      <div class="login_panel_form">
        <div class="login_panel_form_title">
          <!--          <img-->
          <!--            class="login_panel_form_title_logo"-->
          <!--            src=""-->
          <!--            alt-->
          <!--          >-->
          <p class="login_panel_form_title_p">ZPSHOP 系统后台</p>
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
                    <user />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="formData.password"
              placeholder="请输入密码"
              type="password"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <lock />
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
                  alt="请输入验证码"
                  @click="loginVerify()"
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
      <div class="login_panel_right" />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { captchaApi } from '@/api/base.js'

// form ref
const form = ref(null)
// form 表单数据
const formData = ref({
  captchaId: '',
  username: '',
  password: '',
  captcha: ''
})
// 验证码图片地址
const picPath = ref('')

// 获取验证码
const getCaptcha = async() => {
  const res = await captchaApi()
  if (res.code === 0) {
    picPath.value = res.data.picPath
    formData.value.captchaId = res.data.captchaId
  }
}
getCaptcha()
</script>

<style lang="scss" scoped>
@import "@/style/newLogin";
</style>
