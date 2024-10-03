import React, { useEffect, useState } from "react";
import { Link, useLoaderData } from "react-router-dom";

import ScrapForm from "@/widgets/ScrapForm";
import ScrapList from "@/widgets/ScrapList";
import { useTitleContext } from "@/entities/title";
import { Scrap } from "@/entities/scrap";
import { Card } from "@/widgets/Common/Card";

export const HomePage: React.FC = () => {
    const { setPageTitle } = useTitleContext(); // LayoutContext에서 상태 업데이트 함수 가져오기

    useEffect(() => setPageTitle("Scrap Home"), [setPageTitle]); // 페이지가 로드될 때 제목 설정

    // 스크랩 목록 상태 관리
    const preloadedScraps = useLoaderData() as { list: Scrap[]; count: number };
    const [scraps, setScraps] = useState(preloadedScraps);

    const handleScrapAdd = (newScrap: Scrap) => {
        scraps.list = [newScrap, ...scraps.list];
        scraps.count += 1;
        setScraps(scraps);
    };

    return (
        <div style={{ textAlign: "center" }}>
            <ScrapForm onScrapAdd={handleScrapAdd} />
            <ScrapList scraps={scraps.list} />
            <Card>
                <Link style={{ color: "black", textDecoration: "none" }} to="scraps">
                    <strong style={{ fontSize: "1.2rem", padding: "2rem" }}>more</strong>
                </Link>
            </Card>
        </div>
    );
};

export default HomePage;
