import axios from "axios";
import AuthService from "@/services/auth.service";

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json'
    },
});

api.interceptors.request.use((request) => {
    request.headers.Authorization = AuthService.getToken();
    return request;
});

const RolesMap = {
    Guest: "0",
    LPP: "1",
    Person: "2",
    Admin: "3"
}

api.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalConfig = error.config;
        if (error.response) {
            const role = error.response.headers.role;
            if (role) {
                if (role === RolesMap.Guest) {
                    const accessToken = await AuthService.refresh();
                    if (!accessToken) {
                        await AuthService.logout();
                    }
                    return axios(originalConfig);
                }
            }

            return Promise.reject(error);
        }

        return Promise.reject(error);
    },
);

export {api}