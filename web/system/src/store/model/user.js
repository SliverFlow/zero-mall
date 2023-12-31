import { defineStore } from 'pinia'
import { ref } from 'vue'
import { loginApi } from '@/api/system/base.js'
import { useRouterStore } from '@/store/model/router.js'
import router from '@/router/index.js'
import { useMenuStore } from '@/store/model/menu.js'

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

      const routerStore = useRouterStore()
      await routerStore.setAsyncRouter()
      routerStore.asyncRouterList.forEach(i => router.addRoute('Layout', i))

      isLogin.value = true
      return true
    } else {
      return false
    }
  }

  // 退出方法
  const logout = async() => {
    // const res = await logoutApi()
    const res = { code: 0 }
    if (res['code'] === 0) {
      // ElMessage({
      //   message: res['message'],
      //   type: 'success',
      //   showClose: true,
      // })
      // 将 token 置空
      setToken('')
      isLogin.value = false
      localStorage.removeItem('zp-aside-store')
      localStorage.removeItem('zp-router-store')
      localStorage.removeItem('token')
      const asideStore = useMenuStore()
      asideStore.tabList = [{ title: '仪表盘', path: '/layout/dashboard' }]
      return true
    }
    return false
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

