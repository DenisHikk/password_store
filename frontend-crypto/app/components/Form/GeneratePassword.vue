<script setup lang="ts">
    import { useClipboard } from '@vueuse/core'

    import passfather from 'passfather';

    const emit = defineEmits(['update:modelValue'])

    defineProps(['modelValue'])

    const charSet = ref([
        {
            label: 'Заглавные буквы (A–Z)',
            value: 'uppercase'
        }, 
        {
            label: 'Строчные буквы (a-z)',
            value: 'lowercase'
        }, 
        {
            label: 'Цифры (0-9)', 
            value: 'numbers'
        },
        {
            label: "Специальные символы (!@#$%^&*()_+~`|}{[]\\:;?><,./-=\\')",
            value: 'specialchar'
        }
    ])  
    const charSetValue = ref(['uppercase','lowercase', 'numbers'])

    const options = ref([
        {
            label: 'Исключить похожий набор символов',
            description: 'Исключает повторение паролей (рекомендуется)',
            value: 'ExcludeSimilar'
        },
        {
            label: 'Требовать каждый набор символов',
            description: 'Гарантирует, что пароль будет содержать хотя бы по одному символу из каждой выбранной категории',
            value: 'RequireEach'
        }
    ])
    const optionsValue = ref(['ExcludeSimilar'])

    const lenghtPassword = ref(50)
    const password = ref('')

    const { copy, copied } = useClipboard()



    const generatePassord = () => {

        var set = {
            uppercase: false,
            lowercase: false,
            numbers: false,
            symbols: false,
            length: lenghtPassword.value,
        }

        charSetValue.value.forEach(charVal => {
            switch (charVal) {
                case 'uppercase':
                    set.uppercase = true;
                    break;
                case 'lowercase':
                    set.lowercase = true;
                    break;
                case 'numbers':
                    set.numbers = true;
                    break;
                case 'symbols':
                    set.symbols = true;
                    break;
                default:
                    break;
            }
        });

        password.value = passfather(set);

    }

    const updateModel = () => {
        emit('update:modelValue', 'Новое значение из ребенка')
    }

    
</script>

<template>
    <div class="p-10 ">
        <h2 class="text-xl mb-5">Генерация пароля</h2>
        <UForm  class="flex flex-col gap-5 overflow-y-auto" >

        <UFormField label="Длина пароля">
            <USlider 
                size="xs" 
                tooltip  
                :min="16"
                :max="64"
                color="info" 
                class="mt-3"  
                v-model="lenghtPassword" />
        </UFormField>
            

            <UFormField  label="Наборы символов">
                <UCheckboxGroup 
                    class="mt-3"
                    size="lg" 
                    :ui="{
                        fieldset: 'gap-y-3'
                    }"
                    color="info" 
                    v-model="charSetValue" 
                    :items="charSet" />
            </UFormField>

            <UFormField label="Наборы символов">
                <UCheckboxGroup
                    class="mt-3" 
                    color="info" 
                    size="lg" 
                    :ui="{
                        fieldset: 'gap-y-3'
                    }"
                    v-model="optionsValue" 
                    :items="options" />
            </UFormField>
            <div class=" inline-flex flex-col gap-y-4">
                <UButton 
                    class="mb-10 justify-center" 
                    variant="soft" 
                    color="info" 
                    type="submit"
                    @click="generatePassord">
                    Сгенерировать пароль
                </UButton>
                <UInput

                    v-model="password"
                    @input="$emit('update:modelValue', $event.target.value)"
                    placeholder="Пароль"
                    color="info"
                    :ui="{ trailing: 'pr-3' }"
                    size="xl"
                    type='text'
                >
                <template v-if="password?.length" #trailing>
                    <UButton
                        color="info" 
                        variant="link"
                        :icon="copied ? 'i-lucide-copy-check' : 'i-lucide-copy'"
                        aria-label="Скопировать пароль"
                        @click="copy(password)"
                        />
                </template>

                </UInput>
                <UButton 
                    class="justify-center" 
                    color="info"  
                    type="submit"
                    @click="updateModel">
                    Вставить в форму
                </UButton>
            </div>
        </UForm>
    </div>
</template>