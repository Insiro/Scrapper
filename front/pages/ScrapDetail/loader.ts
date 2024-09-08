import { LoaderFunction } from "react-router-dom";
import { getScrap } from "../../services/scrapService";

// Scrap 상세 데이터를 로드하는 loader 함수 정의
export const scrapDetailLoader: LoaderFunction = async ({ params }) => {
    return await getScrap(params.scrapId)
}