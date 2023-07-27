import { createRouter, createWebHashHistory } from 'vue-router'
import Npregress from 'nprogress'
import 'nprogress/nprogress.css'

Npregress.configure({ showSpinner: false })

const routes = [
  {
    name: 'Login',
    path: '/login',
    component: () => import('@/view/login/index.vue')
  },
  {
    name: 'Layout',
    path: '/layout',
    component: () => import('@/view/layout/index.vue')
  },
  {
    name: 'Dashboard',
    path: '/dashboard',
    component: () => import('@/view/dashboard/index.vue')
  },
  // {
  //   path: '/404',
  //   name: '404',
  //   component: () => import('@/view/error/404.vue')
  // },
  // {
  //   path: '/:pathMatch(.*)',
  //   redirect: '/404'
  // }
]
const router = createRouter({
  history: createWebHashHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  // 开启进度条
  Npregress.start()
  next()
})

router.afterEach(() => {
  // 关闭进度条
  Npregress.done()
})

export default router
