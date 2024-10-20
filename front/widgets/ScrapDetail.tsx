import React, { useState, CSSProperties, FC, PropsWithChildren } from "react";
import { Card } from "./Common/Card";
import { Button } from "./Common/Button";
import ConfirmModal from "./Common/Confirm";

import { scrapApi } from "@/entities/scrap";
import { api as imageApi } from "@/entities/scrapImage";
import { Scrap } from "@/entities/scrap";
import { color } from "@/shared/constant";
import Config from "@/shared/config";
import { useIsSmallScreen } from "@/shared/smallScreen/useIsSmallSize";

type inlineProps = { title: string; content: string; children: undefined };
type ContentLineProps = inlineProps | PropsWithChildren<{ title: string; content?: string }>;

const ContentLine: FC<ContentLineProps> = ({ title, content, children }) => {
    const titleStyle = { fontWeight: "bold", color: "#333", fontSize: "1.1rem", marginRight: "0.5rem" };
    return (
        <>
            <strong style={titleStyle}>{title}</strong>
            <div style={{ whiteSpace: "pre-line", color: "#555", lineHeight: "1.5" }}>
                {content?.split("\n").map((line, index) => (
                    <React.Fragment key={index}>
                        {line}
                        <br />
                    </React.Fragment>
                ))}
                {children}
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
    const [isDelModalOpen, setIsDelModalOpen] = useState<boolean>(false);
    const [isReScrapModalOpen, setIsReScrapModalOpen] = useState<boolean>(false);

    const isSmallScreen = useIsSmallScreen();

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
            <Card pin={scrap.pin}>
                <h2>Content</h2>
                <div
                    style={
                        isSmallScreen
                            ? { ...styles.contentContainer, display: "block" }
                            : { ...styles.contentContainer }
                    }
                >
                    <ContentLine title="Source">
                        <a style={{ wordBreak: "break-all" }} href={scrap.url} target="_blank">
                            {scrap.source} : {scrap.url}
                        </a>
                    </ContentLine>
                    <ContentLine title="Author" content={`${scrap.author_name} (@${scrap.author_tag})`} />
                    <ContentLine title="Tag">
                        {scrap.tags.map((it) => (
                            <span style={{ margin: "0.5rem" }}>#{it}</span>
                        ))}
                    </ContentLine>
                    <ContentLine title="Content" content={scrap.content} />

                    {scrap.comment && <ContentLine title="Comment" content={scrap.comment} />}
                </div>
            </Card>
            <Card style={{ marginTop: "1rem" }}>
                <h2>Images</h2>

                <div
                    style={
                        isSmallScreen
                            ? { ...styles.imageContainer, gridTemplateColumns: "repeat(1, 1fr)" }
                            : styles.imageContainer
                    }
                >
                    {scrap.images.map((image) => (
                        <div key={image.id} style={styles.imageWrapper}>
                            <input
                                type="checkbox"
                                style={styles.checkbox}
                                checked={selectedImages.has(image.id)}
                                onChange={() => handleImageSelect(image.id)}
                            />
                            <img
                                src={`${Config.hostPath}/media/${image.file_name}`}
                                alt={image.file_name}
                                style={styles.image}
                            />
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
                <ConfirmModal
                    isOpen={isReScrapModalOpen}
                    title="Confirm Re:Scrap"
                    message="Are you sure re:scrap this?"
                    action={actionReScrap}
                    disabled={isReScraping}
                />
                <ConfirmModal
                    isOpen={isDelModalOpen}
                    title="Confirm Delete"
                    message="Are you sure delete the selected images?"
                    action={actionDeleteImage}
                />
            </Card>
        </>
    );
};

export default ScrapDetail;
