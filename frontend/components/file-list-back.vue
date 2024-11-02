<script setup lang="ts">
import FolderIcon from "~/components/icons/folder-icon.vue";
import {useRoute} from "#app";

defineProps<{
  id: string
}>()

const route = useRoute()

const dir = () => route.query['directory'] || '/'


const calcPath = () => {
  const [path, query] = route.fullPath.split('?')
  const searchParams = new URLSearchParams(query || "")
  const currentDir = searchParams.get('directory')

  if(currentDir && currentDir.includes('/')) {
    const parts = currentDir.split('/')
    if(parts.length > 1) {
      parts.pop()
      const partsS = parts.join('/')
      if(partsS) searchParams.set('directory', partsS)
      else searchParams.delete('directory')
    }
  }

  return `${path}?${searchParams.toString()}`
}

</script>

<template>
  <li v-if="dir() != '/'" class="py-2 px-4 first:rounded-t-lg last:rounded-b-lg border-b border-gray-500 last:border-b-0 flex items-center justify-between">
    <!-- File/Directory name and icon -->
    <div class="flex items-center space-x-3 w-full">
      <FolderIcon class="text-gray-300" />
      <RouterLink class="transition hover:text-blue-400 w-full block" :to="calcPath()">
        ..
      </RouterLink>
    </div>
  </li>
</template>
