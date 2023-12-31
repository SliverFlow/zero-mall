import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'

// https://vitejs.dev/config/
export default defineConfig((common, mode) => {
  const env = loadEnv(mode, process.cwd(), '')
  const config = {
    plugins: [vue()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    server: { // 服务相关配置
      host: '127.0.0.1',
      port: env.VITE_CLI_PORT,
      open: `http://127.0.0.1:${env.VITE_CLI_PORT}/#/index`,
      proxy: { // 请求代理
        '/api': {
          target: `http://127.0.0.1:20011/`,	// 实际请求地址
          // rewrite: (path) => path.replace(/^\/api/, '')
        },
      }
    }
  }

  return config
})
