import React from "react";
import { useLoaderData } from "react-router-dom";
import ScrapForm from "../../components/ScrapForm";
import ScrapList from "../../components/ScrapList";
import { Scrap } from "../../types";

const HomePage: React.FC = () => {
    const scraps = useLoaderData() as Scrap[];

    const handleScrapAdd = (newScrap: Scrap) => {
        // 상태 관리를 통해 새로운 스크랩을 목록에 추가할 수 있습니다.
        console.log(newScrap);
    };

    return (
        <div style={{ padding: "2rem" }}>
            <h1>Scrap List</h1>
            <ScrapForm onScrapAdd={handleScrapAdd} />
            <ScrapList scraps={scraps} />
        </div>
    );
};

export default HomePage;
