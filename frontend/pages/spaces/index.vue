<script setup lang="ts">
import type {Space} from "~/types/space";
import {onMounted} from "vue";
import {apiFetch} from "~/scripts/request";

const spaces = ref<Array<Space>>([])

const load = async () => {
  spaces.value = []
  try {
    const res = await apiFetch('/spaces')
    const data = await res.json()
    spaces.value = data as Array<Space>
  } catch (e: unknown) {
    console.error(e)
    alert('Failed to load your spaces')
  }
}

onMounted(load)

const create = () => {
  const val = prompt('Enter your space name:')
  if(!val) return alert('No value')

  apiFetch('/spaces', {
    method: 'POST',
    body: JSON.stringify({name: val})
  }).then(load)
}
</script>

<template>
  <div class="px-10 py-5 max-w-screen-md mx-auto">
    <div class="flex items-center justify-between">
      <h1 class="fredoka text-3xl mb-5">Spaces</h1>
      <green-button @click="create">Create</green-button>
    </div>

    <div v-for="(space, i) in spaces" :key="i" class="bg-gray-100 rounded border border-gray-200 py-2 px-4 mt-3 flex items-center justify-between">
      <div>
        <h2 class="fredoka">{{ i }} - {{ space.name }}</h2>
        <p>Size: {{ space.size }}mb</p>
      </div>
      <green-link :to="`/spaces/${space.id}`">Open</green-link>
    </div>

    <div v-if="spaces.length == 0">
      <p>You have no spaces yet.</p>
    </div>
  </div>
</template>