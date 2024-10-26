<script setup lang="ts">
import type {Space} from "~/types/space";
import {onMounted} from "vue";

const env = useRuntimeConfig()
const spaces = ref<Array<Space>>([])

onMounted(async () => {
  try {
    const res = await fetch(`${env.public.api}/spaces`)
    const data = await res.json()
    spaces.value = data as Array<Space>
  } catch (e: unknown) {
    console.error(e)
    alert('Failed to load your spaces')
  }
})
</script>

<template>
  <div class="px-10 py-5 max-w-screen-md mx-auto">
    <h1 class="fredoka text-3xl mb-5">Spaces</h1>

    <div v-for="(space, i) in spaces" :key="i" class="bg-gray-100 rounded border border-gray-200 py-2 px-4 mt-3 flex items-center justify-between">
      <div>
        <h2 class="fredoka">{{ i }} - {{ space.name }}</h2>
        <p>Size: {{ space.size }}mb</p>
        <p>ID: {{ space.id }}</p>
      </div>
      <green-link :to="`/space/${space.id}`">Open</green-link>
    </div>
  </div>
</template>