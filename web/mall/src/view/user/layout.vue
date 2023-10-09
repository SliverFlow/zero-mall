<template>
  <div class="panel">
    <div class="panel-left" />
    <div class="panel-right">
      <div class="header">
        <div class="left">
          <img src="https://cdn.cnbj1.fds.api.mi-img.com/mi.com-assets/shop/img/logo-mi2.png" alt="">
          <span>zero-mall</span>
        </div>
        <div class="right">
          <span>用户协议</span>
          <span>隐私政策</span>
          <span>帮助中心</span>
          <span>|</span>
          <span>中文简体</span>
        </div>
      </div>
      <div style="height: calc(100vh - 100px);width: 100%">
        <div class="panel">
          <div class="login_panel">
            <el-tabs v-model="activeName" @tab-change="tabChange">
              <el-tab-pane label="登录" name="Login" />
              <el-tab-pane label="注册" name="Register" />
            </el-tabs>
            <router-view />
          </div>
        </div>
      </div>
      <div class="footer">
        <el-affix position="bottom" :offset="20">
          <div class="github">
            <p>@ Powered by</p>
            <span @click="toGithub">Zero-Mall</span>
            <p>Copyright</p>
            <span @click="toGithub">SliverFlow</span>
          </div>
        </el-affix>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const activeName = ref(route.name)

const tabChange = (name) => {
  router.push({ name: name })
}
</script>

<style scoped lang="scss">
@import '@/style/user_layout';

</style>
