import { CSSProperties } from "react";

// Card 스타일 객체
const cardStyle: CSSProperties = {
    marginTop: "1rem",
    padding: "1.5rem",
    border: "1px solid #ddd",
    borderRadius: "8px",
    backgroundColor: "#fff",
    maxWidth: "800px",
    marginLeft: "auto",
    marginRight: "auto",
};

// Card 컴포넌트 정의
export const Card: React.FC<{ children: React.ReactNode; style?: CSSProperties }> = ({ children, style, ...props }) => {
    return (
        <div style={{ ...cardStyle, ...(style ?? {}) }} {...props}>
            {children}
        </div>
    );
};
