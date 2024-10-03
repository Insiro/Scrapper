import { CSSProperties, PropsWithChildren } from "react";
import Pin from "@/shared/assets/pin.png";
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

interface CardProps {
    style?: CSSProperties;
    pin?: boolean;
}

// Card 컴포넌트 정의
export const Card: React.FC<PropsWithChildren<CardProps>> = ({ children, style, pin, ...props }) => {
    return (
        <div style={{ ...cardStyle, ...(style ?? {}) }} {...props}>
            {pin && <img style={{ height: "1rem", float: "right", left: "1rem", top: "1rem" }} src={Pin} />}
            {children}
        </div>
    );
};
