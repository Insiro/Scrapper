import { color } from "@/shared/constant";
import React from "react";
import { Button } from "./Button";

interface ModalProps {
    isOpen: boolean;
    title: string;
    message: string;
    action: (confirm: boolean) => void;
    disabled?: boolean;
}

const modalOverlayStyle: React.CSSProperties = {
    position: "fixed",
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: "rgba(0, 0, 0, 0.5)",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    zIndex: 1000,
};

const modalStyle: React.CSSProperties = {
    backgroundColor: "#fff",
    padding: "2rem",
    borderRadius: "8px",
    width: "300px",
    textAlign: "center",
};

const buttonContainerStyle: React.CSSProperties = {
    display: "flex",
    justifyContent: "space-between",
    marginTop: "1rem",
};

const ConfirmModal: React.FC<ModalProps> = ({ isOpen, title, message, action, disabled }) => {
    if (!isOpen) return null; // 모달이 열려있지 않으면 렌더링하지 않음

    return (
        <div style={modalOverlayStyle}>
            <div style={modalStyle}>
                <h2>{title}</h2>
                <p>{message}</p>
                <div style={buttonContainerStyle}>
                    <Button onClick={() => action(false)} style={{ backgroundColor: color.gray1 }} disabled={disabled}>
                        Cancel
                    </Button>
                    <Button
                        onClick={() => action(true)}
                        style={{ backgroundColor: color.red, color: "#fff" }}
                        disabled={disabled}
                    >
                        Confirm
                    </Button>
                </div>
            </div>
        </div>
    );
};

export default ConfirmModal;
