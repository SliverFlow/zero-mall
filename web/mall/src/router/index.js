import { createRouter, createWebHashHistory } from 'vue-router'
import { useUserStore } from '@/pinia/model/user.js'
import { ElMessageBox } from 'element-plus'

const routes = [
  {
    path: '/',
    name: 'Layout',
    component: () => import('@/view/layout/index.vue'),
    children: [
      {
        path: '',
        name: 'Index',
        component: () => import('@/view/index/index.vue'),
      }, {
        path: 'product/list',
        name: 'ProductList',
        component: () => import('@/view/product/index.vue')
      }, {
        path: 'product/detail',
        name: 'ProductDetail',
        component: () => import('@/view/product/detail.vue')
      }
    ]
  }, {
    path: '/',
    name: 'UserLayout',
    component: () => import('@/view/user/layout.vue'),
    children: [
      {
        path: 'login',
        name: 'Login',
        component: () => import('@/view/user/login/index.vue')
      },
      {
        path: 'register',
        name: 'Register',
        component: () => import('@/view/user/register/index.vue')
      }
    ]
  },
  {
    path: '/cart',
    name: 'Cart',
    component: () => import('@/view/cart/index.vue')
  }, {
    path: '/order',
    name: 'Order',
    component: () => import('@/view/order/index.vue')
  }
]

const router = createRouter({
  routes,
  history: createWebHashHistory()
})

router.beforeEach(async(to) => {
  if (to.name === 'Login' || to.name === 'Register') {
    return true
  }
  const userStore = useUserStore()

  if (to.name === 'Order' || to.name === 'Cart') {
    if (!userStore.isLogin) {
      const res = await ElMessageBox.confirm(
        '加入购物车需要您登录账号，是否前往进入登录界面',
        '提示',
        {
          confirmButtonText: '确认',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
      if (res === 'confirm') {
        return { name: 'Login', query: { 'path': to.path }}
      } else {
        return false
      }
    } else {
      return true
    }
  }
})

export default router
