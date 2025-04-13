import {Media} from "../types/Media.ts";
import MediaCard from "../components/MediaCard.tsx";
import SearchBar from "../components/SearchBar.tsx";
import {useSearchContext} from "../contexts/SearchContext.tsx";
import {useState} from "react";

export default function Home() {
    const { results, setResults } = useSearchContext()
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState<Error | null>(null);

    const renderResults = () => {
        if(!results)
            return <></>

        if(results.length === 0)
            return <p className="text-center text-lg font-semibold text-neutral">No results found.</p>

        return results.map((media: Media, i) => <MediaCard key={i} media={media} />)
    }

    return (
        <div className="p-4 flex flex-col items-center min-h-[80vh]">
            <div className="w-full max-w-xl mt-8">
                <SearchBar setResults={setResults} setLoading={setLoading} setError={setError} />
            </div>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-8 mt-10 w-full max-w-5xl">
                {renderResults()}
            </div>
        </div>
    );
}