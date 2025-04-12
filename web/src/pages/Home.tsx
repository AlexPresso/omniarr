import {useEffect, useState} from "react";
import * as React from "react";
import {Media} from "../types/Media.ts";
import MediaCard from "../components/MediaCard.tsx";

export default function Home() {
    const [searchTerms, setSearchTerms] = useState("");
    const [searched, setSearched] = useState(false);
    const [results, setResults] = useState([] as Media[]);

    const renderResults = () => {
        if(!searched)
            return <></>

        if(results.length === 0 && searchTerms !== "")
            return <p className="text-center text-lg font-semibold text-neutral">No results found.</p>

        return results.map((media: Media, i) => <MediaCard key={i} media={media} />)
    }

    const handleSearch = async (e: React.FormEvent | null) => {
        e && e.preventDefault();
        if (!searchTerms)
            return;

        try {
            const response = await fetch(`/api/medias/search?q=${searchTerms}`);
            const json = await response.json();
            setResults(json.data || []);
        } catch (error) {
            console.error("Error fetching search results:", error);
            setResults([]);
        }
    };

    useEffect(() => {
        let timerId = setTimeout(async () => {
            setSearched(false);
            await handleSearch(null);
            setSearched(true);
        }, 300);

        return () => clearTimeout(timerId);
    }, [searchTerms])

    return (
        <div className="p-4 flex flex-col items-center min-h-[80vh]">
            <div className="w-full max-w-xl mt-8">
                <form onSubmit={handleSearch} className="relative w-full">
                    <input
                        type="text"
                        className="input input-bordered w-full pr-16 rounded-full"
                        placeholder="Search movies, series..."
                        value={searchTerms}
                        onChange={(e: React.ChangeEvent<HTMLInputElement>) => setSearchTerms(e.target.value)}
                    />
                    <button type="submit" className="btn btn-primary absolute top-0 right-0 rounded-full h-full px-4">
                        Search
                    </button>
                </form>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-6 mt-8 w-full max-w-5xl">
                {renderResults()}
            </div>
        </div>
    );
}