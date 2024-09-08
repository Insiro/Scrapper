import React from 'react';
import { Scrap } from '../types';

interface ScrapDetailProps {
    scrap: Scrap;
}

const ScrapDetail: React.FC<ScrapDetailProps> = ({ scrap }) => {
    return (
        <div>
            <h1>Scrap Details</h1>
            <p><strong>URL:</strong> {scrap.url}</p>
            <p><strong>Author:</strong> {scrap.author_name} (@{scrap.author_tag})</p>
            <p><strong>Content:</strong> {scrap.content}</p>
            {scrap.comment && <p><strong>Comment:</strong> {scrap.comment}</p>}
            <div>
                <h2>Images</h2>
                {scrap.image_names.length > 0 ? (
                    <ul>
                        {scrap.image_names.map((imageName) => (
                            <li key={imageName}>
                                <img 
                                    src={`http://localhost:8000/media/${imageName}`}  // 서버에서 서빙되는 이미지 경로
                                    alt={imageName}
                                    style={{ maxWidth: '100%' }} 
                                />
                            </li>
                        ))}
                    </ul>
                ) : (
                    <p>No images available</p>
                )}
            </div>
        </div>
    );
};

export default ScrapDetail;