import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouterStore } from '@/store/model/router.js'

export const useUserStore = defineStore('user', () => {
  const token = ref(window.localStorage.getItem('token') || '')
  const isLogin = ref(false)
  // 用户信息
  const userInfo = ref({})
  // 用户菜单 列表
  // 存放 token
  const setToken = (val) => {
    token.value = val
    window.localStorage.setItem('token', val)
  }

  // 登录
  const login = (data) => {
    const routerStore = useRouterStore()
    routerStore.asyncRouter()
    isLogin.value = true
    return true
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
    setToken,
    isLogin,
  }
}, {
  persist: {
    key: 'zp-user-store',
    storage: window.localStorage
  }
})

