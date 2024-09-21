import React from "react";
import { Link } from "react-router-dom";
import { Scrap } from "@/entities/scrap";
import { color } from "@/shared/constant";
import { Card } from "./Common/Card";

interface ScrapListProps {
    scraps: Scrap[];
}

const ScrapList: React.FC<ScrapListProps> = ({ scraps }) => {
    const linkStyle = {
        color: color.blue,
        textDecoration: "underline",
    };

    const authorStyle = {
        fontSize: "1.2rem",
        fontWeight: "bold",
    };

    const contentStyle = {
        fontSize: "1rem",
        marginBottom: "0.5rem",
    };

    return (
        <>
            {scraps.map((scrap) => (
                <Card key={scrap.id}>
                    <span style={authorStyle}>
                        {scrap.author_name} (@{scrap.author_tag}){" "}
                    </span>
                    <small style={{ margin: "0.3rem" }}>{scrap.source}</small>

                    <p style={contentStyle}>{scrap.content}</p>
                    <Link to={`/scraps/${scrap.id}`} style={linkStyle}>
                        View Details
                    </Link>
                </Card>
            ))}
        </>
    );
};

export default ScrapList;
