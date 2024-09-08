import { useLoaderData } from "react-router-dom";
import ScrapDetail from "../../components/ScrapDetail";
import { Scrap } from "../../types";

const ScrapDetailPage: React.FC = () => {
    const scrap = useLoaderData() as Scrap; // useLoaderData로 로드된 데이터를 가져옵니다.
    console.log(scrap);

    return (
        <div style={{ padding: "2rem" }}>
            <ScrapDetail scrap={scrap} /> {/* ScrapDetail 컴포넌트에 로드된 데이터를 전달 */}
        </div>
    );
};
export default ScrapDetailPage;
