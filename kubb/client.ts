export type RequestConfig<TData = unknown> = {
    baseURL?: string;
    url: string;
    method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH' | "OPTIONS";
    params?: object;
    data?: TData | FormData;
    responseType?: 'arraybuffer' | 'json' | 'text' | 'blob' | 'document';
    signal?: AbortSignal;
    headers?: HeadersInit;
};

export type ResponseConfig<TData = unknown, TError = unknown> = {
    data: TData;
    status: number;
    statusText: string;
};

export type ResponseErrorConfig<TError = unknown> = {
    data: TError;
    status: number;
    statusText: string;
};


// Use a variável de ambiente para a baseURL
const isDev = process.env.NODE_ENV === 'development' || __DEV__;
const API_BASE_URL = process.env.EXPO_PUBLIC_API_BASE_URL 

const client = async <TData, TError = unknown, TVariables = unknown>(
    config: RequestConfig<TVariables>
): Promise<ResponseConfig<TData, TError>> => {
    // Monta a baseURL
    const baseURL = config.baseURL ?? API_BASE_URL;
    let fullUrl = `${baseURL}${config.url}`;

    // Adiciona params como query string se existirem
    if (config.params) {
        const searchParams = new URLSearchParams();
        Object.entries(config.params).forEach(([key, value]) => {
            if (value !== undefined && value !== null) {
                searchParams.append(key, String(value));
            }
        });
        const queryString = searchParams.toString();
        if (queryString) {
            fullUrl += (fullUrl.includes('?') ? '&' : '?') + queryString;
        }
    }

    // Configurar headers padrão
    const defaultHeaders: HeadersInit = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    };

    // Merge headers customizados com os padrão
    const finalHeaders = {
        ...defaultHeaders,
        ...config.headers,
    };

    // Se o body é FormData, remover Content-Type para deixar o browser definir
    if (config.data instanceof FormData) {
        delete (finalHeaders as any)['Content-Type'];
    }

    console.log('=== REQUEST DEBUG ===');
    console.log('Method:', config.method?.toUpperCase() || 'GET');
    console.log('URL:', fullUrl);
    console.log('Headers:', finalHeaders);

    // Log completo do body da request
    if (config.data instanceof FormData) {
        console.log('Body Type: FormData');
        console.log('FormData entries:');
        for (const [key, value] of (config.data as any).entries()) {
            console.log(`  ${key}:`, value);
        }
    } else if (config.data) {
        console.log('Body Type: JSON');
        console.log('Body (raw):', config.data);
        console.log('Body (stringified):', JSON.stringify(config.data, null, 2));
    } else {
        console.log('Body: No body data');
    }
    console.log('====================');

    try {
        const response = await fetch(fullUrl, {
            method: config.method?.toUpperCase() || 'GET',
            body: config.data instanceof FormData ? config.data : config.data ? JSON.stringify(config.data) : undefined,
            signal: config.signal,
            headers: finalHeaders,
        });

        console.log('=== RESPONSE DEBUG ===');
        console.log('Status:', response.status);
        console.log('StatusText:', response.statusText);
        console.log('Headers:', Object.fromEntries(response.headers.entries()));
        console.log('=====================');

        let data: TData | TError | string | undefined;
        const contentType = response.headers.get('content-type');

        if (contentType && contentType.includes('application/json')) {
            data = await response.json().catch(() => undefined);
        } else {
            const textData = await response.text().catch(() => '');
            // console.log('Non-JSON response:', textData);
            data = textData;
        }

        if (!response.ok) {
            console.log('=== ERROR RESPONSE ===');
            console.log('Error Data:', data);
            console.log('====================');

            throw {
                data: data as TError,
                status: response.status,
                statusText: response.statusText,
            } as ResponseErrorConfig<TError>;
        }

        console.log('=== SUCCESS RESPONSE ===');
        console.log('Success Data:', data);
        console.log('=======================');

        return {
            data: data as TData,
            status: response.status,
            statusText: response.statusText,
        };
    } catch (error) {
        console.log('=== NETWORK ERROR ===');
        console.log('Error:', error);
        console.log('====================');
        throw error;
    }
};

export default client;
