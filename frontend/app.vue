<template>
    <NuxtRouteAnnouncer />
    <Navbar/>
    <main class="min-h-[calc(100vh-var(--nav-height))] mt-[var(--nav-height)] relative w-full">
      <NuxtPage />
    </main>
</template>

<script setup lang="ts">
import {onMounted} from "vue";
import {useAuthStore} from "~/stores/auth";
import {useRoute} from "#app";

useHead({
  title: ''
})

const route = useRoute()
const auth = useAuthStore()

onMounted(async () => {
  await auth.handle()
})

watch(() => route.fullPath, () => auth.handle().then(() => console.log('Checking auth state')))
</script>