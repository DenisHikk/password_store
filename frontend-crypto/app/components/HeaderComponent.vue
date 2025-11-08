<script lang="ts" setup>
const userStore = useUserStore()
</script>
<template>
    <UHeader title="CryptoStore">
        <template #right>
            <div v-if="userStore.isAuth">
                <UButton>
                    Кнопка
                </UButton>
            </div>
        </template>
    </UHeader>
</template>