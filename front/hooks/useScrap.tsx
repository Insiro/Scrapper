import { useCallback } from 'react';
import { createScrap } from '../services/scrapService';

export const useCreateScrap = () => {
    return useCallback(async (url: string, force: boolean) => {
        try {
            const result = await createScrap(url, force);
            console.log('Scrap created:', result);
        } catch (error) {
            console.error('Error creating scrap:', error);
        }
    }, []);
};