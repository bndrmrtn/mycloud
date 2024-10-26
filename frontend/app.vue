<template>
    <NuxtRouteAnnouncer />
    <Navbar/>
    <main class="min-h-[calc(100vh-var(--nav-height))] mt-[var(--nav-height)] relative w-full">
      <NuxtPage />
    </main>
</template>

<script setup lang="ts">
useHead({
})
</script>