import {useEffect, useRef} from "react";
import * as React from "react";
import {useSearchContext} from "../contexts/SearchContext.tsx";
import {streamRequest} from "../utils/Requester.tsx";

export default function SearchBar(props: { setResults: Function, onData: (data: any) => void, setLoading: Function, setError: Function }) {
    const { searchQuery, setSearchQuery } = useSearchContext();
    const isFirstRender = useRef(true);
    const lastSearchedQuery = useRef(searchQuery);

    const handleSearch = async (e: React.FormEvent | null) => {
        e && e.preventDefault();
        if(searchQuery === lastSearchedQuery.current)
            return;


        props.setResults(null);
        if (!searchQuery) {
            props.setResults(null);
            return;
        }

        streamRequest(`/medias/search?q=${searchQuery}`, "GET", null, props.onData, props.setLoading, props.setError)
    };

    useEffect(() => {
        if(isFirstRender.current) {
            isFirstRender.current = false;
            return;
        }

        let timerId = setTimeout(() => handleSearch(null), 300);
        return () => clearTimeout(timerId);
    }, [searchQuery])

    return (
        <form onSubmit={handleSearch} className="relative w-full">
            <input
                type="text"
                className="input input-bordered w-full pr-16 rounded-full"
                placeholder="Search movies, series..."
                value={searchQuery}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => setSearchQuery(e.target.value)}
            />
            <button type="submit" className="btn btn-primary absolute top-0 right-0 rounded-full h-full px-4">
                Search
            </button>
        </form>
    )
}