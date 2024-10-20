import { createContext } from "react";

export interface TitleContextType {
    pageTitle: string;
    setPageTitle: (title: string) => void;
}
export const TitleContext = createContext<TitleContextType | undefined>(undefined);
