import {ref} from 'vue'
import {defineStore} from 'pinia'
import AuthService from "@/services/auth.service";

export const useAuthStore = defineStore('auth', () => {
    const user = ref(null);

    async function signIn(formFields) {
        const response = await AuthService.signIn(formFields);
        // const token = response.data.token;
        // const userFromToken = null;
        // // const user = {
        // //     id:
        // // }
        // user.value = user;
    }

    async function signUp(formFields) {
        const response = await AuthService.signUp(formFields);
        await signIn(formFields);
        return response;
    }

    function logout() {
        user.value = null;
    }

    return {
        signIn,
        signUp,
        logout,
    };
})
