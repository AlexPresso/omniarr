import { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router';
import {MediaDetails} from "../types/Media.ts";
import DownloadList from "../components/DownloadList.tsx";
import {request} from "../utils/Requester.tsx";
import {getCoverURL} from "../types/Media.ts";

export default function Media() {
    const { id } = useParams<{ id: string }>();

    const [media, setMedia] = useState<MediaDetails | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {request(`/medias/${id}`, "GET", null, setLoading, setError).then(setMedia)}, [id]);

    if (loading) {
        return (
            <div className="flex justify-center my-24">
                <span className="loading loading-ring loading-lg"></span>
            </div>
        );
    }

    if (error || !media) {
        return (
            <div className="card bg-base-200 shadow-xl max-w-md mx-auto mt-24 p-8 text-center">
                <p className="text-error">{error || "Media not found."}</p>
                <Link to="/" className="btn btn-primary mt-4">Return Home</Link>
            </div>
        );
    }

    return (
        <div className="container mx-auto px-4 py-10">
            <div className="card lg:card-side shadow-xl bg-base-200 overflow-hidden">
                <figure className="lg:w-1/3 bg-neutral">
                    <img
                        src={media.cover ? getCoverURL(media.type, media.cover) : '/placeholder.png'}
                        alt={media.title}
                        className="object-cover h-full w-full"
                    />
                </figure>
                <div className="card-body lg:w-2/3">
                    <h2 className="card-title text-3xl font-bold">{media.title}</h2>
                    <p className="text-base-content">{media.description}</p>

                    <div className="text-sm space-y-2 mt-2">
                        <p><strong>Release Date:</strong> {media.releaseDate || "Unknown"}</p>
                        <p><strong>Popularity:</strong> {media.popularity.toFixed(2)}</p>
                    </div>
                </div>
            </div>

            <div className="mt-12">
                <h2 className="text-2xl font-bold mb-4">Torrents</h2>
                <DownloadList media={media} />
            </div>
        </div>
    );
}