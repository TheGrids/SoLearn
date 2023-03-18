import axios, {HttpStatusCode} from "axios";
import AuthService from "@/services/auth.service";

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    headers: {
        'Content-Type': 'application/json'
    },
    withCredentials: true,
});

api.interceptors.request.use((request) => {
    const authHeader = `Bearer ${AuthService.getToken()}`
    request.headers.Authorization = authHeader;
    return request;
});

api.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalConfig = error.config;
        if (error.response) {
            if (error.response.status === HttpStatusCode.Unauthorized && !originalConfig._retry) {
                originalConfig._retry = true;

                // Do something, call refreshToken() request for example;
                // return a request
                await AuthService.refresh();
                return axios(originalConfig);
            }

            return Promise.reject(error);
        }

        return Promise.reject(error);
    },
);

export {api}