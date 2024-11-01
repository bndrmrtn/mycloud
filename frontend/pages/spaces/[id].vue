<script setup lang="ts">
import {useRoute, useRouter} from "#app";
import type {SpaceFile} from "~/types/space";
import {onMounted, watch} from "vue";
import prettyBytes from "pretty-bytes";
import FileIcon from "~/components/icons/file-icon.vue";
import DotsHandleIcon from "~/components/icons/dots-handle-icon.vue";
import FolderIcon from "~/components/icons/folder-icon.vue";
import {apiFetch} from "~/scripts/request";

const dirs = ref<Array<string>>([])
const files = ref<Array<SpaceFile>>([])

const route = useRoute()
const router = useRouter()
const id = route.params.id as string

const fetchDirs = async () => {
  try {
    const res = await apiFetch(`/spaces/${id}/fs?path=${route.query['dir'] || '/'}`)
    if(res.status != 200) return router.push('/')
    dirs.value = await res.json()
  } catch (e: unknown) {
    console.error(e)
  }
}

const fetchFiles = async () => {
  try {
    const res = await apiFetch(`/spaces/${id}/files?path=${route.query['dir'] || '/'}`)
    if(res.status != 200) return router.push('/')
    files.value = await res.json()
  } catch (e: unknown) {
    console.error(e)
  }
}

const dirRoute = (d: string): string => {
  const path = route.fullPath.split('?')[0]
  const searchParams = new URLSearchParams(route.fullPath.split('?')[1] || "")
  const currentDir = searchParams.get('dir')
  if(!currentDir) searchParams.set('dir', '/' + d)
  else searchParams.set('dir', currentDir + '/' + d)
  return `${path}?${searchParams.toString()}`
}

onMounted(async () => {
  await fetchFiles()
  await fetchDirs()
})

watch(route, () => {
    dirs.value = []
    files.value = []
    fetchFiles()
    fetchDirs()
})
</script>

<template>
  <div class="px-10 py-5 max-w-screen-md mx-auto">
    <h1 class="fredoka text-3xl">Space:</h1>
    <p class="my-2 text-sm text-gray-700">Path: <span class="px-1 py-0.5 rounded bg-gray-100">{{ route.query['dir'] || '/' }}</span></p>
    <ul class="border rounded-lg mt-5">
      <li v-if="route.query['dir']" class="flex items-center space-x-3 py-2 px-4 hover:bg-gray-100 transition rounded-t-lg border-b">
        <folder-icon class="text-gray-600" />
        <RouterLink class="w-full" :to="`/spaces/${route.params?.id}`">..</RouterLink>
      </li>
      <li
          class="flex items-center justify-between py-2 px-4 hover:bg-gray-100 transition"
          v-for="(dir, inx) in dirs" :key="inx"
          :class="{'border-b': inx != dirs.length, 'rounded-t-lg': inx == 0}"
      >
        <!-- File name and icon -->
        <div class="flex items-center space-x-3">
          <folder-icon class="text-gray-600" />
          <RouterLink class="transition hover:underline hover:text-blue-400" :to="dirRoute(dir)">{{ dir }}</RouterLink>
        </div>
        <!-- edit icon -->
        <button>
          <dots-handle-icon />
        </button>
      </li>
      <li
          class="flex items-center justify-between py-2 px-4 hover:bg-gray-100 transition"
          :class="{'border-b': inx != files.length - 1, 'rounded-t-lg': inx == 0 && dirs.length == 0, 'rounded-b-lg': inx == files.length - 1}"
          v-for="(file, inx) in files" :key="file"
      >
        <!-- File name and icon -->
        <div class="flex items-center space-x-3">
          <file-icon class="text-gray-600" />
          <p>{{ file.file_name }}</p>
        </div>
        <!-- File size and edit icon -->
        <div class="flex items-center space-x-2">
          <p class="text-gray-700 text-sm">{{ prettyBytes(file.fileinfo.info.size) }}</p>
          <button>
            <dots-handle-icon />
          </button>
        </div>
      </li>
    </ul>
  </div>
</template>