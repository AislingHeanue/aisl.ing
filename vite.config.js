import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import {resolve} from 'path'

const root = resolve(__dirname, "src")
const outDir = resolve(__dirname,"dist")
const publicDir = resolve(__dirname,"public")

// https://vitejs.dev/config/
export default defineConfig({
  root,
  publicDir,
  plugins: [
    react(),
  ],
  build: {
    outDir,
    emptyOutDir: true,
    rollupOptions: {
      input : {
        main: resolve(root,"index.html"),
        demo: resolve(root,"collatz","index.html")
      }
    }
  }
})
