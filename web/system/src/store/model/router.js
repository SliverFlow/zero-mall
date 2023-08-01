import { defineStore } from 'pinia'
import router from '@/router/index.js'
import { menuListA } from '@/view/menu.js'
import { ref } from 'vue'

export const useRouterStore = defineStore('router', () => {
  const menuList = ref(menuListA)
  const asyncRouterList = ref([])
  // 设置用户的路由
  const asyncRouter = () => {
    asyncRouterList.value = mapRouter(menuList.value)
    asyncRouterList.value.forEach(i => router.addRoute('Layout', i))
  }

  const modules = import.meta.glob('../../view/**/*.vue')
  // 获取路由数组
  const mapRouter = (val) => {
    return val.map(i => {
      const r = {
        name: i.name,
        path: i.path,
        meta: i.meta,
        component: modules['../../view/' + i.component]
      }
      if (i.children) {
        r.children = mapRouter(i.children)
      }
      return r
    })
  }
  return {
    asyncRouter,
    menuList,
    asyncRouterList
  }
}, {
  persist: {
    key: 'zp-router-store',
    storage: window.sessionStorage
  }
})
