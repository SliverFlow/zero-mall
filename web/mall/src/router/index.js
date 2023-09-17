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
      }
    ]
  }, {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue'),
  }, {
    path: '/bub/cart',
    name: 'Cart',
    component: () => import('@/view/cart/index.vue')
  }
]

const router = createRouter({
  routes,
  history: createWebHashHistory()
})

export default router
