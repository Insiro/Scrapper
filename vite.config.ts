import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tsconfigPaths from "vite-tsconfig-paths";

let basePath = process.env.SCRAPER_BASE_PATH ?? ""
if (basePath[0] === "/") basePath = basePath.slice(1)

// https://vitejs.dev/config/
export default defineConfig({
  base: basePath,
  envPrefix: ["SCRAPER"],
  plugins: [react(), tsconfigPaths()],
})
