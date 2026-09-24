import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/recipes': 'http://localhost:8080',
      '/ingredients': 'http://localhost:8080',
    }
  },
  build: {
    outDir: '../backend/dist',
    emptyOutDir: true 
  }
})