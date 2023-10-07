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

<style>
* {
  margin: 0;
  padding: 0;
  list-style: none;
}

a {
  text-decoration: none;
  color: inherit;
}

#app {
  width: 100%;
  height: 100%;
}

::-webkit-scrollbar {
  display: none;
}
@font-face {
  font-family: "fm";
  src: url('@/assets/汉仪竹心楷.ttf');
}
</style>
