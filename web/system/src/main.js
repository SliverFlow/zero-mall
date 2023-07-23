import { createApp } from 'vue'
import App from './App.vue'
import store from './store/index.js'
import router from '@/router/index.js'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './style/element_visiable.scss'
// element icon
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// 处理组件中的英文
import zhCn from 'element-plus/es/locale/lang/zh-cn'

const app = createApp(App)
app.use(store)
app.use(ElementPlus, { locale: zhCn })
app.use(router)
app.mount('#app')

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
