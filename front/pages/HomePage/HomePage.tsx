import React, { useEffect, useState } from "react";
import { useLoaderData } from "react-router-dom";

import ScrapForm from "../../widgets/ScrapForm";
import ScrapList from "../../widgets/ScrapList";
import { useLayoutContext } from "../../entities/title/lib/useLayoutContext";
import { Scrap } from "../../entities/scrap/Scrap";

const HomePage: React.FC = () => {
    const { setPageTitle } = useLayoutContext(); // LayoutContext에서 상태 업데이트 함수 가져오기

    useEffect(() => {
        setPageTitle("Scrap List");
    }, [setPageTitle]); // 페이지가 로드될 때 제목 설정

    // 초기 스크랩 데이터를 로드
    const initialScraps = useLoaderData() as Scrap[];

    // 스크랩 목록 상태 관리
    const [scraps, setScraps] = useState<Scrap[]>(initialScraps);

    const handleScrapAdd = (newScrap: Scrap) => setScraps((prevScraps) => [newScrap, ...prevScraps]);

    return (
        <div>
            <ScrapForm onScrapAdd={handleScrapAdd} />
            <ScrapList scraps={scraps} />
        </div>
    );
};

export default HomePage;
