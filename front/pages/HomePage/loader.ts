import { LoaderFunction } from "react-router-dom";
import { scrapApi } from "@/entities/scrap";

// Scrap 데이터를 로드하는 loader 함수 정의
export const homePageLoader: LoaderFunction = async () => {
    return await scrapApi.getScrapList({ pined: true });
};
