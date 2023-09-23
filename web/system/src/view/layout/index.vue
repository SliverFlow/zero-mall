<template>
  <el-container class="layout-container">
    <el-aside :width="asideWidth" class="aside-container">
      <div class="zp-title" :class="{hidden: asideStore.collapseType}">
        <img
          src="https://www.gin-vue-admin.com/img/logo.png"
          alt=""
        >
        <p v-if="!asideStore.collapseType && userInfo.role === 3" style="font-family: 'fm'">zero-mall&nbsp;&nbsp;后台管理</p>
        <p v-if="!asideStore.collapseType && userInfo.role === 2" style="font-family: 'fm'">zero-mall&nbsp;&nbsp;商家管理</p>
      </div>
      <Aside />
    </el-aside>
    <el-container class="container " :class="{hiddenContainer: !asideStore.collapseType}">
      <el-header class="header-con">
        <div class="head-left">
          <div v-if="!asideStore.collapseType" @click="asideStore.changeCollapse()">
            <svg
              t="1690699441652"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="10466"
              width="32"
              height="32"
            >
              <path
                d="M550.4 494.933333l192 192 29.866667-29.866666-162.133334-162.133334 162.133334-162.133333-29.866667-29.866667-192 192z m-256 0l192 192 29.866667-29.866666-162.133334-162.133334 162.133334-162.133333-29.866667-29.866667-192 192z"
                fill="#515151"
                p-id="10467"
              />
            </svg>
          </div>
          <div v-if="asideStore.collapseType" @click="asideStore.changeCollapse()">
            <svg
              t="1690699399407"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="5031"
              width="32"
              height="32"
            >
              <path
                d="M516.266667 494.933333l-192 192-29.866667-29.866666 162.133333-162.133334-162.133333-162.133333 29.866667-29.866667 192 192z m256 0l-192 192-29.866667-29.866666 162.133333-162.133334-162.133333-162.133333 29.866667-29.866667 192 192z"
                fill="#515151"
                p-id="5032"
              />
            </svg>
          </div>
          <div class="title"><span>{{ route.matched[1].path === route.path ? '' : route.matched[1].meta.title }} <span
            v-if="route.matched[1].path !== route.path"
            class="fn-span"
          >/</span> {{ route.meta.title }}</span></div>
        </div>
        <div class="head-right">
          <div @click="toGithub">
            <svg
              t="1690716772490"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="14449"
              width="32"
              height="32"
            >
              <path
                d="M427.392 853.504a61.44 61.44 0 0 1-1.450667 15.530667 92.586667 92.586667 0 0 1-10.965333 27.605333c-15.061333 25.301333-40.661333 42.154667-73.642667 42.154667-77.653333 0-108.117333-38.101333-146.261333-133.504C169.216 740.693333 157.013333 725.461333 128 725.461333v-85.333333c77.653333 0 108.117333 38.101333 146.261333 133.504 25.856 64.597333 38.058667 79.786667 67.072 79.786667 0-12.373333-0.170667-23.296-0.512-38.144-0.853333-34.816-0.938667-41.941333 0.554667-51.2 0.64-20.352 5.888-34.773333 16.384-49.066667-95.232-20.736-159.445333-63.573333-196.309333-132.992l-13.824-32.426667C134.186667 510.848 128 467.072 128 416.426667c0-58.24 17.749333-110.336 50.944-153.856-10.368-41.386667-8.96-91.989333 13.909333-149.077334l7.466667-18.688 19.2-6.101333c2.56-0.853333 5.632-1.578667 9.301333-2.133333 37.290667-5.888 90.325333 8.106667 159.701334 52.48a565.930667 565.930667 0 0 1 127.274666-14.208c38.741333 0 77.226667 3.882667 114.048 11.605333 67.456-42.24 119.04-55.509333 155.306667-49.92 3.626667 0.597333 6.741333 1.322667 9.258667 2.133333l19.285333 6.101334 7.466667 18.773333c20.010667 50.218667 23.424 96.469333 16.128 136.96C875.434667 296.32 896 352.597333 896 416.512c0 53.888-3.84 94.378667-14.933333 133.802667l-11.733334 32.170666c-30.677333 69.333333-98.304 112.64-202.538666 133.248 10.837333 15.018667 15.872 30.250667 15.872 52.394667v42.666667l-0.042667 42.666666a13.013333 13.013333 0 0 0 0.341333 2.730667L682.666667 938.794667c-36.352 0-63.36-17.706667-76.672-45.653334a88.277333 88.277333 0 0 1-8.661334-40.277333v-84.736c0-3.584-0.128-3.797333-8.832-12.501333-23.296-23.296-33.834667-40.874667-33.834666-72.832v-38.613334l38.4-3.84c114.346667-11.52 176.512-43.221333 197.12-89.6l9.6-26.325333c7.68-27.562667 10.88-61.098667 10.88-107.946667 0-49.706667-17.365333-90.837333-50.218667-123.648L742.4 274.773333l7.381333-24.448c6.528-21.717333 8.106667-47.402667 1.152-76.714666a158.634667 158.634667 0 0 0-3.584 0.938666c-22.826667 6.4-51.370667 20.053333-85.76 43.008l-15.658666 10.453334-18.304-4.522667a467.754667 467.754667 0 0 0-111.829334-13.269333c-42.709333 0-84.906667 5.418667-123.946666 16.042666l-19.029334 5.205334-16.256-11.136c-35.541333-24.32-65.109333-38.826667-88.746666-45.568a158.293333 158.293333 0 0 0-4.864-1.28c-8.234667 33.92-4.992 61.781333 3.413333 82.688l9.984 25.088-18.346667 19.797333C228.693333 332.629333 213.333333 370.986667 213.333333 416.512c0 41.685333 4.864 76.202667 13.824 102.229333l11.178667 26.453334c27.904 52.352 87.168 83.84 192.853333 95.146666l38.144 4.096v38.357334c0 32-10.538667 49.536-33.834666 72.832-8.704 8.704-8.832 8.96-8.832 12.501333l-0.725334 7.893333c-0.512 2.56-0.512 9.258667 0.170667 37.205334 0.298667 12.8 0.469333 22.997333 0.554667 33.621333a33.962667 33.962667 0 0 1 0.725333 6.656z"
                p-id="14450"
                fill="#707070"
              />
            </svg>
          </div>
          <div v-if="!isFullscreen" @click="changeFull">
            <svg
              t="1690713183321"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="11828"
              width="32"
              height="32"
            >
              <path
                d="M285.866667 810.666667H384v42.666666H213.333333v-170.666666h42.666667v98.133333l128-128 29.866667 29.866667-128 128z m494.933333 0l-128-128 29.866667-29.866667 128 128V682.666667h42.666666v170.666666h-170.666666v-42.666666h98.133333zM285.866667 256l128 128-29.866667 29.866667-128-128V384H213.333333V213.333333h170.666667v42.666667H285.866667z m494.933333 0H682.666667V213.333333h170.666666v170.666667h-42.666666V285.866667l-128 128-29.866667-29.866667 128-128z"
                fill="#515151"
                p-id="11829"
              />
            </svg>
          </div>
          <div v-if="isFullscreen" @click="changeFull">
            <svg
              t="1690713327556"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="12271"
              width="32"
              height="32"
            >
              <path
                d="M354.133333 682.666667H256v-42.666667h170.666667v170.666667H384v-98.133334L243.2 853.333333l-29.866667-29.866666L354.133333 682.666667z m358.4 0l140.8 140.8-29.866666 29.866666-140.8-140.8V810.666667h-42.666667v-170.666667h170.666667v42.666667h-98.133334zM354.133333 384L213.333333 243.2l29.866667-29.866667L384 354.133333V256h42.666667v170.666667H256V384h98.133333z m358.4 0H810.666667v42.666667h-170.666667V256h42.666667v98.133333L823.466667 213.333333l29.866666 29.866667L712.533333 384z"
                fill="#515151"
                p-id="12272"
              />
            </svg>
          </div>
          <el-dropdown style="cursor: pointer">
            <div class="el-dropdown-link dropdown-con">
              <div>
                <el-avatar
                  :size="30"
                  style="margin-right: 10px"
                  :src="userInfo.avatar"
                />
              </div>
              <span>{{ userInfo.nickname }}</span>
              <el-icon class="el-icon--right" style="position: relative;top:0;margin-left: 8px;">
                <arrow-down />
              </el-icon>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>
                  <span style="font-weight: 600;">
                    当前角色：{{ userInfo.role === 3 ? '系统管理员' : '系统商家' }}
                  </span>
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="userInfo.username==='admin'"
                  @click="changeRole(1)"
                >
                  <span>
                    切换为：系统管理员
                  </span>
                </el-dropdown-item>
                <el-dropdown-item
                  v-if="userInfo.username==='admin'"
                  @click="changeRole(2)"
                >
                  <span>
                    切换为：系统商家角色
                  </span>
                </el-dropdown-item>
                <el-dropdown-item icon="avatar" divided>个人信息</el-dropdown-item>
                <el-dropdown-item icon="reading-lamp" divided @click="toLogout">登 出</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      <tab />
      <el-main style="padding: 10px 14px">
        <router-view class="router-con" />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import Aside from './aside/index.vue'
import Tab from './tab/index.vue'
import variables from '@/style/variables.module.scss'
import screenfull from 'screenfull'
import { computed, ref } from 'vue'
import { useMenuStore } from '@/store/model/menu.js'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/store/model/user.js'
import { userChangeRoleApi } from '@/api/system/user.js'
import { ElMessage } from 'element-plus'
import { useRouterStore } from '@/store/model/router.js'

const route = useRoute()
const router = useRouter()
const routerStore = useRouterStore()
// userStore
const userStore = useUserStore()
// 用户信息
const userInfo = userStore.userInfo
const asideStore = useMenuStore()
// aside 宽度
const asideWidth = computed(() => {
  return asideStore.collapseType ? variables['aside-hidden-width'] : variables['aside-width']
})
// 页面全屏状态
const isFullscreen = ref(screenfull.isFullscreen)
// 切换全屏
const changeFull = () => {
  screenfull.toggle()
  isFullscreen.value = !screenfull.isFullscreen
}
// 去往 github
const toGithub = () => {
  window.open('https://github.com')
}
const toLogout = async() => {
  const flag = await userStore.logout()
  if (flag) {
    setTimeout(() => {
      router.push({ name: 'Login' })
    }, 400)
  }
}

// 修改系统用户角色
const changeRole = async(val) => {
  const res = await userChangeRoleApi({ username: userInfo.username, role: val })
  if (res.code === 0) {
    ElMessage({
      message: '更换管理员角色成功',
      type: 'success',
      showClose: true,
    })
    // 重新加载路由标签
    asideStore.tabList.splice(0, asideStore.tabList.length)
    // 重新设置 router
    setTimeout(async() => {
      await routerStore.setAsyncRouter()
      routerStore.asyncRouterList.forEach(i => router.addRoute('Layout', i))
    }, 500)
  }
}
</script>

<style scoped lang="scss">
@import "@/style/variables.module";
@import "@/style/layout";
</style>
