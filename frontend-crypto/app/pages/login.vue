<script setup lang="ts">
const tabs = [
    { label: 'Sign In', icon: '', slot: 'signin', value: 'signin' },
    { label: 'Sign Up', icon: '', slot: 'signup', value: 'signup' },
]
const tab = ref<'signin' | 'signup'>('signin');

</script>

<template>
    <UContainer class="mt-6">
        <UCard class="flex justify-center" variant="soft" :ui="{ root: 'border-none' }">
            <div class="text-center mt-6">
                <p class="text-4xl">Добро пожаловать в CryptoStore</p>
            </div>
            <UTabs :items="tabs" v-model="tab" :content="false" class="mt-6 w-fit mx-auto" variant="link" :ui="{
                list: '!border-b-1 gap-2',
                indicator: 'h-1 rounded'
            }">
            </UTabs>

            <div class="w-full max-w-[420px] mx-auto mt-6 transition duration-150 ease-in-out">
                <!-- sign in -->
                <div v-if="tab === 'signin'" class="space-y-4 mx-auto">
                    <AuthSignInForm />
                </div>

                <!-- sign up -->
                <div v-else class="space-y-4 mx-auto">
                    <AuthSignUpForm />
                </div>
            </div>

        </UCard>
    </UContainer>
</template>