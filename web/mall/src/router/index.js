import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Layout',
    component: () => import('@/view/layout/index.vue'),
    children: [
      {
        path: 'index',
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
      },
      {
        path: 'cart',
        name: 'Cart',
        component: () => import('@/view/cart/index.vue')
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
  }
]

const router = createRouter({
  routes,
  history: createWebHashHistory()
})

export default router
