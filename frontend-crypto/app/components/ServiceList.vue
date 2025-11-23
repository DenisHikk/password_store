<script setup lang="ts">
import { h, resolveComponent } from 'vue'
import type { TableColumn } from '@nuxt/ui'
import type { Row } from '@tanstack/vue-table'
import { useClipboard } from '@vueuse/core'

const UButton = resolveComponent('UButton')
const UBadge = resolveComponent('UBadge')

const toast = useToast()
const { copy } = useClipboard()

type Service = {
    id: string
    name: string
    login: string
    tag: string
    updatedAt: string
    status: 'active' | 'inactive' | 'maintenance'
    password: string
}

const data = ref<Service[]>([
    {
        id: '1',
        name: 'API Gateway',
        login: 'admin@company.com',
        tag: 'backend',
        updatedAt: '2024-03-11T15:30:00',
        status: 'active',
        password: 'password123'
    },
    {
        id: '2',
        name: 'User Service',
        login: 'user-service',
        tag: 'microservice',
        updatedAt: '2024-03-11T10:10:00',
        status: 'active',
        password: '21sadfdsf'
    },
    {
        id: '3',
        name: 'Database Cluster',
        login: 'db-admin',
        tag: 'database',
        updatedAt: '2024-03-11T08:50:00',
        status: 'maintenance',
        password: '21sadfdsf'
    },
    {
        id: '4',
        name: 'Cache Service',
        login: 'redis-user',
        tag: 'cache',
        updatedAt: '2024-03-10T19:45:00',
        status: 'inactive',
        password: '21sadfdsf'
    },
    {
        id: '5',
        name: 'File Storage',
        login: 'storage@company.com',
        tag: 'storage',
        updatedAt: '2024-03-10T15:55:00',
        status: 'active',
        password: '21sadfdsf'
    }
])
// Реф для отслеживания открытого dropdown
const openDropdownId = ref<string | null>(null)

// Функция для переключения dropdown
const toggleDropdown = (id: string) => {
    openDropdownId.value = openDropdownId.value === id ? null : id
}

// Закрыть dropdown при клике вне его
const closeDropdown = () => {
    openDropdownId.value = null
}
const columns: TableColumn<Service>[] = [
    {
        accessorKey: 'name',
        header: 'Сервис',
        cell: ({ row }) => h('div', { class: 'font-medium' }, row.getValue('name'))
    },
    {
        accessorKey: 'login',
        header: 'Логин/Email',
        cell: ({ row }) => {
            const login = row.getValue('login') as string
            const isEmail = login.includes('@')

            return h('div', {
                class: `flex items-center gap-2 ${isEmail ? 'text-gray-600' : 'text-gray-900'}`
            }, [
                isEmail ? h('i', { class: 'i-lucide-mail w-4 h-4' }) : h('i', { class: 'i-lucide-user w-4 h-4' }),
                login
            ])
        }
    },
    {
        accessorKey: 'tag',
        header: 'Тег',
        cell: ({ row }) => {
            const tag = row.getValue('tag') as string
            const colorMap: Record<string, string> = {
                backend: 'blue',
                microservice: 'green',
                database: 'purple',
                cache: 'orange',
                storage: 'cyan'
            }

            const color = colorMap[tag] || 'purple'

            return h(UBadge, {
                class: 'capitalize',
                variant: 'soft',
                color: color as any
            }, () => tag)
        }
    },
    {
        accessorKey: 'updatedAt',
        header: 'Обновлено',
        cell: ({ row }) => {
            return new Date(row.getValue('updatedAt')).toLocaleString('ru-RU', {
                day: 'numeric',
                month: 'short',
                hour: '2-digit',
                minute: '2-digit',
                hour12: false
            })
        }
    },
    {
        id: 'actions',
        cell: ({ row }) => {
            const isOpen = openDropdownId.value === row.original.id

            return h(
                'div',
                {
                    class: 'text-right relative',
                    onClick: (e: Event) => e.stopPropagation()
                },
                [
                    h(UButton, {
                        icon: 'i-lucide-ellipsis-vertical',
                        color: 'neutral',
                        variant: 'ghost',
                        class: 'ml-auto',
                        'aria-label': 'Открыть меню действий',
                        onClick: () => toggleDropdown(row.original.id)
                    }),
                    isOpen && h(
                        'div',
                        {
                            class: 'absolute right-0 top-full mt-1 w-48 bg-white rounded-md shadow-lg border border-gray-200 z-50 py-1',
                        },
                        getRowItems(row).map((item, index) => {
                            if (item.type === 'separator') {
                                return h('div', { class: 'border-t border-gray-200 my-1', key: index })
                            }

                            if (item.type === 'label') {
                                return h('div', {
                                    class: 'px-3 py-2 text-xs font-semibold text-gray-500 uppercase tracking-wide',
                                    key: index
                                }, item.label)
                            }

                            return h(
                                'button',
                                {
                                    class: 'w-full text-left px-3 py-2 text-sm text-gray-700 hover:bg-gray-100 flex items-center gap-2',
                                    onClick: () => {
                                        item.onSelect?.()
                                        closeDropdown()
                                    },
                                    key: index
                                },
                                [
                                    item.icon && h('i', { class: `${item.icon} w-4 h-4` }),
                                    item.label
                                ]
                            )
                        })
                    )
                ]
            )
        }
    }
]

// Обновите функцию getRowItems чтобы убрать type из обычных items
function getRowItems(row: Row<Service>) {
    return [
        {
            label: 'Скопировать логин',
            icon: 'i-lucide-copy',
            onSelect() {
                copy(row.original.login)
                toast.add({
                    title: 'Логин скопирован!',
                    color: 'success',
                    icon: 'i-lucide-circle-check'
                })
            }
        },
        {
            label: 'Скопировать пароль',
            icon: 'i-lucide-copy',
            onSelect() {
                copy(row.original.password)
                toast.add({
                    title: 'Пароль скопирован!',
                    color: 'success',
                    icon: 'i-lucide-circle-check'
                })
            }
        },
        {
            type: 'separator'
        },
        {
            label: 'Перейти к сервису',
            icon: 'i-lucide-external-link',
            onSelect() { /* логика перехода */ }
        },
        {
            label: 'Просмотреть логи',
            icon: 'i-lucide-file-text',
            onSelect() { /* логика просмотра логов */ }
        },
        {
            label: row.original.status === 'active' ? 'Остановить сервис' : 'Запустить сервис',
            icon: row.original.status === 'active' ? 'i-lucide-square' : 'i-lucide-play',
            onSelect() { /* логика управления сервисом */ }
        }
    ]
}
</script>

<template>

    <div class="flex justify-center gap-4 mb-4 mt-4">
        <UButton to="/add/service" variant="solid" color="neutral">Добавить сервис</UButton>
        <UButton variant="solid" color="primary">Импорт/Экспорт</UButton>
    </div>

    <div @click="closeDropdown">
        <UTable :data="data" :columns="columns" class="flex-1" />
    </div>
</template>
<style>

</style>
