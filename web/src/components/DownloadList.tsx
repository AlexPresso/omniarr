import {Download} from "../types/Download.ts";
import {MediaDetails} from "../types/Media.ts";
import {useEffect, useState} from "react";
import apiRequest from "../utils/Requester.tsx";

export default function DownloadList(props: {media: MediaDetails}) {
    const [downloads, setDownloads] = useState<Download[] | null>([]);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    useEffect(() => apiRequest(`/downloads/search?q=${props.media.title}`, setDownloads, setLoading, setError), [])

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
                            <td>{download.size}</td>
                            <td>{download.seeders}</td>
                            <td>{download.leechers}</td>
                            <td>
                                <a className="btn btn-primary btn-sm" href={download.downloadUrl} target="_blank" rel="noopener noreferrer">
                                    Download
                                </a>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}