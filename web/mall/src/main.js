import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router/index.js'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import store from '@/pinia/index.js'

const app = createApp(App)
app.use(router)
app.use(ElementPlus)
app.use(store)
app.mount('#app')

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
