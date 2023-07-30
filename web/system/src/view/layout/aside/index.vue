<template>
  <el-menu
    :unique-opened="true"
    :collapse-transition="false"
    :collapse="asideStore.collapseType"
    router
    class="aside-con"
  >
    <el-menu-item index="/layout/dashboard">
      <el-icon>
        <setting />
      </el-icon>
      <span>仪表盘</span>
    </el-menu-item>
  </el-menu>
</template>

<script setup>
import { useASideStore } from '@/store/model/aside.js'

const asideStore = useASideStore()

</script>

<style scoped lang="scss">
@import "@/style/variables.module";

.aside-con {
  height: calc(100% - $TitleHeight);
}
</style>
