import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react-swc";
import tsconfigPaths from "vite-tsconfig-paths";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd(), "");
    let basePath = env.SCRAPER_BASE_PATH ?? "";
    if (basePath[0] === "/") basePath = basePath.slice(1);
    return {
        base: basePath,
        envPrefix: ["SCRAPER"],
        plugins: [react(), tsconfigPaths()],
    };
});
