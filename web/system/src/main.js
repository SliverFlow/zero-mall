import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/notification/style/css'
import 'element-plus/es/components/message-box/style/css'
import './style/element_visiable.scss'

import 'element-plus/theme-chalk/src/index.scss'
import { createApp } from 'vue'
import App from './App.vue'
import store from './store/index.js'
import router from '@/router/index.js'

import './core/gin-vue-admin'
import run from '@/core/gin-vue-admin.js'
// import 'virtual:windi.css'

const app = createApp(App)
app.use(store)
app.use(router)
app.use(run)
app.mount('#app')

export default app
