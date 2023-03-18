import {api} from "@/http";

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
}

export default new AuthService();