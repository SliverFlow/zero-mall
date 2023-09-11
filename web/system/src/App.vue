<template>
  <div id="app">
    <el-config-provider :locale="zhCn">
      <router-view />
    </el-config-provider>

  </div>
</template>

<script setup>

import { zhCn } from 'element-plus/es/locale/index'
</script>

<style lang="scss">
// 引入初始化样式
@import '@/style/main.scss';
@import '@/style/base.scss';
@import '@/style/mobile.scss';

#app {
  background: #f0f2f5;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}

.el-button {
  font-weight: 400 !important;
}
@font-face {
  font-family: "fm";
  src: url('@/assets/汉仪竹心楷.ttf');
}
</style>

