import { CSSProperties, FC, PropsWithChildren } from "react";

const ButtonStyle = {
    marginTop: "0.5rem",
    padding: "0.5rem 1rem",
    color: "#fff",
    border: "none",
    borderRadius: "4px",
    cursor: "pointer",
    fontSize: "1rem",
};

interface BtnProps extends PropsWithChildren, React.ComponentPropsWithRef<"button"> {
    backgroundColor?: string;
    style?: CSSProperties;
}

export const Button: FC<BtnProps> = ({ backgroundColor, onClick, style, children, ...props }) => {
    return (
        <button style={{ ...ButtonStyle, backgroundColor, ...(style ?? {}) }} onClick={onClick} {...props}>
            {children}
        </button>
    );
};
