const basePath = import.meta.env.SCRAPER_BASE_PATH || "/";
const apiHost = import.meta.env.SCRAPER_API_HOST || "";
const apiPort = import.meta.env.SCRAPPER_API_PORT || 8000;
const hostPath = new URL(basePath, `${apiHost}:${apiPort}`).href

const Config = {
    basePath,
    apiHost,
    apiPort,
    hostPath,
    api: hostPath + "/api"
};
console.log(Config)
export default Config