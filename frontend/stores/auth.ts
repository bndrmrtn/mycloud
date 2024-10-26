import {defineStore} from "pinia";
import type {User} from "~/types/user";
import {computed} from "vue";

export const useAuthStore = defineStore('auth', () => {
    const user = ref<User|null>(null)

    const logout = () => {
        user.value = null
    }

    const isLoggedIn = computed(() => user.value != null)

    const handle = async (api: string) => {
        try {
            const res = await fetch(`${api}/me`)
            if(res.status != 200) return
            const data = await res.json()
            user.value = data as User
        } catch (e: unknown) {
            console.error(e)
        }
    }

    return {user, logout, isLoggedIn, handle}
})