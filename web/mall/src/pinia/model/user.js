import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const userinfo = ref({})
  const token = ref(window.localStorage.getItem('token') || '')
  const isLogin = ref(false)

  const setToken = (val) => {
    token.value = val
  }

  const getToken = () => {
    return token.value
  }

  // 登出操作
  const logout = () => {
    setToken('')
    isLogin.value = false
    localStorage.clear()
  }

  return {
    userinfo,
    setToken,
    getToken,
    isLogin,
    logout
  }
}, {
  persist: {
    key: 'zero-mall-user-store',
    storage: window.localStorage
  }
})

