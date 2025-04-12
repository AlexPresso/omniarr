import {Media} from "../types/Media.ts";

export default function MediaCard(props: { media: Media }) {
    return (
        <div className="card shadow-lg">
            <figure className="h-64 w-full">
                <img src={props.media.cover} alt={props.media.title} className="object-cover h-full w-full" />
            </figure>
            <div className="card-body flex items-center justify-center p-4">
                <h2 className="text-lg font-bold text-center">{props.media.title}</h2>
            </div>
        </div>
    )
}