import {ApiResponse} from "../types/ApiResponse.ts";

export async function getRequest(url: string, setLoading: Function, setError: Function) {
    try {
        setLoading(true);
        const response = await fetch(`/api${url}`);
        if(!response.ok) {
            throw new Error(`(${response.status}) An error occurred.`);
        }

        const json: ApiResponse = await response.json();
        if(json.error) {
            throw new Error(`Error: ${json.error}`);
        }

        return json.data;
    } catch (err) {
        console.error(err);
        setError('An error occurred. Check logs for additional information.');
    } finally {
        setLoading(false);
    }
}

export async function postRequest(url: string, data: any, setLoading: Function, setError: Function) {
    try {
        setLoading(true);
        const response = await fetch(`/api${url}`, {
            method: 'POST',
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(data),
        });

        if(!response.ok) {
            throw new Error(`(${response.status}) An error occurred.`);
        }

        const json = await response.json();
        if (json.error) {
            throw new Error(`Error: ${json.error}`);
        }

        return json.data;
    } catch (err) {
        console.error(err);
        setError('An error occurred. Check logs for additional information.');
    } finally {
        setLoading(false);
    }
}

