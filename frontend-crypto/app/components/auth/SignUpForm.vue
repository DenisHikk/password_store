<script setup lang="ts">
import { z } from 'zod'

type ZX = typeof import('zxcvbn');
const strongPassword = ref<number>(20);
const show = ref<boolean>(false);
const modzxc = ref<ZX | null>(null);

onMounted(async () => {
    const mod = await import('zxcvbn');
    modzxc.value = mod.default ? (mod as any).default : (mod as any); 
})


const form = reactive({
    email: '',
    password: '',
    repeatPassword: '',
    masterPassword: '',
    rememberMe: false
})

function zxcvbnPercent(res: ReturnType<ZX>): number {
    const entropyBits = res.guesses_log10 * 3.32192809489; // log2(10)
    const capped = Math.min(entropyBits, 60);
    return Math.round((capped / 60) * 100); // 0..100
}

function scorePassword(pw: string) {
    if(modzxc.value) {
        const res = modzxc.value!(pw)
        return zxcvbnPercent(res)
    }
    return 0;
}

const passwordScore = computed(() => scorePassword(form.password))
const progressColor = computed(() => {
    const s = passwordScore.value;
    if (s < 40) return 'error'
    if (s < 70) return 'warning'
    if (s < 90) return 'success'
    return 'success'
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
                size="xl" class="w-full mt-6" autocomplete="new-password" variant="subtle">
                <template #trailing>
                    <UButton variant="ghost" color="info"
                        :trailing-icon="show ? 'i-heroicons-eye-20-solid' : 'i-heroicons-eye-slash-20-solid'"
                        @click="show = !show" />
                </template>
            </UInput>
        </UFormField>
        <UProgress :max=100 :color="progressColor" v-model="passwordScore"/>
        <UFormField required>
            <UInput placeholder="Repeat password" v-model="form.repeatPassword" type='password' color="neutral" size="xl" class="w-full mt-6"
                autocomplete="new-password" variant="subtle" />
        </UFormField>
        <UFormField required>
            <UInput placeholder="Master Password" v-model="form.masterPassword" type='password' color="neutral" size="xl" class="w-full mt-6"
                autocomplete="new-password" variant="subtle" />
        </UFormField>
        <UFormField class="mt-6">
            <UCheckbox indicator="end" v-model="form.rememberMe" label="Запомнить меня" />
        </UFormField>
        <UButton class="mt-6 mb-6" variant="solid" color="secondary" block size="lg" type="submit">Зарегестрироваться
        </UButton>
    </UForm>
</template>