import {Download} from "./Download.ts";

export type Media = {
    id: string;
    title: string;
    originalTitle: string;
    description: string;
    popularity: number;
    releaseDate: string;
    cover: string;
    type: string;
};

export type MediaDetails = Media & {
    downloads: Download[];
}

export function getCoverURL(type: string, cover: string): string {
    switch(type) {
        case "movie":
        case "tv":
            return `https://image.tmdb.org/t/p/w500/${cover}`;
        case "book":
            return `https://covers.openlibrary.org/b/id/${cover}-M.jpg`;
        default:
            return ""
    }
}