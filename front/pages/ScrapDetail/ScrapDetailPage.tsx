import { useEffect, useMemo, useState } from "react";
import { useLoaderData, useLocation, useNavigate } from "react-router-dom";

import ScrapDetail from "@/widgets/ScrapDetail";
import { useTitleContext } from "@/entities/title";
import { Scrap, scrapApi } from "@/entities/scrap";

export const ScrapDetailPage: React.FC = () => {
    const initialScrap = useLoaderData() as Scrap; // useLoaderData로 로드된 데이터를 가져옵니다.
    const location = useLocation();
    const scrapId = useMemo(() => {
        location.pathname.split("/").at(-1);
    }, [location]);
    const [scrap, setScrap] = useState(initialScrap);
    const { setPageTitle } = useTitleContext();
    const navigate = useNavigate();
    const refreshScrap = async () => {
        if (scrapId === undefined) return;
        const sc = await scrapApi.getScrap(parseInt(scrapId));
        if (sc === null) return navigate("404");
        setScrap(sc);
    };
    useEffect(() => setPageTitle("Scrap Detail"), [setPageTitle]); // 페이지가 로드될 때 제목 설정

    return <ScrapDetail scrap={scrap} refreshScrap={refreshScrap} />;
};
