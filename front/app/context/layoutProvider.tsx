import { ReactNode, useState } from "react";
import { TitleContext } from "../../entities/title/lib/titleContext";
import { useIsSmallScreenHook } from "@/shared/smallScreen/useIsSmallSize";
import { SmallScreenContext } from "@/shared/smallScreen/SmallScreenContext";

interface LayoutProviderProps {
    children: ReactNode;
}

export const LayoutProvider: React.FC<LayoutProviderProps> = ({ children }) => {
    const [pageTitle, setPageTitle] = useState<string>("Home");
    const [isSmallScreen, setIsSmallScreen] = useState<boolean>(window.innerWidth <= 600);

    useIsSmallScreenHook(setIsSmallScreen);

    return (
        <SmallScreenContext.Provider value={isSmallScreen}>
            <TitleContext.Provider value={{ pageTitle, setPageTitle }}> {children} </TitleContext.Provider>)
        </SmallScreenContext.Provider>
    );
};
