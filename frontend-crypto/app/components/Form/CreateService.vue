<script setup lang="ts">

const tags = ref([''])

const note = ref('')

const password = ref('')


</script>

<template>
    <UForm class="flex flex-col gap-5" >
        <UFormField 
            label="Сервис/Сайт/Приложение"
            size="xl" 
            required>
            <UInput placeholder="Например: интернет-магазин" class="flex "/>
        </UFormField>

        <UFormField 
            label="URL"
            size="xl" 
            required>
            <UInput placeholder="Например: интернет-магазин.рф" class="flex " />
        </UFormField>

        <UFormField 
            label="Логин/Email" 
            size="xl" 
            required>
            <UInput placeholder="Например: user@mail.ru" class="flex " />
        </UFormField>

        <UFormField 
            label="Пароль"
            size="xl" 
            required >
            <UInput v-model="password" type="password" class="flex " />
        </UFormField>
        <UModal
            :ui="{
                content: 'max-h-[calc(100dvh)] sm:max-h-[calc(100dvh)]'
            }"
        >
            <UButton 
                variant="soft" 
                color="neutral" 
                type="submit" 
                class="block"
            >
                Сгенерировать пароль
            </UButton>
            
            <template #content>
                <FormGeneratePassword />
            </template>
        </UModal>

        <UFormField  label="Теги">
            <UInputTags class="flex"  v-model="tags" placeholder="Например: покупки, личное" />
        </UFormField>

        <UFormField  label="Заметки">
            <UTextarea class="flex" v-model="note" />
        </UFormField>
        

        <div class="flex justify-center gap-4">
            <UButton variant="soft" color="neutral" type="submit">
                Отменить
            </UButton>
            <UButton type="submit">
                Сохранить
            </UButton>
        </div>
    </UForm>
</template>
