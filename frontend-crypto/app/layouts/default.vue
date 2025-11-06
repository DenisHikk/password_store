<template>
    <div class="flex min-h-dvh flex-col">
        <HeaderComponent />
        <main class="flex-1">
            <slot />
        </main>
        <FooterComponent class="mt-auto"/>
    </div>
</template>