import React from "react";
import { Link } from "react-router-dom";

const NotFoundPage: React.FC = () => {
    return (
        <div style={{ textAlign: "center", padding: "2rem" }}>
            <h1>404 - Not Found</h1>
            <p>The page you are looking for does not exist.</p>
            <Link to="/" style={{ color: "blue", textDecoration: "underline" }}>
                Go to Home
            </Link>
        </div>
    );
};

export default NotFoundPage;
