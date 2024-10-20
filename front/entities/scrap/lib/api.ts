import { api } from "../../../shared/api";
import { Scrap } from "../Scrap";

export const getScrapList = async ({ page = 1, pined = false }: { page?: number; pined?: boolean }) => {
    const params = { page, limit: 10, pined };

    try {
        const response = await api.get<{ list: Scrap[]; count: number }>("scraps", { params });
        return response.data;
    } catch {
        return { list: [], count: 0 };
    }
};

export const getScrap = async (scrapId: number | string | undefined) => {
    if (scrapId === undefined) return null;
    try {
        const response = await api.get<Scrap>(`/scraps/${scrapId}`);
        return response.data;
    } catch {
        return undefined;
    }
};

export const createScrap = async (url: string) => {
    try {
        const response = await api.post<Scrap>("/scraps", { url });
        return response.data;
    } catch {
        return undefined;
    }
};

export const reScrap = async (id: number) => {
    try {
        const response = await api.post<Scrap>(`/scraps/${id}`);
        return response.data;
    } catch {
        return undefined;
    }
};
