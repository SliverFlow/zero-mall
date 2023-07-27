import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'
import fullImportPlugin from './vitePlugin/fullImport/fullImport.js'
import GvaPosition from './vitePlugin/gvaPosition/index.js'
import WindiCSS from 'vite-plugin-windicss'

// https://vitejs.dev/config/
export default defineConfig(() => {
  const config = {
    plugins: [
      vue(),
      GvaPosition(),
      WindiCSS()
    ],
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
    }
  }
  config.plugins.push(fullImportPlugin())
  return config
})
