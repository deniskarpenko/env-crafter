import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  publicDir: 'public',  // Указываем папку для статических файлов
  resolve: {
    alias: {
      '@': '/src'
    }
  }
})
