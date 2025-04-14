import {useEffect} from "react";
import * as React from "react";
import {useSearchContext} from "../contexts/SearchContext.tsx";
import { getRequest } from "../utils/Requester.tsx";

export default function SearchBar(props: { setResults: Function, setLoading: Function, setError: Function }) {
    const { searchQuery, setSearchQuery } = useSearchContext();
    const handleSearch = async (e: React.FormEvent | null) => {
        e && e.preventDefault();
        if (!searchQuery) {
            props.setResults(null);
            return;
        }

        getRequest(`/medias/search?q=${searchQuery}`, props.setResults, props.setLoading, props.setError);
    };

    useEffect(() => {
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