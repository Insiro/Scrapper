import { LoaderFunction } from "react-router-dom";
import { scrapApi } from "@/entities/scrap";

// Scrap 데이터를 로드하는 loader 함수 정의
export const scrapListLoader: LoaderFunction = async ({ request }) => {
    const url = new URL(request.url)
    const page = url.searchParams.get("page") ?? 1
    return await scrapApi.getScrapList(Number(page))
};