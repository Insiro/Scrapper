import React, { useState } from "react";
import { Scrap } from "@/entities/scrap";
import { Button } from "./Common/Button";
import { scrapApi } from "@/entities/scrap";
import { Card } from "./Common/Card";
import { color } from "@/shared/constant";

interface ScrapFormProps {
    onScrapAdd: (scrap: Scrap) => void;
}

const ScrapForm: React.FC<ScrapFormProps> = ({ onScrapAdd }) => {
    const [url, setUrl] = useState<string>("");
    const [error, setError] = useState<string | null>(null);

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            const scrap = await scrapApi.createScrap(url);
            onScrapAdd(scrap);
            setUrl("");
        } catch {
            setError("Failed to add scrap.");
        }
    };

    // 스타일 정의
    const formStyle = {
        display: "flex",
        alignItems: "center", // 입력 필드와 버튼을 수직 가운데 정렬
        gap: "1rem", // 입력 필드와 버튼 사이의 간격 추가
    };
    const inputStyle = {
        marginTop: "0.5rem",
        padding: "0.5rem",
        width: "300px",
        borderRadius: "4px",
        border: "1px solid #ddd",
        flexGrow: 1, // 입력 필드가 가능한 넓게 확장되도록 설정
        height: "20px", // 버튼과 동일한 높이 설정
    };

    const errorStyle = {
        color: "red",
        marginBottom: "1rem",
    };

    return (
        <Card style={formStyle}>
            {error && <p style={errorStyle}>{error}</p>}
            <input type="text" value={url} onChange={(e) => setUrl(e.target.value)} placeholder="Enter URL to scrap" style={inputStyle} />
            <Button backgroundColor={color.blue} onClick={handleSubmit} style={{ height: "40px" }}>
                Scrap URL
            </Button>
        </Card>
    );
};

export default ScrapForm;
