import { LoaderFunction } from "react-router-dom";
import { scrapApi } from "@/entities/scrap";

// Scrap 상세 데이터를 로드하는 loader 함수 정의
export const scrapDetailLoader: LoaderFunction = async ({ params }) => {
    return await scrapApi.getScrap(params.scrapId);
};
