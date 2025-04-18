import {getCoverURL, Media} from "../types/Media.ts";
import {useNavigate} from "react-router";

export default function MediaCard(props: { media: Media }) {
    const navigate = useNavigate();
    const navigateToDetails = () => {
        navigate(`/media/${props.media.type}:${props.media.id}`)
    }

    return (
        <div onClick={navigateToDetails} className="card shadow-xl cursor-pointer group relative overflow-hidden">
            <figure className="w-full h-full">
                <img
                    src={getCoverURL(props.media.type, props.media.cover)}
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