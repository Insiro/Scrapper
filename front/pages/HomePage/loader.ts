import { LoaderFunction } from "react-router-dom";
import { getScrapList } from "../../entities/scrap/lib/scrapService";

// Scrap 데이터를 로드하는 loader 함수 정의
export const homePageLoader: LoaderFunction = async () => {
    return await getScrapList()
};