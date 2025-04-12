import {BrowserRouter, Route, Routes} from "react-router";
import Home from "./pages/Home.tsx";
import Header from "./components/Header.tsx";
import Settings from "./pages/Settings.tsx";

export default function App() {
    return (
        <BrowserRouter>
            <div className="min-h-screen bg-base-100 text-base-content">
                <Header />
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/settings" element={<Settings />} />
                </Routes>
            </div>
        </BrowserRouter>
    )
}
