import { createRouter, createWebHashHistory } from 'vue-router'
import Npregress from 'nprogress'
import 'nprogress/nprogress.css'
import { useUserStore } from '@/store/model/user.js'

Npregress.configure({ showSpinner: false })

const routes = [
  {
    name: 'Login',
    path: '/login',
    component: () => import('@/view/login/index.vue')
  },
  {
    name: 'Index',
    path: '/',
    redirect: '/layout/dashboard',
    component: () => import('@/view/login/index.vue')
  },
  {
    name: 'Layout',
    path: '/layout',
    redirect: '/layout/dashboard',
    component: () => import('@/view/layout/index.vue'),
    children: [
      {
        name: 'Dashboard',
        path: 'dashboard',
        component: () => import('@/view/dashboard/index.vue')
      }
    ]
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
  const useStore = useUserStore()
  // 判断是否登录
  if (!useStore.isLogin && to.name !== 'Login') {
    // 这里的query就是为了记录用户最后一次访问的路径，这个路径是通过to的参数获取
    // 后续在登录成功以后，就可以根据这个path的参数，然后调整到你最后一次访问的路径
    router.push({ name: 'Login', query: { 'path': to.path }})
    return
  }
  // 继续进行
  next()
})

router.afterEach(() => {
  // 关闭进度条
  Npregress.done()
})

export default router
