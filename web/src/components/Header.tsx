import { FiSettings } from 'react-icons/fi';
import {Link} from "react-router";

export default function Header() {
    return (
        <div className="navbar bg-base-100 shadow-md sticky top-0 z-50">
            <div className="flex-1">
                <Link to={"/"} className="btn btn-ghost normal-case text-xl">Omniarr</Link>
            </div>
            <div className="flex-none gap-2">
                <Link to={"/settings"} className="btn btn-ghost btn-circle" aria-label="Settings">
                    <FiSettings size={20} />
                </Link>
            </div>
        </div>
    )
}
