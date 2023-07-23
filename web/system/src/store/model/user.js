import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref(window.localStorage.getItem('token') || '')

  const userInfo = ref({})

  const setToken = (val) => {
    token.value = val
    window.localStorage.setItem('token', val)
  }

  const login = () => {
  }

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

