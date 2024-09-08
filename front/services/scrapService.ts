
import axios from 'axios';
import { Scrap } from '../types';

const API_URL = 'http://localhost:8000/api/scraps';

export const getScraps = async () => {
    const response = await axios.get<Scrap[]>(API_URL);
    console.log(response.data)
    return response.data;
};


export const getScrap = async (scrapId: number | string | undefined) => {
    if (scrapId === undefined)
        return null
    const response = await axios.get<Scrap>(API_URL + `/${scrapId}`);
    console.log(response.data)
    return response.data;
};


export const createScrap = async (url: string, force: boolean) => {
    const response = await axios.post(API_URL, { url, force });
    return response.data;
};