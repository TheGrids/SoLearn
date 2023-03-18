import {api} from "@/http";

const ACCESS_TOKEN_KEY = 'access_token';

class AuthService {
    constructor() {
        this.httpClient = api;
    }

    async signIn(formFields) {
        return this.httpClient.post('login', {
            email: formFields.email,
            password: formFields.password,
        });
    }

    async signUp(formFields) {
        return this.httpClient.post('register', {
            email: formFields.email,
            password: formFields.password,
            first_name: formFields.firstName,
            last_name: formFields.lastName,
        });
    }

    async logout() {
        try {
            this.deleteToken();
            await this.httpClient.get('logout');
        } catch (error) {
            // pass
        }
    }

    async refresh() {
        const response = await this.httpClient.get('refresh');
        this.saveToken(response.data.access);
    }

    getToken() {
        return localStorage.getItem(ACCESS_TOKEN_KEY);
    }

    saveToken(token) {
        localStorage.setItem(ACCESS_TOKEN_KEY, token);
    }

    deleteToken() {
        localStorage.removeItem(ACCESS_TOKEN_KEY);
    }
}

export default new AuthService();