import {ApiResponse} from "../types/ApiResponse.ts";

export default function apiRequest(endpoint: string, setData: Function, setLoading: Function, setError: Function) {
    (async () => {
        try {
            setLoading(true);
            const response = await fetch(`/api${endpoint}`);
            if(!response.ok) {
                throw new Error(`(${response.status}) An error occurred.`);
            }

            const json: ApiResponse = await response.json();
            if(json.error) {
                throw new Error(`Error: ${json.error}`);
            }

            setData(json.data);
        } catch (err) {
            console.error(err);
            setError('An error occurred. Check logs for additional information.');
        } finally {
            setLoading(false);
        }
    })()
}