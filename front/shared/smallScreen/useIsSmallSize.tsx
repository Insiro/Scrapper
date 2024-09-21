import { useContext, useEffect } from "react";
import { SmallScreenContext } from "./SmallScreenContext";

export const useIsSmallScreenHook = (setIsSmallScreen: (arg: boolean) => void) => {
    useEffect(() => {
        // 화면 크기 변경 감지
        const handleResize = () => setIsSmallScreen(window.innerWidth <= 600);
        window.addEventListener("resize", handleResize);

        return () => window.removeEventListener("resize", handleResize);
    }, [setIsSmallScreen]);
    return <></>;
};

export const useIsSmallScreen = () => {
    return useContext(SmallScreenContext);
};
