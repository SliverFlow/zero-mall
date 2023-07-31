export const menuListA = [
  {
    id: 1,
    parentId: 0,
    title: '仪表盘',
    icon: 'Odometer',
    name: 'Dashboard',
    path: '/layout/dashboard',
    component: 'dashboard/index.vue',
    meta: { title: '仪表盘' },
    children: []
  },
  {
    id: 2,
    parentId: 0,
    title: '超级管理员',
    icon: 'User',
    name: 'System',
    path: '/layout/system',
    component: 'system/index.vue',
    meta: { title: '超级管理员' },
    children: [
      {
        parentId: 2,
        name: 'User',
        title: '用户管理',
        icon: 'UserFilled',
        path: '/layout/system/user',
        component: 'system/user/index.vue',
        meta: { title: '用户管理' },
      }, {
        parentId: 2,
        name: 'Menu',
        title: '菜单管理',
        icon: 'Memo',
        path: '/layout/system/menu',
        component: 'system/index.vue',
        meta: { title: '菜单管理' },
      }
    ]
  },
  {
    id: 6,
    parentId: 0,
    title: '关于我们',
    icon: 'InfoFilled',
    name: 'About',
    path: '/layout/about',
    component: 'about/index.vue',
    meta: { title: '关于我们' },
    children: []
  }
]
