import { defineStore } from "pinia";

interface User {
    email: string,
    password: string,
    name: string,
}

export const useUserStore = defineStore('userStore', ()=> {
    const user = ref<User | null>(null);
    const isAuth = ref<boolean>(false);

    return {
        user, isAuth
    }
})