import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'
import fullImportPlugin from './vitePlugin/fullImport/fullImport.js'
import GvaPosition from './vitePlugin/gvaPosition/index.js'
// import WindiCSS from 'vite-plugin-windicss'

// https://vitejs.dev/config/
export default defineConfig(({ common, mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  const config = {
    plugins: [
      vue(),
      GvaPosition(),
      // WindiCSS()
    ],
    build: {
      chunkSizeWarningLimit: 2000,
    },
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: `@use "@/style/element/index.scss" as *;`,
        }
      }
    },
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    server: { // 服务相关配置
      host: '127.0.0.1',
      port: env.VITE_CLI_PORT,
      open: `http://127.0.0.1:${env.VITE_CLI_PORT}`,
      proxy: { // 请求代理
        '/api': {
          target: `http://${env.VITE_SERVER_URL}:${env.VITE_SERVER_PORT}`,	// 实际请求地址
          rewrite: (path) => path.replace(/^\/api/, '')
        },
      }
    }
  }
  config.plugins.push(fullImportPlugin())
  return config
})
