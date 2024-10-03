import React, { useEffect, useState } from "react";
import { useLoaderData, useSearchParams } from "react-router-dom";

import ScrapForm from "@/widgets/ScrapForm";
import ScrapList from "@/widgets/ScrapList";
import { useTitleContext } from "@/entities/title";
import { Scrap, scrapApi } from "@/entities/scrap";
import { PageNation } from "@/widgets/PageNation";

export const ScrapListPage: React.FC = () => {
    const { setPageTitle } = useTitleContext(); // LayoutContext에서 상태 업데이트 함수 가져오기
    const [searchParams, setSearchParams] = useSearchParams();

    useEffect(() => setPageTitle("Scrap List"), [setPageTitle]); // 페이지가 로드될 때 제목 설정

    // 스크랩 목록 상태 관리
    const preloadedScraps = useLoaderData() as { list: Scrap[]; count: number };
    const [scraps, setScraps] = useState(preloadedScraps);
    const [page, setPage] = useState(Number(searchParams.get("page") ?? 1));

    useEffect(() => {
        scrapApi.getScrapList({ page }).then((it) => {
            setScraps(it);
            searchParams.set("page", String(page));
            setSearchParams(searchParams);
        });
    }, [page, searchParams, setSearchParams]);

    const handleScrapAdd = (newScrap: Scrap) => {
        scraps.list = [newScrap, ...scraps.list];
        scraps.count += 1;
        setScraps(scraps);
    };

    return (
        <div style={{ textAlign: "center" }}>
            <ScrapForm onScrapAdd={handleScrapAdd} />
            <ScrapList scraps={scraps.list} />
            <PageNation totalPage={Math.floor(scraps.count / 10) + 1} visiblePage={10} current={page} setPage={setPage} />
        </div>
    );
};

export default ScrapListPage;
