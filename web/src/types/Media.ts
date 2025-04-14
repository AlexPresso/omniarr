import {Download} from "./Download.ts";

export type Media = {
    id: string;
    title: string;
    description: string;
    popularity: number;
    releaseDate: string;
    cover: string;
    type: string;
};

export type MediaDetails = Media & {
    downloads: Download[];
}