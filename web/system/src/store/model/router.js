import { defineStore } from 'pinia'
import { ref } from 'vue'
import { menuTreeListApi } from '@/api/system/menu.js'
import router from '@/router/index.js'

export const useRouterStore = defineStore('router', () => {
  const menuList = ref([])
  const asyncRouterList = ref([])
  // 设置用户的路由
  const setAsyncRouter = async() => {
    const res = await menuTreeListApi()
    if (res.code !== 0) {
      return false
    }
    menuList.value = res.data.list
    asyncRouterList.value = mapRouter(menuList.value)
    return true
  }

  const setLocalAsyncRouter = async() => {
    mapRouter(menuList.value).forEach(i => router.addRoute('Layout', i))
  }

  const modules = import.meta.glob('../../view/**/*.vue')
  // 获取路由数组
  const mapRouter = (val) => {
    if (val.length === 0) {
      return []
    }
    return val.map(i => {
      const r = {
        name: i.name,
        path: i.path,
        meta: { title: i.meta.title },
        component: modules['../../view/' + i.component]
      }
      if (i.children) {
        r.children = mapRouter(i.children)
      }
      return r
    })
  }
  return {
    setAsyncRouter,
    menuList,
    asyncRouterList,
    setLocalAsyncRouter
  }
}, {
  persist: {
    key: 'zp-router-store',
    storage: window.localStorage
  }
})
