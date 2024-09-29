const basePath = import.meta.env.SCRAPER_BASE_PATH || "/";
const apiHost = import.meta.env.SCRAPER_API_HOST || "";

let hostPath: string
if (apiHost) {
    const apiPort = import.meta.env.SCRAPPER_API_PORT || 80;
    hostPath = new URL(basePath, `${apiHost}:${apiPort}`).href
} else hostPath = basePath



const Config = {
    basePath,
    apiHost,
    hostPath,
};
console.log(Config)
export default Config