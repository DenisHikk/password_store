<script setup lang="ts">
const show = ref<boolean>(false);

const form = reactive({
    email: '',
    password: ''
})
</script>

<template>
    <UForm>
        <UFormField required>
            <UInput placeholder="Email" v-model="form.email" type="email" color="neutral" size="xl" class="w-full"
                autocomplete="email" variant="subtle" />
        </UFormField>
        <UFormField required>
            <UInput placeholder="Password" v-model="form.password" :type="show ? 'text' : 'password'" color="neutral"
                size="xl" class="w-full mt-6" autocomplete="current-password" variant="subtle">
                <template #trailing>
                    <UButton variant="ghost" color="info"
                        :trailing-icon="show ?  'i-heroicons-eye-20-solid' : 'i-heroicons-eye-slash-20-solid'"
                        @click="show = !show" />
                </template>
            </UInput>
        </UFormField>
        <UFormField class="mt-6">
            <UCheckbox indicator="end" label="Запомнить меня" />
        </UFormField>
        <UButton class="mt-6 !text-gray-500" block size="xs" variant="ghost" type="button">Забыли
            пароль?</UButton>
        <UButton class="mt-6 mb-6" variant="solid" color="secondary" block size="lg" type="submit">Войти
        </UButton>
    </UForm>
</template>