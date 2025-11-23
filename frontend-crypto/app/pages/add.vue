<script setup lang="ts">
// You might choose this based on an API call or logged-in status
const layout = 'add'
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const tabs = [
    { label: 'Сервисы', icon: '', slot: 'service', value: 'service' },
    { label: 'SSH', icon: '', slot: 'ssh', value: 'ssh' },
    { label: 'Заметки', icon: '', slot: 'notes', value: 'notes' },
]

// Инициализация таба на основе query параметров
const getInitialTab = () => {
    const searchParams = new URLSearchParams(route.fullPath.split('?')[1])
    if (searchParams.has('service')) return 'service'
    if (searchParams.has('ssh')) return 'ssh'
    if (searchParams.has('notes')) return 'notes'
    return 'service' // значение по умолчанию
}

const tab = ref<'service' | 'ssh' | 'notes'>(getInitialTab())

// Следим за изменениями маршрута
watch(() => route.fullPath, (newPath) => {
    const searchParams = new URLSearchParams(newPath.split('?')[1])
    if (searchParams.has('service')) tab.value = 'service'
    else if (searchParams.has('ssh')) tab.value = 'ssh'
    else if (searchParams.has('notes')) tab.value = 'notes'
})
</script>

<template>
  <NuxtLayout :name="layout">
    
    <div class="max-w-120 mx-auto">
      <UTabs :items="tabs" v-model="tab" variant="link" :ui="{ trigger: 'grow' }" class="gap-4 w-full">
        
      </UTabs>

      <div v-if="tab === 'service'" class="space-y-4 mx-auto">
        <FormCreateService>

        </FormCreateService>
      </div>
      <div v-if="tab === 'ssh'" class="space-y-4 mx-auto">
        ssh
      </div>
      <div v-if="tab === 'notes'" class="space-y-4 mx-auto">
        notes
      </div>
    </div>


  </NuxtLayout>
</template>
