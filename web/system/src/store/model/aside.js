import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useASideStore = defineStore('aside', () => {
  // 菜单展开类型
  const collapseType = ref(false)
  // 修改菜单展开类型
  const changeCollapse = () => {
    collapseType.value = !collapseType.value
  }
  return {
    collapseType, changeCollapse
  }
})
