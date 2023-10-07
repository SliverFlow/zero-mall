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
        <span :class="{isActive: usernameFocus}">手机号码 / 用户登录名</span>
        <span v-if="usernameInfo" class="info">请输入账号</span>
      </div>
      <div class="input_item">
        <input
          v-model="formData.password"
          :class="{isFocus: passwordFocus}"
          :type="!isLock ? 'password' : 'text'"
          @focus="passwordFocus = true"
          @focusout="passwordFocusOut"
        >
        <span :class="{isActive: passwordFocus || passwordInfo}">密码</span>
        <div class="input_icon" @click="passwordFocus = true;isLock = !isLock;">
          <el-icon size="22" color="#959595">
            <Unlock v-if="isLock" />
            <lock v-if="!isLock" />
          </el-icon>
        </div>
        <span v-if="passwordInfo" class="info">请输入密码</span>
      </div>
      <div class="info">
        <el-checkbox v-model="isActiveInfo" :false-label="false" :true-label="true" size="large" />
        <span>已阅读并同意 zero-mall 用户协议 和 隐私政策</span>
      </div>
      <button>登录</button>
      <div class="phone">
        <span>忘记密码？</span>
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
import { ref } from 'vue'

const passwordFocus = ref()
const passwordInfo = ref(false)
const usernameFocus = ref()
const usernameInfo = ref(false)
const isLock = ref(false)
const isActiveInfo = ref(false)

const formData = ref({
  password: '',
  username: ''
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

</script>

<style scoped lang="scss">
@import '@/style/login.scss';
</style>
