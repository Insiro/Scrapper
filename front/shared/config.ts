let basePath = import.meta.env.SCRAPER_BASE_PATH || "";
if (basePath === "/") basePath = "";
else if (basePath[0] === "/") basePath = basePath.slice(1);

const { hostname, protocol } = window.location;

const apiHost = import.meta.env.SCRAPER_API_HOST || protocol + "//" + hostname;

let hostPath: string;
if (apiHost) {
    hostPath = basePath ? new URL(basePath, apiHost).href : apiHost;
} else hostPath = basePath;

const Config = {
    basePath,
    apiHost,
    hostPath,
};

export default Config;
