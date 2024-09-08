import React from "react";
import { Link } from "react-router-dom";
import { Scrap } from "../types/Scrap";

interface ScrapListProps {
    scraps: Scrap[];
}

const ScrapList: React.FC<ScrapListProps> = ({ scraps }) => {
    return (
        <ul style={{ listStyleType: "none", padding: 0 }}>
            {scraps.map((scrap) => (
                <li key={scrap.id} style={{ marginBottom: "1rem" }}>
                    <h2>
                        {scrap.author_name} (@{scrap.author_tag})
                    </h2>
                    <p>{scrap.content}</p>
                    <Link to={`/scraps/${scrap.id}`} style={{ color: "blue", textDecoration: "underline" }}>
                        View Details
                    </Link>
                </li>
            ))}
        </ul>
    );
};

export default ScrapList;
