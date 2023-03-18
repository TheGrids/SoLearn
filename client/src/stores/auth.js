import {ref} from 'vue'
import {defineStore} from 'pinia'

export const useAuthStore = defineStore('auth', () => {
    const user = ref({
        id: null,
        email: null,
    })

    async function signUp() {

    }

    function signIn() {

    }

    function logout() {

    }

    return {count, doubleCount, increment}
})
