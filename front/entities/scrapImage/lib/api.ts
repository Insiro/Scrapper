import { api } from "../../../shared/api";

export const deleteImage = async (imageIds: number[]) => {
    try {
        const response = await api.delete("/images", { data: { images: imageIds } });
        return response.data;
    } catch {
        return undefined;
    }
};
