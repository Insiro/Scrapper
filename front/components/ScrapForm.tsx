import React, { useState } from "react";
import axios from "axios";
import { Scrap } from "../types/Scrap";

interface ScrapFormProps {
    onScrapAdd: (scrap: Scrap) => void;
}

const ScrapForm: React.FC<ScrapFormProps> = ({ onScrapAdd }) => {
    const [url, setUrl] = useState<string>("");
    const [force, setForce] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        try {
            const response = await axios.post<Scrap>("/api/scraps", {
                url,
                force,
            });
            onScrapAdd(response.data);
            setUrl("");
            setForce(false);
        } catch {
            setError("Failed to add scrap.");
        }
    };

    return (
        <form onSubmit={handleSubmit} style={{ marginBottom: "2rem" }}>
            {error && <p style={{ color: "red" }}>{error}</p>}
            <div style={{ marginBottom: "1rem" }}>
                <input type="text" value={url} onChange={(e) => setUrl(e.target.value)} placeholder="Enter URL to scrap" style={{ padding: "0.5rem", width: "300px" }} />
            </div>
            <div style={{ marginBottom: "1rem" }}>
                <label>
                    <input type="checkbox" checked={force} onChange={(e) => setForce(e.target.checked)} />
                    Force update
                </label>
            </div>
            <button type="submit" style={{ padding: "0.5rem 1rem" }}>
                Scrap URL
            </button>
        </form>
    );
};

export default ScrapForm;
