import { createRouter, createWebHashHistory } from 'vue-router'
import Npregress from 'nprogress'
import 'nprogress/nprogress.css'
import { useUserStore } from '@/store/model/user.js'
import { useRouterStore } from '@/store/model/router.js'

export const ng = Npregress.configure({ showSpinner: false })

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

router.beforeEach(async(to) => {
  // 开启进度条
  ng.start()
  const useStore = useUserStore()
  if (!useStore.isLogin && to.name === 'Login') {
    return true
  }

  // 判断是否登录
  if (!useStore.isLogin && to.name !== 'Login') {
    // 这里的query就是为了记录用户最后一次访问的路径，这个路径是通过to的参数获取
    // 后续在登录成功以后，就可以根据这个path的参数，然后调整到你最后一次访问的路径
    return { name: 'Login', query: { 'path': to.path }}
  }
  const routerStore = useRouterStore()
  // await routerStore.setAsyncRouter()
  await routerStore.setLocalAsyncRouter()
  // 如果刷新出现空白的问题，那么就使用下面这行代码
  if (!to.name && hasRoute(to)) {
    return { ...to }
  }

  if (useStore.isLogin && to.name === 'Login') {
    return { name: 'Dashboard' }
  }

  // 继续进行
  return true
})

router.afterEach(() => {
  // 关闭进度条
  ng.done()
})

// 判断当前路由是否存在动态添加的路由数据中
function hasRoute(to) {
  const item = router.getRoutes().find((item) => item.path === to.path)
  return !!item
}

export default router
