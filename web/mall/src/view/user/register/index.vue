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
          <p class="captcha">获取验证码</p>
        </div>
        <span :class="{isActive: captchaFocus || captchaInfo}">验证码</span>
        <span v-if="captchaInfo" class="info">请输入验证码</span>
      </div>
      <div class="info">
        <el-checkbox v-model="isActiveInfo" :false-label="false" :true-label="true" size="large" />
        <span>已阅读并同意 zero-mall 用户协议 和 隐私政策</span>
      </div>
      <button>注册</button>
      <div class="phone">
        <span>收不到验证码？</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

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

</script>

<style scoped lang="scss">
@import '@/style/login.scss';

.isPhoneActive {
  left: calc(30% + 16px) !important;
  top: 8px !important;
  font-size: 14px !important;
}
.input_item {

  .phone_span {
    position: relative;
    left: calc(30% + 10px) !important;
  }
}
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
