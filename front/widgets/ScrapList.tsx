import React from "react";
import { Link } from "react-router-dom";
import { Scrap } from "@/entities/scrap";
import { color } from "@/shared/constant";

interface ScrapListProps {
    scraps: Scrap[];
}

const ScrapList: React.FC<ScrapListProps> = ({ scraps }) => {
    const listStyle = {
        listStyleType: "none",
        padding: 0,
    };

    const listItemStyle = {
        marginBottom: "1rem",
        padding: "1rem",
        border: "1px solid #ddd",
        borderRadius: "8px",
        backgroundColor: "#fff",
    };

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
        <ul style={listStyle}>
            {scraps.map((scrap) => (
                <li key={scrap.id} style={listItemStyle}>
                    <div style={authorStyle}>
                        {scrap.author_name} (@{scrap.author_tag})
                    </div>
                    <p style={contentStyle}>{scrap.content}</p>
                    <Link to={`/scraps/${scrap.id}`} style={linkStyle}>
                        View Details
                    </Link>
                </li>
            ))}
        </ul>
    );
};

export default ScrapList;
