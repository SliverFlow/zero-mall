<template>
  <div>
    <div class="login_panel">
      <div class="input_item">
        <input
          v-model="formData.username"
          :class="{isFocus: usernameFocus || usernameInfo}"
          type="text"
          @focus="usernameFocus = true"
          @focusout="usernameFocusOut"
        >
        <span
          v-if="loginType === 'password'"
          :class="{isActive: usernameFocus || usernameInfo}"
        >手机号码 / 用户登录名</span>
        <span v-if="loginType === 'captcha'" :class="{isActive: usernameFocus || usernameInfo}">手机号码</span>
        <span v-if="usernameInfo && loginType === 'password'" class="info">请输入账号</span>
        <span v-if="usernameInfo && loginType === 'captcha'" class="info">请输入手机号码</span>
      </div>
      <div class="input_item">
        <input
          v-if="loginType === 'password'"
          v-model="formData.password"
          :class="{isFocus: passwordFocus || passwordInfo}"
          :type="!isLock ? 'password' : 'text'"
          @focus="passwordFocus = true"
          @focusout="passwordFocusOut"
        >
        <div v-if="loginType === 'captcha'" style="display: flex;position: relative">
          <input
            v-model="formData.password"
            :class="{isFocus: passwordFocus || passwordInfo}"
            type="text"
            @focus="passwordFocus = true"
            @focusout="passwordFocusOut"
          >
          <a class="captcha" @click.prevent="sendCaptchaMessage">{{ message }}</a>
        </div>
        <span v-if="loginType === 'password'" :class="{isActive: passwordFocus || passwordInfo}">密码</span>
        <span v-if="loginType === 'captcha'" :class="{isActive: passwordFocus || passwordInfo}">验证码</span>
        <div v-if="loginType === 'password'" class="input_icon" @click="passwordFocus = true;isLock = !isLock;">
          <el-icon size="22" color="#959595">
            <Unlock v-if="isLock"/>
            <lock v-if="!isLock"/>
          </el-icon>
        </div>
        <span v-if="passwordInfo && loginType ==='password'" class="info">请输入密码</span>
        <span v-if="passwordInfo && loginType ==='captcha'" class="info">请输入验证码</span>
      </div>
      <div class="info">
        <el-checkbox v-model="isActiveInfo" size="large"/>
        <span>已阅读并同意 zero-mall 用户协议 和 隐私政策</span>
      </div>
      <button :class="{'is-disabled': !isDisabled}" :disabled="!isDisabled" @click="enterLogin">登录</button>
      <div class="phone">
        <span>忘记密码？</span>
        <span v-if="loginType === 'password'" @click="initForm(); loginType = 'captcha'">手机验证码登录</span>
        <span v-if="loginType === 'captcha'" @click="initForm(); loginType = 'password'">密码登录</span>
      </div>
      <div class="other">
        <span>其他方式登录</span>
        <div class="info_img">
          <img src="@/assets/zfb.svg" alt="">
          <img src="@/assets/wx.svg" alt="">
          <img src="@/assets/qq.svg" alt="">
          <img src="@/assets/wb.svg" alt="">
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watchEffect } from 'vue'
import { ElMessage } from 'element-plus'
import { phoneCaptchaApi, userLoginByPhoneApi, userLoginByPhoneOrUsernameApi } from '@/api/user.js'
import { useUserStore } from '@/pinia/model/user.js'
import { useRoute, useRouter } from 'vue-router'

const passwordFocus = ref()
const passwordInfo = ref(false)
const usernameFocus = ref()
const usernameInfo = ref(false)
const isLock = ref(false)
const isActiveInfo = ref(false)
const loginType = ref('password')

const formData = ref({
  password: '',
  username: '',

})

const passwordFocusOut = () => {
  if (formData.value.password === '') {
    passwordFocus.value = false
    passwordInfo.value = true
  } else {
    passwordFocus.value = true
    passwordInfo.value = false
  }
}
const usernameFocusOut = () => {
  if (formData.value.username === '') {
    usernameFocus.value = false
    usernameInfo.value = true
  } else {
    usernameFocus.value = true
    usernameInfo.value = false
  }
}

// -----------------------------------------------------------------
// 显示的消息
const message = ref('获取验证码')
// 倒计时
const count = ref(60)
// 定时器
const timer = ref(null)
// 是否可以登录
const isDisabled = ref(false)
// 用户信息存储
const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
// 记录路径
const path = route.query.path || '/'

// 发送验证码消息
const sendCaptchaMessage = async() => {
  if (!/^1[3456789]\d{9}$/.test(formData.value.username)) {
    ElMessage({
      message: '请按照要求填写信息',
      type: 'warning',
      showClose: true,
    })
    return
  }
  // 发送验证码
  const res = await phoneCaptchaApi({ phone: formData.value.username })
  if (res.code === 0) {
    ElMessage({
      message: '验证码发送成功',
      type: 'success',
      showClose: true,
    })
    message.value = `${ count.value }秒后重新发送`
    if (timer.value) clearInterval(timer.value)
    timer.value = setInterval(() => {
      if (count.value <= 1) {
        clearInterval(timer.value)
        message.value = '获取验证码'
        count.value = 60
        return
      }
      count.value = count.value - 1
      message.value = `${ count.value }秒后重新发送`
    }, 1000)
  }
}

// 提交登录
const enterLogin = async() => {
  let res
  if (loginType.value === 'password') {
    res = await userLoginByPhoneOrUsernameApi({ username: formData.value.username, password: formData.value.password })
  } else {
    res = await userLoginByPhoneApi({ phone: formData.value.username, captcha: formData.value.password })
  }
  if (res.code === 0) {
    userStore.userinfo = res.data.user
    userStore.setToken(res.data.token)
    userStore.isLogin = true
    // await router.push({ path: path })
    window.location.href = '/#' + path
  }
}

// 初始化表单
const initForm = () => {
  formData.value = {
    password: '',
    username: '',
  }
}

watchEffect(() => {
  if (formData.value.username !== '' && formData.value.password !== '' && isActiveInfo.value) {
    isDisabled.value = true
  } else {
    isDisabled.value = false
  }
})
</script>

<style scoped lang="scss">
@import '@/style/login.scss';
</style>
