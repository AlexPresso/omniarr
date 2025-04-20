import {createContext, useState, useContext, ReactNode, Dispatch, SetStateAction} from 'react';
import { Media } from '../types/Media';

interface SearchContextType {
    searchQuery: string;
    setSearchQuery: (query: string) => void;
    results: Media[] | null;
    setResults: Dispatch<SetStateAction<Media[] | null>>;
}

const SearchContext = createContext<SearchContextType | undefined>(undefined);

export const SearchProvider = ({ children }: { children: ReactNode }) => {
    const [searchQuery, setSearchQuery] = useState("");
    const [results, setResults] = useState<Media[] | null>(null);

    return (
        <SearchContext.Provider value={{ searchQuery, setSearchQuery, results, setResults }}>
            {children}
        </SearchContext.Provider>
    );
};

export const useSearchContext = (): SearchContextType => {
    const context = useContext(SearchContext);
    if (context === undefined) {
        throw new Error("useSearchContext must be used within a SearchProvider");
    }
    return context;
};