import { ReactNode, useState } from "react";
import { LayoutContext } from "../../entities/title/lib/layoutContext";

interface LayoutProviderProps {
    children: ReactNode;
}

export const LayoutProvider: React.FC<LayoutProviderProps> = ({ children }) => {
    const [pageTitle, setPageTitle] = useState<string>("Home");

    return <LayoutContext.Provider value={{ pageTitle, setPageTitle }}> {children} </LayoutContext.Provider>;
};
