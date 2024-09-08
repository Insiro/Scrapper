import React, { useState, useEffect, CSSProperties } from "react";
import { Card } from "./Common/Card";
import { Button } from "./Common/Button";
import { scrapApi } from "@/entities/scrap";
import { api as imageApi } from "@/entities/scrapImage";
import { Scrap } from "@/entities/scrap";
import { color } from "@/shared/constant";
import ConfirmModal from "./Common/Confirm";

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
    refreshScrap: () => void;
}

const ScrapDetail: React.FC<ScrapDetailProps> = ({ scrap, refreshScrap }) => {
    const [selectedImages, setSelectedImages] = useState<Set<number>>(new Set());
    const [isReScraping, setIsReScraping] = useState<boolean>(false);
    const [isSmallScreen, setIsSmallScreen] = useState<boolean>(window.innerWidth <= 600);
    const [isDelModalOpen, setIsDelModalOpen] = useState<boolean>(false);
    const [isReScrapModalOpen, setIsReScrapModalOpen] = useState<boolean>(false);

    useEffect(() => {
        // 화면 크기 변경 감지
        const handleResize = () => setIsSmallScreen(window.innerWidth <= 600);
        window.addEventListener("resize", handleResize);

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

    const actionDeleteImage = async (confirm: boolean) => {
        if (!confirm) {
            setIsDelModalOpen(false);
            setSelectedImages(new Set());
            return;
        }
        try {
            await imageApi.deleteImage(Array.from(selectedImages));
            alert("Selected images have been deleted.");
            setSelectedImages(new Set());
            refreshScrap();
        } catch (error) {
            alert("Failed to delete selected images.");
            console.error("Error deleting images:", error);
        }
        setIsDelModalOpen(false);
    };

    const handleDeleteSelectedImages = () => setIsDelModalOpen(true);

    const handleReScrap = () => setIsReScrapModalOpen(true);

    const actionReScrap = async (confirm: boolean) => {
        if (!confirm) {
            setIsReScrapModalOpen(false);
            return;
        }
        setIsReScraping(true);
        try {
            await scrapApi.reScrap(scrap.id);
            alert("Scrap has been successfully updated.");
            refreshScrap();
        } catch (error) {
            alert("Failed to re:scrap the URL.");
            console.error("Error re:scraping:", error);
        }
        setIsReScraping(false);
        setIsReScrapModalOpen(false);
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
                <hr />
                <Button backgroundColor={color.red} onClick={handleDeleteSelectedImages}>
                    Delete Selected Images
                </Button>
                <Button backgroundColor={color.blue} onClick={handleReScrap} disabled={isReScraping}>
                    {isReScraping ? "Re:Scraping..." : "Re:Scrap"}
                </Button>
                <ConfirmModal isOpen={isReScrapModalOpen} title="Confirm Re:Scrap" message="Are you sure re:scrap this?" action={actionReScrap} />
                <ConfirmModal isOpen={isDelModalOpen} title="Confirm Delete" message="Are you sure delete the selected images?" action={actionDeleteImage} />
            </Card>
        </>
    );
};

export default ScrapDetail;
