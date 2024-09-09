const Config = {
    basePath: import.meta.env.SCRAPER_BASE_PATH || "/", // 환경변수에서 basePath 가져오기 (기본값은 "/")
    apiHost: import.meta.env.SCRAPER_API_HOST || "http://localhost", // 환경변수에서 apiHost 가져오기 (기본값은 로컬 호스트)
    apiPort: import.meta.env.SCRAPPER_API_PORT || 8000,

    get hostPath() {
        return `${this.apiHost}:${this.apiPort}${this.basePath}`;
    },
    get api() {
        // basePath와 apiHost를 결합하여 api 경로 구성
        return `${this.apiHost}:${this.apiPort}${this.basePath}api`;
    }
};

export default Config