let basePath = import.meta.env.SCRAPER_BASE_PATH || "";
if (basePath === "/") basePath = ""
const apiHost = import.meta.env.SCRAPER_API_HOST || "";

let hostPath: string


if (apiHost) {
    const apiPort = import.meta.env.SCRAPER_API_PORT || 80;
    console.log(import.meta.env.SCRAPER_API_PORT)
    const apiPath = `${apiHost}:${apiPort}`
    hostPath = basePath ? (new URL(basePath, apiPath)).href : apiPath
} else hostPath = basePath

const Config = {
    basePath,
    apiHost,
    hostPath,
};
console.log(Config)
console.log(import.meta.env)
export default Config