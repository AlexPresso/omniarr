import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig(() => {
    return {
        plugins: [
            tailwindcss(),
            react()
        ],
        server: {
            proxy: {
                '/api': {
                    target: `http://localhost:8080/`,
                    changeOrigin: true,
                }
            }
        }
    }
})
