<script setup lang="ts">
import {onMounted, useLoaderStore} from "#imports";
import type {Space} from "~/types/space";
import {useToast} from "vue-toastification";
import {fetchSpaces} from "~/scripts/fetch-spaces";
import {useRouter} from "#app";
import prettyBytes from "pretty-bytes";
import SpaceLayout from "~/layouts/space.vue";
import CreateSpaceModal from "~/components/modals/create-space-modal.vue";

definePageMeta({
  middleware: ['auth'],
})

const router = useRouter()
const spaces = ref<Array<Space>>([])
const createModal = ref(false)

const fetchData = async () => {
  const data = await fetchSpaces()
  if(data) return spaces.value = data

  if(process.client) useToast().warning('Failed to load your spaces.')
  await router.push('/')
}

const addSpace = (s: Space) => {
  spaces.value.push(s)
}

onMounted(async () => {
  await fetchData()
  useLoaderStore().finish()
})
</script>

<template>
  <SpaceLayout>
      <div class="flex items-center justify-between">
        <h1 class="fredoka text-3xl mb-5">Spaces</h1>
        <buttons-button-pinkle @click="createModal = true" class="!w-min">Create</buttons-button-pinkle>
      </div>

      <ul class="mt-5">
        <li
            class="w-full bg-widget bg-blur py-3 px-4 rounded-lg drop-shadow-sm flex items-center justify-between mb-3"
            v-for="space in spaces" :key="space.id"
        >
          <div>
            <h2 class="fredoka">{{ space.name }}</h2>
            <p class="text-gray-400">Size: {{ prettyBytes(space.size) }}</p>
          </div>
          <buttons-button-bluish :to="`/spaces/${space.id}`" class="!w-min">Open</buttons-button-bluish>
        </li>
      </ul>
  </SpaceLayout>

  <CreateSpaceModal v-if="createModal" @close="createModal = false" @finish="addSpace" />
</template>