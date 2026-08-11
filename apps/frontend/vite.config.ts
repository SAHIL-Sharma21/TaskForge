import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  define: {
    "process.env": process.env,
  },
  server: {
    port: 3000,
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@taskForge/openapi": path.resolve(__dirname, "../../packages/openapi/src"),
      "@taskForge/zod": path.resolve(__dirname, "../../packages/zod/src"),
    },
  },
});
