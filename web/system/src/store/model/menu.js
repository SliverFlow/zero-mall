import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMenuStore = defineStore('aside', () => {
  // 菜单展开类型
  const collapseType = ref(false)
  // 修改菜单展开类型
  const changeCollapse = () => {
    collapseType.value = !collapseType.value
  }
  const tabList = ref([{ title: '仪表盘', path: '/layout/dashboard' }])
  const activeTabPath = ref()
  return {
    collapseType, changeCollapse, tabList, activeTabPath
  }
}, {
  persist: {
    key: 'zp-aside-store',
    storage: window.localStorage
  }
})
