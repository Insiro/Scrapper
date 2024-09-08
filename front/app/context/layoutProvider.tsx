import { ReactNode, useState } from "react";
import { TitleContext } from "../../entities/title/lib/titleContext";

interface LayoutProviderProps {
    children: ReactNode;
}

export const LayoutProvider: React.FC<LayoutProviderProps> = ({ children }) => {
    const [pageTitle, setPageTitle] = useState<string>("Home");

    return <TitleContext.Provider value={{ pageTitle, setPageTitle }}> {children} </TitleContext.Provider>;
};
