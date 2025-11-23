<script lang="ts" setup>
const userStore = useUserStore()
const route = useRoute()

const items = computed(() => [
    {
        label: 'Пароли',
        icon: route.path === '/' ? 'i-lucide-lock-keyhole-open' : 'i-lucide-lock-keyhole',
        to: '/',
        active: route.path === '/',
    },
    {
        label: 'Заметки',
        icon: 'i-lucide-sticky-note',
        to: '/notes',
    },
    {
        label: 'Добавить сервис',
        icon: 'i-lucide-plus',
        to: '/add',
        children: [
            {
                label: 'Сервис',
                icon: 'i-lucide-monitor-cloud',
                to: '/add?service',
                active: route.fullPath === '/add?service'
            },
            {
                label: 'SSH',
                icon: 'i-lucide-terminal',
                to: '/add?ssh',
                active: route.fullPath === '/add?ssh'
            }
            ,
            {
                label: 'Заметки',
                icon: 'i-lucide-sticky-note'
                ,
                to: '/add?note',
                active: route.fullPath === '/add?note'
            }
        ],
    }

])

</script>
<template>
    <UHeader title="CryptoStore">
        <template #left>
            <UNavigationMenu
                :ui="{
                    list: 'gap-2',
                    childList: 'flex flex-col', 
                    viewport: 'max-w-52'
                }"
                :collapsible="true"
                v-if="userStore.isAuth" 
                :items="items" />
        </template>
        <template #right>
            <div class="flex justify-center gap-4">
                
                <!-- <UColorModeButton /> -->
                <div v-if="userStore.isAuth" class="flex justify-center gap-4">
                    <UInput icon="i-lucide-search" size="md" variant="outline" placeholder="Search..." />
                    <UButton icon="i-lucide-settings" size="md" color="neutral" variant="soft"></UButton>
                    <UAvatar src="https://github.com/benjamincanac.png" />
                </div>
            </div>
        </template>
    </UHeader>
</template>