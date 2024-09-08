import React, { useState, useEffect, CSSProperties } from "react";
import { Card } from "./Common/Card";
import { Button } from "./Common/Button";
import { reScrap } from "../entities/scrap/lib/scrapService";
import { deleteImage } from "../entities/scrapImage/lib/api";
import { Scrap } from "../entities/scrap/Scrap";

// ScrapDetail 스타일 객체
const styles: Record<string, CSSProperties> = {
    title: {
        fontSize: "1.8rem",
        fontWeight: "bold",
        marginBottom: "1rem",
    },
    detail: {
        marginBottom: "0.5rem",
        fontSize: "1rem",
    },
    imageContainer: {
        display: "grid",
        gridTemplateColumns: "repeat(2, 1fr)", // 기본적으로 두 개의 열로 구성
        gap: "1rem",
        marginTop: "1rem",
    },
    imageWrapper: {
        position: "relative",
    },
    checkbox: {
        position: "absolute",
        top: "8px",
        left: "8px",
        zIndex: 1,
    },
    image: {
        width: "100%",
        height: "auto",
        borderRadius: "4px",
        objectFit: "cover",
    },
};

interface ScrapDetailProps {
    scrap: Scrap;
}

const ScrapDetail: React.FC<ScrapDetailProps> = ({ scrap }) => {
    const [selectedImages, setSelectedImages] = useState<Set<number>>(new Set());
    const [isReScraping, setIsReScraping] = useState<boolean>(false);
    const [isSmallScreen, setIsSmallScreen] = useState<boolean>(window.innerWidth <= 600);

    useEffect(() => {
        // 화면 크기 변경 감지
        const handleResize = () => setIsSmallScreen(window.innerWidth <= 600);

        window.addEventListener("resize", handleResize);
        // 컴포넌트 언마운트 시 이벤트 리스너 제거
        return () => window.removeEventListener("resize", handleResize);
    }, []);

    const handleImageSelect = (imageId: number) => {
        setSelectedImages((prevSelectedImages) => {
            const newSelectedImages = new Set(prevSelectedImages);
            if (newSelectedImages.has(imageId)) newSelectedImages.delete(imageId);
            else newSelectedImages.add(imageId);

            return newSelectedImages;
        });
    };

    const handleDeleteSelectedImages = async () => {
        if (selectedImages.size === 0) {
            alert("Please select images to delete.");
            return;
        }

        try {
            await deleteImage(Array.from(selectedImages));
            alert("Selected images have been deleted.");
            setSelectedImages(new Set());
            // 이미지 삭제 후 로직 추가 (예: 목록 갱신)
        } catch (error) {
            alert("Failed to delete selected images.");
            console.error("Error deleting images:", error);
        }
    };

    const handleReScrap = async () => {
        setIsReScraping(true);
        try {
            await reScrap(scrap.id);
            alert("Scrap has been successfully updated.");
            // 여기서 서버로부터 새롭게 갱신된 스크랩 데이터를 가져와 업데이트할 수 있음
        } catch (error) {
            alert("Failed to re_scrap the URL.");
            console.error("Error re_scraping:", error);
        } finally {
            setIsReScraping(false);
        }
    };

    return (
        <>
            <Card>
                <h2>Content</h2>
                <p style={styles.detail}>
                    <strong>URL:</strong> {scrap.url}
                </p>
                <p style={styles.detail}>
                    <strong>Author:</strong> {scrap.author_name} (@{scrap.author_tag})
                </p>
                <p style={styles.detail}>
                    <strong>Content:</strong> {scrap.content}
                </p>
                {scrap.comment && (
                    <p style={styles.detail}>
                        <strong>Comment:</strong> {scrap.comment}
                    </p>
                )}
            </Card>
            <Card style={{ marginTop: "1rem" }}>
                <h2>Images</h2>
                {scrap.images.length > 0 ? (
                    <div style={isSmallScreen ? { ...styles.imageContainer, gridTemplateColumns: "repeat(1, 1fr)" } : styles.imageContainer}>
                        {scrap.images.map((image) => (
                            <div key={image.id} style={styles.imageWrapper}>
                                <input type="checkbox" style={styles.checkbox} checked={selectedImages.has(image.id)} onChange={() => handleImageSelect(image.id)} />
                                <img src={`http://localhost:8000/media/${image.file_name}`} alt={image.file_name} style={styles.image} />
                            </div>
                        ))}
                    </div>
                ) : (
                    <p>No images available</p>
                )}
                <Button backgroundColor="#d9534f" onClick={handleDeleteSelectedImages}>
                    Delete Selected Images
                </Button>
                <Button backgroundColor="#007bff" onClick={handleReScrap} disabled={isReScraping}>
                    {isReScraping ? "Re_scraping..." : "Re_scrap"}
                </Button>
            </Card>
        </>
    );
};

export default ScrapDetail;
