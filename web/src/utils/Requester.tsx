import {ApiResponse} from "../types/ApiResponse.ts";

export async function request(url: string, method: string, data: any, setLoading: Function, setError: Function) {
    try {
        setLoading(true);
        const request: RequestInit = {
            method: method,
            headers: {
                "Content-Type": "application/json"
            }
        };

        if(method === "POST" && data) {
            request.body = JSON.stringify(data);
            // @ts-ignore
            request.headers["Content-Type"] = 'application/json';
        }

        const response = await fetch(`/api${url}`, request);

        if(!response.ok) {
            throw new Error(`(${response.status}) An error occurred.`);
        }

        const json: ApiResponse = await response.json();
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

export async function streamRequest(
    url: string,
    method: string,
    body: object | null,
    onData: (data: any) => void,
    setLoading: Function,
    setError: Function
) {
    try {
        setLoading(true);
        const request: RequestInit = {
            method: method,
            headers: {
                'Accept': 'application/x-ndjson',
            }
        }

        if(method === "POST") {
            request.body = JSON.stringify(body);
            // @ts-ignore
            request.headers["Content-Type"] = 'application/json';
        }

        const response = await fetch(`/api${url}`, request);

        if (!response.ok) {
            throw new Error(`(${response.status}) An error occurred.`);
        }

        const reader = response.body?.getReader();
        if (!reader) {
            throw new Error('No reader available');
        }

        const stream = new ReadableStream({
            async start(controller) {
                try {
                    while (true) {
                        const { done, value } = await reader.read();
                        if (done)
                            break;

                        controller.enqueue(value);
                    }
                } finally {
                    controller.close();
                    reader.releaseLock();
                }
            }
        });

        const ndjsonStream = stream
            .pipeThrough(new TextDecoderStream())
            .pipeThrough(new TransformStream({
                transform(chunk, controller) {
                    chunk.split('\n')
                        .filter(line => line.trim())
                        .forEach(line => {
                            try {
                                const data = JSON.parse(line);
                                controller.enqueue(data);
                            } catch (e) {
                                console.error('Error parsing JSON:', e);
                            }
                        });
                }
            }));

        const reader2 = ndjsonStream.getReader();
        while (true) {
            const { done, value } = await reader2.read();
            if (done) break;
            onData(value);
        }

    } catch (err) {
        console.error(err);
        setError('An error occurred. Check logs for additional information.');
    } finally {
        setLoading(false);
    }
}
