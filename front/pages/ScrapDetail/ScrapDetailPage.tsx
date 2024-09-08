import { useEffect } from "react";
import { useLoaderData, useOutletContext } from "react-router-dom";

import ScrapDetail from "../../widgets/ScrapDetail";
import { LayoutContextType } from "../../entities/title/lib/layoutContext";
import { Scrap } from "../../entities/scrap/Scrap";

const ScrapDetailPage: React.FC = () => {
    const scrap = useLoaderData() as Scrap; // useLoaderData로 로드된 데이터를 가져옵니다.
    const { setPageTitle } = useOutletContext<LayoutContextType>();

    useEffect(() => setPageTitle("Scrap Detail"), [setPageTitle]); // 페이지가 로드될 때 제목 설정

    return <ScrapDetail scrap={scrap} />;
};
export default ScrapDetailPage;
