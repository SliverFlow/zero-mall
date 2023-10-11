<template>
  <el-dropdown style="position: relative;top:-1px">
    <span>
      {{ userinfo.nickname }}
      <el-icon style="margin-left: 2px;position: relative;top: 2px">
        <arrow-down />
      </el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item>个人中心</el-dropdown-item>
        <el-dropdown-item>评价晒单</el-dropdown-item>
        <el-dropdown-item>我的喜欢</el-dropdown-item>
        <el-dropdown-item divided @click="logout">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>
<script setup>
import { useUserStore } from '@/pinia/model/user.js'
import { useRouter } from 'vue-router'

const userStore = useUserStore()
const userinfo = userStore.userinfo
const logout = () => {
  userStore.logout()
  window.location.reload()
}
</script>

<style lang="scss" scoped>
// 下拉框黑边问题
:deep(:focus-visible) {
  outline: none;
}

:deep(.el-dropdown-menu__item) {
  font-size: 12px;
}

:deep(.el-dropdown-menu__item):hover {
  color: #ff6700;
  background-color: rgba(204, 204, 204, 0.17);
}
</style>
