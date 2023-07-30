<template>
  <el-container class="layout-container">
    <el-aside :width="asideWidth" class="aside-container">
      <div class="zp-title" :class="{hidden: asideStore.collapseType}">
        <img
          src="https://img0.baidu.com/it/u=3368678403,249914024&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1690822800&t=b5d263590f4770cab940422bea705612"
          alt=""
        >
        <p v-if="!asideStore.collapseType">kva 后台管理系统</p>
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
        </div>
        <div class="head-right">
          <div v-if="!isFullscreen" @click="changeFull">
            <svg t="1690713183321" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="11828" width="32" height="32"><path d="M285.866667 810.666667H384v42.666666H213.333333v-170.666666h42.666667v98.133333l128-128 29.866667 29.866667-128 128z m494.933333 0l-128-128 29.866667-29.866667 128 128V682.666667h42.666666v170.666666h-170.666666v-42.666666h98.133333zM285.866667 256l128 128-29.866667 29.866667-128-128V384H213.333333V213.333333h170.666667v42.666667H285.866667z m494.933333 0H682.666667V213.333333h170.666666v170.666667h-42.666666V285.866667l-128 128-29.866667-29.866667 128-128z" fill="#515151" p-id="11829" /></svg>
          </div>
          <div v-if="isFullscreen" @click="changeFull">
            <svg t="1690713327556" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12271" width="32" height="32"><path d="M354.133333 682.666667H256v-42.666667h170.666667v170.666667H384v-98.133334L243.2 853.333333l-29.866667-29.866666L354.133333 682.666667z m358.4 0l140.8 140.8-29.866666 29.866666-140.8-140.8V810.666667h-42.666667v-170.666667h170.666667v42.666667h-98.133334zM354.133333 384L213.333333 243.2l29.866667-29.866667L384 354.133333V256h42.666667v170.666667H256V384h98.133333z m358.4 0H810.666667v42.666667h-170.666667V256h42.666667v98.133333L823.466667 213.333333l29.866666 29.866667L712.533333 384z" fill="#515151" p-id="12272" /></svg>
          </div>
          <el-dropdown style="cursor: pointer">
            <div class="el-dropdown-link dropdown-con">
              <div>
                <el-avatar
                  :size="35"
                  style="margin-right: 10px"
                  src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
                />
              </div>
              <span>系统管理员</span>
              <el-icon class="el-icon--right" style="position: relative;top:1px;margin-left: 8px;">
                <arrow-down />
              </el-icon>
            </div>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>
                  <span style="font-weight: 600;">
                    当前角色：超级管理员
                  </span>
                </el-dropdown-item>
                <el-dropdown-item>
                  <span>
                    切换为：系统商家角色
                  </span>
                </el-dropdown-item>
                <el-dropdown-item>
                  <span>
                    切换为：系统测试角色
                  </span>
                </el-dropdown-item>
                <el-dropdown-item icon="avatar" divided>个人信息</el-dropdown-item>
                <el-dropdown-item icon="reading-lamp" divided>登 出</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>
      <el-main class="main-con">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import Aside from './aside/index.vue'
import variables from '@/style/variables.module.scss'
import screenfull from 'screenfull'
import { computed, ref } from 'vue'
import { useASideStore } from '@/store/model/aside.js'

const asideStore = useASideStore()
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
</script>

<style scoped lang="scss">
@import "@/style/variables.module";

.layout-container {
  position: relative;
  width: 100%;
  height: 100%;
}

.aside-container {
  height: 100%;
}

.container {
  width: calc(100% - $AsideWidth);
  height: 100%;
  position: fixed;
  top: 0;
  right: 0;
  z-index: 9;
  transition: all 0.3s;

  &.hiddenContainer {
    transition: all 0.3s;
    width: calc(100% - $AsideHiddenWidth);
  }
}

.header-con {
  box-sizing: border-box;
  border-bottom: 1px solid rgba(0, 0, 0, 0.11);
  background-color: white;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;

  .head-left {
    margin-left: 14px;

    div {
      .icon {
        width: 30px;
        height: 30px;
        cursor: pointer;
      }
    }
  }

  .head-right {
    margin-right: 20px;
    display: flex;
    align-items: center;
    .icon {
      cursor: pointer;
      margin-right: 12px;
      width: 22px;
      height: 22px;
    }
  }
}

.zp-title {
  height: $TitleHeight;
  display: flex;
  align-items: center;
  background-color: white;

  img {
    margin-left: 14px;
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }

  p {
    margin: 0 10px 0 10px;
    font-size: 20px;
    font-weight: 700;
  }

  &.hidden {
    justify-content: center;

    img {
      margin-left: 0;
    }
  }
}

.dropdown-con {
  display: flex;
  align-items: center;
}

:deep(.el-header) {
  padding: 0;
}
// 下拉框黑边问题
:deep(:focus-visible) {
  outline: none;
}
.main-con {
  margin: 10px 15px;
}
</style>
