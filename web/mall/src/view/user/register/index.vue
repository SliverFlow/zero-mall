<template>
  <div>
    <div class="login_panel">
      <div class="input_item">
        <input
          v-model="formData.country"
          type="text"
          disabled
        >
        <span class="isActive">国家/地区</span>
      </div>
      <div class="input_item" style="display: flex;">
        <div style="width: 30%;background-color: #f9f9f9;display: flex;flex-direction: column;">
          <span style="font-size: 14px;position: relative;top: 6px">国家码</span>
          <p style="font-size: 18px;position: relative;top: 8px;display: flex;left: 17px">+86</p>
        </div>
        <input
          v-model="formData.phone"
          style="width: 70%"
          :class="{isFocus: phoneFocus || phoneInfo}"
          type="text"
          @focus="phoneFocus = true"
          @focusout="phoneFocusOut"
        >
        <span class="phone_span" :class="{isPhoneActive: phoneFocus || phoneInfo}">手机号</span>
        <span v-if="phoneInfo" class="info">请输入电话号码</span>
      </div>
      <div class="input_item">
        <div style="display: flex;position: relative">
          <input
            v-model="formData.captcha"
            :class="{isFocus: captchaFocus || captchaInfo}"
            type="text"
            @focus="captchaFocus = true"
            @focusout="captchaFocusOut"
          >
          <a class="captcha" @click.prevent="sendCaptchaMessage">{{ message }}</a>
        </div>
        <span :class="{isActive: captchaFocus || captchaInfo}">验证码</span>
        <span v-if="captchaInfo" class="info">请输入验证码</span>
      </div>
      <div class="info">
        <el-checkbox v-model="isActiveInfo" size="large" />
        <span>已阅读并同意 zero-mall 用户协议 和 隐私政策</span>
      </div>
      <button :class="{'is-disabled': !isDisabled}" :disabled="isDisabled">注册</button>
      <div class="phone">
        <span>收不到验证码？</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watchEffect } from 'vue'

const captchaFocus = ref()
const captchaInfo = ref(false)
const phoneFocus = ref()
const phoneInfo = ref(false)
const isActiveInfo = ref(false)

const formData = ref({
  phone: '',
  captcha: '',
  country: '中国'
})

const captchaFocusOut = () => {
  if (formData.value.captcha === '') {
    captchaFocus.value = false
    captchaInfo.value = true
  } else {
    captchaFocus.value = true
    captchaInfo.value = false
  }
}
const phoneFocusOut = () => {
  if (formData.value.phone === '') {
    phoneFocus.value = false
    phoneInfo.value = true
  } else {
    phoneFocus.value = true
    phoneInfo.value = false
  }
}

// ------------------------------------------
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

watchEffect(() => {
  if (formData.value.phone !== '' && formData.value.captcha !== '' && isActiveInfo.value) {
    isDisabled.value = true
  } else {
    isDisabled.value = false
  }
})
</script>

<style scoped lang="scss">
@import '@/style/login.scss';
</style>
