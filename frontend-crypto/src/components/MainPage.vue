<script setup lang="ts">
import { ref, computed } from 'vue'

type Req = {
    length: number
    use_lower: boolean
    use_upper: boolean
    use_digit: boolean
    use_symbol: boolean
}
type Res = { password: string }

const form = ref({
    length: 16,
    useLower: true,
    useUpper: true,
    useDigit: true,
    useSymbol: false,
})

const password = ref('')
const loading = ref(false)
const snackbar = ref({ show: false, text: '' })

const anySet = computed(
    () => form.value.useLower || form.value.useUpper || form.value.useDigit || form.value.useSymbol
)

function notify(t: string) {
    snackbar.value.text = t
    snackbar.value.show = true
}

async function generate() {
    if (!anySet.value) { notify('Выберите набор символов'); return }
    if (form.value.length < 4 || form.value.length > 128) { notify('Длина 4–128'); return }

    loading.value = true
    try {
        const payload: Req = {
            length: form.value.length,
            use_lower: form.value.useLower,
            use_upper: form.value.useUpper,
            use_digit: form.value.useDigit,
            use_symbol: form.value.useSymbol,
        }
        const res = await fetch('/password', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload),
        })
        if (!res.ok) { notify(`Ошибка ${res.status}`); return }
        const data: Res = await res.json()
        password.value = data.password ?? ''
        notify('Готово')
    } catch {
        notify('Сеть недоступна')
    } finally {
        loading.value = false
    }
}

async function copy() {
    if (!password.value) return
    try { await navigator.clipboard.writeText(password.value); notify('Скопировано') }
    catch { notify('Не удалось скопировать') }
}
</script>

<template>
    <v-container class="py-8" max-width="760">
        <v-card elevation="2">
            <v-toolbar color="primary" density="comfortable" title="Генератор паролей" />

            <v-card-text>
                <v-row>
                    <v-col cols="12" sm="6">
                        <v-text-field v-model.number="form.length" type="number" :min="4" :max="128" label="Длина"
                            prepend-inner-icon="mdi-ruler" variant="outlined" density="comfortable" />
                        <v-alert v-if="!anySet" type="warning" variant="tonal" density="compact" class="mt-2">
                            Выберите хотя бы один набор
                        </v-alert>
                    </v-col>

                    <v-col cols="12" sm="6">
                        <v-switch v-model="form.useLower" label="Строчные a–z" inset />
                        <v-switch v-model="form.useUpper" label="Заглавные A–Z" inset />
                        <v-switch v-model="form.useDigit" label="Цифры 0–9" inset />
                        <v-switch v-model="form.useSymbol" label="Символы !@#" inset />
                    </v-col>
                </v-row>

                <v-row class="mt-2" align="center">
                    <v-col cols="12" sm="8" class="d-flex ga-2">
                        <v-btn color="primary" :loading="loading" @click="generate" prepend-icon="mdi-shield-key">
                            Сгенерировать
                        </v-btn>
                        <v-btn :disabled="!password" @click="copy" prepend-icon="mdi-content-copy">
                            Копировать
                        </v-btn>
                    </v-col>
                </v-row>

                <v-text-field class="mt-4" v-model="password" label="Пароль" variant="outlined" density="comfortable"
                    readonly :append-inner-icon="password ? 'mdi-content-copy' : undefined"
                    @click:append-inner="password && copy()" />
            </v-card-text>
        </v-card>

        <v-snackbar v-model="snackbar.show" :timeout="2000">
            {{ snackbar.text }}
        </v-snackbar>
    </v-container>
</template>