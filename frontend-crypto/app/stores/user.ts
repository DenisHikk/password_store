import { defineStore } from "pinia";

interface User {
    email: string,
    password: string,
    name: string,
}

export const useUserStore = defineStore('userStore', ()=> {
    const user = ref<User | null>(null);
    // for debug change to true for unlock another pages without auth
    const isAuth = ref<boolean>(true);

    return {
        user, isAuth
    }
})