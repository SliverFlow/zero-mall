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
        <span v-if="loginType === 'password'" :class="{isActive: usernameFocus || usernameInfo}">手机号码 / 用户登录名</span>
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
            <Unlock v-if="isLock" />
            <lock v-if="!isLock" />
          </el-icon>
        </div>
        <span v-if="passwordInfo && loginType ==='password'" class="info">请输入密码</span>
        <span v-if="passwordInfo && loginType ==='captcha'" class="info">请输入验证码</span>
      </div>
      <div class="info">
        <el-checkbox v-model="isActiveInfo" :false-label="false" :true-label="true" size="large" />
        <span>已阅读并同意 zero-mall 用户协议 和 隐私政策</span>
      </div>
      <button :class="{'is-disabled': !isDisabled}" :disabled="isDisabled">登录</button>
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

// 发送验证码消息
const sendCaptchaMessage = () => {
  if (formData.value.phone === '') {
    alert('请输入手机号')
    return
  }
  message.value = `${count.value}秒后重新发送`
  if (timer.value) clearInterval(timer.value)
  timer.value = setInterval(() => {
    if (count.value <= 1) {
      clearInterval(timer.value)
      message.value = '获取验证码'
      count.value = 60
      return
    }
    count.value = count.value - 1
    message.value = `${count.value}秒后重新发送`
  }, 1000)
}

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
.captcha {
  position: absolute;
  right: 0;
  top: 18px;
  font-size: 15px;
  margin-right: 10px;
  color: #ff6700;
  cursor: pointer;
}

</style>
