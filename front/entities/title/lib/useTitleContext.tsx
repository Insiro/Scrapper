import { useContext } from "react";
import { TitleContext, TitleContextType } from "./titleContext";

// Context를 쉽게 사용하기 위한 커스텀 훅
export const useTitleContext = (): TitleContextType => {
    const context = useContext(TitleContext);
    if (!context) {
        throw new Error("useLayoutContext must be used within a LayoutProvider");
    }
    return context;
};
