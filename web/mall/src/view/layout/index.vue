<template>
  <div>
    <header>
      <div>
        <div class="left">
          <a href="#/">mall官网</a>
          <span>|</span>
          <a href="">mall商城</a>
          <span>|</span>
          <a href="">ZERO</a>
          <span>|</span>
          <a href="">LoT</a>
          <span>|</span>
          <a href="">云服务</a>
          <span>|</span>
          <a href="">gorm</a>
          <span>|</span>
          <a href="">资质证照</a>
          <span>|</span>
          <a href="">小爱开放平台</a>
          <span>|</span>
          <a href="">协议规则</a>
          <span>|</span>
          <a href="" style="color: #ff6700;">成为商家</a>
        </div>
        <div class="right">
          <a v-if="!isLogin" style="cursor: pointer" @click="toLogin">登录</a>
          <span v-if="!isLogin">|</span>
          <a v-if="!isLogin" style="cursor: pointer" @click="toRegister">注册</a>
          <user-info-com v-if="isLogin" style="color: #cecece"/>
          <span>|</span>
          <a @click.prevent="toOrder" style="cursor: pointer">订单列表</a>
          <div class="shop-car" @click="toCart" @mouseover="show = !show" @mouseout="show = !show">
            <svg
              t="1694517489223"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              p-id="4034"
              width="200"
              height="200"
            >
              <path
                d="M352.456912 832.032253c-35.434907 0-63.989249 28.554342-63.989249 63.989249 0 35.434907 28.554342 63.989249 63.989249 63.989249s63.989249-28.554342 63.989249-63.989249C416.446162 860.586595 387.891819 832.032253 352.456912 832.032253L352.456912 832.032253z"
                fill="#b0b0b0"
                p-id="4035"
              />
              <path
                d="M800.55367 832.032253c-35.434907 0-63.989249 28.554342-63.989249 63.989249 0 35.434907 28.554342 63.989249 63.989249 63.989249s63.989249-28.554342 63.989249-63.989249C864.54292 860.586595 835.816563 832.032253 800.55367 832.032253L800.55367 832.032253z"
                fill="#b0b0b0"
                p-id="4036"
              />
              <path
                d="M864.026877 800.037628 344.200235 800.037628c-46.099782 0-86.695112-36.466991-92.199563-83.082815l-54.356459-382.043339L166.853687 156.360826c-1.892155-15.653284-16.169326-28.382328-29.930455-28.382328L95.983874 127.978498c-17.717453 0-31.994625-14.277171-31.994625-31.994625s14.277171-31.994625 31.994625-31.994625l40.767344 0c46.615824 0 87.727196 36.466991 93.403662 83.082815l30.790526 177.86259L315.473879 708.698135c1.720141 14.793214 15.309256 27.350244 28.726356 27.350244l519.826642 0c17.717453 0 31.994625 14.277171 31.994625 31.994625S881.744331 800.037628 864.026877 800.037628z"
                fill="#b0b0b0"
                p-id="4037"
              />
              <path
                d="M384.279523 672.05913c-16.685369 0-30.618512-12.729044-31.82261-29.586427-1.376113-17.545439 11.868974-33.026709 29.586427-34.230808l434.163615-31.994625c15.997312-0.172014 29.414413-12.55703 31.134554-26.834201l50.400134-288.295649c1.204099-10.664875-1.720141-22.533848-8.084663-29.758441-4.128339-4.644381-9.288762-7.052579-15.309256-7.052579L319.946246 224.3064c-17.717453 0-31.994625-14.277171-31.994625-31.994625S302.400806 159.973123 319.946246 159.973123l554.05745 0c24.426004 0 46.959852 10.148833 63.301193 28.554342 18.749538 21.157736 27.178229 50.744163 23.565933 81.706703l-50.400134 288.467663c-5.504452 44.895683-45.927768 81.362674-92.027549 81.362674l-431.755417 31.82261C385.82765 671.887116 384.967579 672.05913 384.279523 672.05913z"
                fill="#b0b0b0"
                p-id="4038"
              />
            </svg>
            <span>购物车</span>
            <span>（0）</span>
<!--            <el-collapse-transition style="transition: all 0.3s">-->
<!--              <div v-show="show" class="show-box">-->
<!--                <div>购物车中还没有商品，赶紧选购吧!</div>-->
<!--              </div>-->
<!--            </el-collapse-transition>-->
          </div>
        </div>
      </div>
    </header>
    <nav>
      <div>
        <img src="https://cdn.cnbj1.fds.api.mi-img.com/mi.com-assets/shop/img/logo-mi2.png" alt="">
        <div class="search-box">
          <div>
            <a href="">手机</a>
            <a href="">笔记本</a>
            <a href="">电视</a>
            <a href="">智慧屏</a>
            <a href="">穿戴</a>
            <a href="">家电</a>
            <a href="">服务中心</a>
            <a href="">社区</a>
          </div>
          <div class="search" style="width: 300px;height: 48px">
            <input
              v-model="keyWord"
              type="text"
              placeholder="输入商品名称"
              :class="{icon_fover: isFocus}"
              @focus="isFocus = !isFocus"
              @focusout="isFocus = !isFocus"
            >
            <el-icon
              :size="25"
              class="search-icon"
              :class="{icon_fover: isFocus,search_icon_hover: isFocus}"
              @click="search"
            >
              <Search/>
            </el-icon>
            <div v-if="isFocus" class="input_word">
              <a href="">全部商品</a>
              <a href="">手机</a>
              <a href="">充电宝</a>
              <a href="">空调</a>
              <a href="">洗衣机</a>
              <a href="">笔记本</a>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <div class="content">
      <router-view/>
    </div>
    <mall-footer/>
  </div>
</template>

<script setup>
import { ref, provide } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { useRoute, useRouter } from 'vue-router'
import MallFooter from '@/view/layout/footer/mallFooter.vue'
import { useActiveStore } from '@/pinia/model/active.js'
import { useUserStore } from '@/pinia/model/user.js'
import UserInfoCom from '@/components/userInfo/UserInfoCom.vue'

const show = ref(false)
const isFocus = ref(false)
const router = useRouter()
const route = useRoute()
const keyWord = ref('')
const userStore = useUserStore()
const isLogin = ref(userStore.isLogin)


const toLogin = () => {
  router.push({ name: 'Login', query: { path: route.path } })
}
const toCart = () => {
  router.push({ name: 'Cart' })
}

const toRegister = () => {
  router.push({ name: 'Register' })
}

const search = () => {
  if (router.currentRoute.value.name !== 'ProductList') {
    router.push({ name: 'ProductList' })
  } else {
    useActiveStore().active = true
  }
}

const toOrder = () => {
  router.push({ name: 'Order' })
}

provide('keyWord', keyWord)

</script>

<style scoped lang="scss">
@import '@/style/layout.scss';

</style>
