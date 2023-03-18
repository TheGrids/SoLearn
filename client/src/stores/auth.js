import {ref} from 'vue'
import {defineStore} from 'pinia'
import AuthService from "@/services/auth.service";
import jwtDecode from 'jwt-decode';

export const useAuthStore = defineStore('auth', () => {
    const user = ref(null);

    function checkAuth() {
        const token = AuthService.getToken();
        if (!token) {
            return;
        }
        try {
            const userFromToken = jwtDecode(token);
            if (!userFromToken) {
                return;
            }
            user.value = userFromToken;
        } catch (error) {
            // pass
        }
    }

    async function signIn(formFields) {
        const response = await AuthService.signIn(formFields);
        const token = response.data.access;
        const userFromToken = jwtDecode(token);
        AuthService.saveToken(token);
        user.value = userFromToken;
    }

    async function signUp(formFields) {
        const response = await AuthService.signUp(formFields);
        await signIn(formFields);
        return response;
    }

    async function logout() {
        await AuthService.logout();
        user.value = null;
    }

    return {
        user,
        signIn,
        signUp,
        logout,
        checkAuth,
    };
})
