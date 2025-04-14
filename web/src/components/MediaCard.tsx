import {Media} from "../types/Media.ts";
import {useNavigate} from "react-router";

const baseUrls: Record<string, string> = {
    "movie": "https://image.tmdb.org/t/p/w500",
    "tv": "https://image.tmdb.org/t/p/w500",
}

export default function MediaCard(props: { media: Media }) {
    const navigate = useNavigate();
    const navigateToDetails = () => {
        navigate(`/media/${props.media.type}:${props.media.id}`)
    }

    return (
        <div onClick={navigateToDetails} className="card shadow-xl cursor-pointer group relative overflow-hidden">
            <figure className="w-full h-full">
                <img
                    src={`${baseUrls[props.media.type]}${props.media.cover}`}
                    alt={props.media.title}
                    className="object-cover w-full h-full"
                />
            </figure>
            <div className="absolute inset-0 bg-transparent group-hover:bg-black/65 transition duration-300 flex justify-center items-center">
                <h2 className="text-xl font-semibold text-white opacity-0 group-hover:opacity-100 transition-opacity duration-300 text-center px-2">
                    {props.media.title}
                </h2>
            </div>
        </div>
    );
}