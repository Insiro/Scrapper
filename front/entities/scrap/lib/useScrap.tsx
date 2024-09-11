import { useCallback } from "react";
import { createScrap } from "./api";

export const useCreateScrap = () => {
    return useCallback(async (url: string) => {
        try {
            const result = await createScrap(url);
            console.log("Scrap created:", result);
        } catch (error) {
            console.error("Error creating scrap:", error);
        }
    }, []);
};
