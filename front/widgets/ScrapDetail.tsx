import React, { useState, useEffect, CSSProperties, FC } from "react";
import { Card } from "./Common/Card";
import { Button } from "./Common/Button";
import { scrapApi } from "@/entities/scrap";
import { api as imageApi } from "@/entities/scrapImage";
import { Scrap } from "@/entities/scrap";
import { color } from "@/shared/constant";
import ConfirmModal from "./Common/Confirm";
import Config from "@/shared/config";

const ContentLine: FC<{ title: string; content: string | undefined }> = ({ title, content }) => {
    return (
        <>
            <strong style={{ fontWeight: "bold", color: "#333", fontSize: "1.1rem", marginRight: "0.5rem" }}>{title}</strong>
            <div style={{ whiteSpace: "pre-line", color: "#555", lineHeight: "1.5" }}>
                {content?.split("\n").map((line, index) => (
                    <React.Fragment key={index}>
                        {line}
                        <br />
                    </React.Fragment>
                ))}
            </div>
        </>
    );
};

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
    contentContainer: {
        marginBottom: "1rem",
        fontSize: "1rem",
        display: "grid",
        gridTemplateColumns: "max-content 1fr",
        gap: "1rem",
        alignItems: "start",
    },
};

interface ScrapDetailProps {
    scrap: Scrap;
    refreshScrap: () => void | Promise<void>;
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
            const result = refreshScrap();
            if (result instanceof Promise) await result;
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
            const result = refreshScrap();

            if (result instanceof Promise) await result;
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
                <div style={styles.contentContainer}>
                    <ContentLine title="Url" content={scrap.url} />
                    <ContentLine title="Author" content={`${scrap.author_name} (@${scrap.author_tag})`} />
                    <ContentLine title="Content" content={scrap.content} />
                    {scrap.comment && <ContentLine title="Comment" content={scrap.comment} />}
                </div>
            </Card>
            <Card style={{ marginTop: "1rem" }}>
                <h2>Images</h2>

                <div style={isSmallScreen ? { ...styles.imageContainer, gridTemplateColumns: "repeat(1, 1fr)" } : styles.imageContainer}>
                    {scrap.images.map((image) => (
                        <div key={image.id} style={styles.imageWrapper}>
                            <input type="checkbox" style={styles.checkbox} checked={selectedImages.has(image.id)} onChange={() => handleImageSelect(image.id)} />
                            <img src={`${Config.hostPath}media/${image.file_name}`} alt={image.file_name} style={styles.image} />
                        </div>
                    ))}
                </div>

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
