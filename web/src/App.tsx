import {BrowserRouter, Route, Routes} from "react-router";
import Home from "./pages/Home.tsx";
import Header from "./components/Header.tsx";
import Settings from "./pages/Settings.tsx";
import {SearchProvider} from "./contexts/SearchContext.tsx";
import Media from "./pages/Media.tsx";

export default function App() {
    return (
        <div className="min-h-screen bg-base-100 text-base-content">
            <BrowserRouter>
                <Header />
                <SearchProvider>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/media/:id" element={<Media />} />
                        <Route path="/settings" element={<Settings />} />
                    </Routes>
                </SearchProvider>
            </BrowserRouter>
        </div>
    )
}
