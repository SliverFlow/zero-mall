<template>
  <el-tabs
    v-model="activeTabPath"
    type="card"
    :closable="!(tabList.length === 1 && tabList[0].path === '/layout/dashboard')"
    class="tab-con"
    @tab-click="tabClick"
    @tab-remove="removeTab"
  >
    <el-tab-pane
      v-for="item in tabList"
      :key="item.path"
      :label="item.title"
      :name="item.path"
    >
    </el-tab-pane>
  </el-tabs>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMenuStore } from '@/store/model/menu.js'

const menuStore = useMenuStore()
const route = useRoute()
const router = useRouter()
// 标签页列表
const tabList = ref(menuStore.tabList)
// 当前激活的标签页
const activeTabPath = ref(menuStore.activeTabPath)

// 添加新标签
const addTab = () => {
  const { path, meta } = route
  const flag = tabList.value.some(item => item.path === path)
  if (!flag) {
    tabList.value.push({
      title: meta.title,
      path,
    })
  }
}

// 删除标签
const removeTab = (e) => {
  const index = tabList.value.findIndex(i => i.path === e)
  tabList.value.splice(index, 1)

  const nextTab = tabList.value[index - 1] || tabList.value[index + 1]
  if (nextTab) {
    router.push({ path: nextTab.path })
  }

  if (tabList.value.length === 0) {
    tabList.value.push({ title: '仪表盘', path: '/layout/dashboard' })
    activeTabPath.value = '/layout/dashboard'
    router.push({ path: '/layout/dashboard' })
  }
}

// 标签页点击事件
const tabClick = (e) => {
  router.push({ path: e.props.name })
}

// 监听路径变化
watch(route, () => {
  addTab()
  activeTabPath.value = route.path
}, { immediate: true })
// 监听标签页变化
watch(tabList.value, () => {
  if (tabList.value.length === 0) {
    tabList.value.push({ title: '仪表盘', path: '/layout/dashboard' })
    activeTabPath.value = '/layout/dashboard'
    router.push({ path: '/layout/dashboard' })
  }
}, { immediate: true })
</script>
<style lang="scss" scoped>
.tab-con {
  background-color: white;
  height: 40px;
}
</style>
