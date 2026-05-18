import { resolve } from 'node:path'
import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'

const rootDir = fileURLToPath(new URL('.', import.meta.url))

export default defineConfig({
  build: {
    emptyOutDir: true,
    rollupOptions: {
      input: {
        popup: resolve(rootDir, 'popup.html'),
        background: resolve(rootDir, 'src/background.ts'),
        content: resolve(rootDir, 'src/content.ts'),
      },
      output: {
        entryFileNames: 'assets/[name].js',
        chunkFileNames: 'assets/[name].js',
        assetFileNames: 'assets/[name][extname]',
      },
    },
  },
})
