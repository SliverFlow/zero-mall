import { defineStore } from 'pinia'
import { ref } from 'vue'
import { loginApi } from '@/api/base.js'

export const useUserStore = defineStore('user', () => {
  const token = ref(window.localStorage.getItem('token') || '')
  const isLogin = ref(false)
  // 用户信息
  const userInfo = ref({})
  // 存放 token
  const setToken = (val) => {
    token.value = val
    window.localStorage.setItem('token', val)
  }

  // 登录
  const login = async(data) => {
    // 开始登录
    const res = await loginApi(data)
    if (res.code === 0) {
      // 登录成功
      setToken(res.data.token)
      // 设置用户信息
      userInfo.value = res.data.user
      isLogin.value = true
      return true
    } else {
      return false
    }
  }

  // 退出方法
  const logout = () => {
    // 将 token 置空
    setToken('')
    isLogin.value = false
    window.sessionStorage.clear()
    return true
  }

  return {
    login,
    logout,
    token,
    userInfo,
    setToken,
    isLogin,
  }
}, {
  persist: {
    key: 'zp-user-store',
    storage: window.sessionStorage
  }
})

