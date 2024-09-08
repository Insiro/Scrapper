export interface Scrap {
    id: number;
    url: string;
    content?: string;
    author_name: string;
    author_tag: string;
    image_names: string[];
    comment?: string;
}