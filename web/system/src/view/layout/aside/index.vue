<template>
  <el-menu
    default-active="/layout/dashboard"
    :unique-opened="true"
    :collapse-transition="false"
    :collapse="asideStore.collapseType"
    router
    class="aside-con"
  >
    <el-sub-menu index="1">
      <template #title>
        <el-icon>
          <location />
        </el-icon>
        <span>Navigator One</span>
      </template>
      <el-menu-item index="/layout/dash11board">
        <el-icon>
          <setting />
        </el-icon>
        <span>仪表盘</span>
      </el-menu-item>
      <el-menu-item index="/layout/dashboard">
        <el-icon>
          <setting />
        </el-icon>
        <span>仪表盘</span>
      </el-menu-item>
    </el-sub-menu>

    <el-menu-item index="/layout/1">
      <el-icon>
        <setting />
      </el-icon>
      <span>仪表盘</span>
    </el-menu-item>

    <el-tooltip
      effect="light "
      content="仪表盘"
      placement="right"
    >
      <el-menu-item index="/layout/1">
        <el-icon>
          <setting />
        </el-icon>
        <span>仪表盘</span>
      </el-menu-item>
    </el-tooltip>

    <el-sub-menu index="1">
      <template #title>
        <el-icon>
          <location />
        </el-icon>
        <span>Navigator One</span>
      </template>

      <el-menu-item index="/layout/dashboard">
        <el-icon>
          <setting />
        </el-icon>
        <span>仪表盘</span>
      </el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script setup>
import { useASideStore } from '@/store/model/aside.js'

const asideStore = useASideStore()

</script>

<style scoped lang="scss">
@import "@/style/variables.module";
@import "@/style/aside";
.aside-con {
  height: calc(100% - $TitleHeight);
}

</style>
