import { defineStore } from 'pinia'
import { menuListA } from '@/view/menu.js'
import { ref } from 'vue'
import { menuTreeListApi } from '@/api/menu.js'

export const useRouterStore = defineStore('router', () => {
  const menuList = ref(menuListA)
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
        meta: { title: i.title },
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
    asyncRouterList
  }
}, {
  persist: {
    key: 'zp-router-store',
    storage: window.sessionStorage
  }
})
