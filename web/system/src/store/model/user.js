import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(window.localStorage.getItem('token') || '')

  // 用户信息
  const userInfo = ref({})

  // 存放 token
  const setToken = (val) => {
    token.value = val
    window.localStorage.setItem('token', val)
  }

  // 登录
  const login = () => {
  }

  // 退出方法
  const logout = () => {
    // 将 token 置空
    setToken('')
  }

  return {
    login,
    logout,
    token,
    userInfo,
    setToken
  }
})

