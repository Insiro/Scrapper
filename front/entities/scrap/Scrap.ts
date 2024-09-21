import { ScrapImage } from "../scrapImage/ScrapImage";

export interface Scrap {
    id: number;
    url: string;
    content?: string;
    author_name: string;
    author_tag: string;
    images: ScrapImage[];
    comment?: string;
    source?: string;
}
