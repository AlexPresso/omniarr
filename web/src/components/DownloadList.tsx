import {Download} from "../types/Download.ts";
import {MediaDetails} from "../types/Media.ts";
import {useEffect, useState} from "react";
import {postRequest} from "../utils/Requester.tsx";

export default function DownloadList(props: {media: MediaDetails}) {
    const [downloads, setDownloads] = useState<Download[] | null>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    useEffect(() => {fetchDownloads(setLoading, setError).then(setDownloads)}, []);

    const fetchDownloads = async (setError: Function, setLoading: Function) => {
        return await postRequest("/downloads/query", {
            externalId: props.media.id,
            title: props.media.title,
            originalTitle: props.media.originalTitle,
            type: props.media.type,
            year: props.media.releaseDate?.split("-")[0] || ""
        }, setLoading, setError)
    }

    const queueDownload = async (_: string) => {}

    if(loading)
        return (
            <div className="flex justify-center my-24">
                <span className="loading loading-ring loading-lg"></span>
            </div>
        )

    if(error)
        return (
            <div className="card bg-base-200 shadow-xl max-w-md mx-auto mt-24 p-8 text-center">
                <p className="text-error">{error}</p>
            </div>
        );

    if(!downloads)
        return <div className="alert alert-error shadow-lg">An error occurred while fetching downloads</div>

    if(downloads.length === 0)
        return (
            <div className="alert alert-warning shadow-lg">
                No download available for this media.
            </div>
        )

    return (
        <div className="overflow-x-auto">
            <table className="table table-zebra table-lg">
                <thead>
                <tr>
                    <th>Title</th>
                    <th>Tracker</th>
                    <th>Size</th>
                    <th>Seeds</th>
                    <th>Peers</th>
                    <th>Action</th>
                </tr>
                </thead>
                <tbody>
                    {downloads.map((download: Download, idx) => (
                        <tr key={idx}>
                            <td>{download.title}</td>
                            <td>{download.indexer}</td>
                            <td>{download.size}</td>
                            <td>{download.seeders}</td>
                            <td>{download.leechers}</td>
                            <td>
                                <button
                                    className="btn btn-primary btn-sm"
                                    onClick={() => queueDownload(download.magnetUri || download.link)}
                                >
                                  Download
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}