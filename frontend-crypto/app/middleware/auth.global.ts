
export default defineNuxtRouteMiddleware((to) => {
    const userStore = useUserStore();
    if(!userStore.isAuth && to.path !== '/login') {
        return navigateTo("/login")
    }
})