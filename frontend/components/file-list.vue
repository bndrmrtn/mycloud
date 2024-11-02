<script setup lang="ts">
import FolderIcon from "~/components/icons/folder-icon.vue";
import FileIcon from "~/components/icons/file-icon.vue";
import prettyBytes from "pretty-bytes";
import DotsHandleIcon from "~/components/icons/dots-handle-icon.vue";
import {useRoute} from "#app";

defineProps<{
  id: string
  type: 'file' | 'directory'
  name: string
  size?: number
}>()

const route = useRoute()

const calcPath = (d: string) => {
  const [path, query] = route.fullPath.split('?')
  const searchParams = new URLSearchParams(query || "")
  const currentDir = searchParams.get('directory') || ""
  searchParams.set('directory', `${currentDir}/${d}`.replace(/^\/+/, '/'))
  return `${path}?${searchParams.toString()}`
}
</script>

<template>
  <li class="py-2 px-4 first:rounded-t-lg last:rounded-b-lg border-b border-gray-500 last:border-b-0 flex items-center justify-between">
    <!-- File/Directory name and icon -->
    <div class="flex items-center space-x-3">
      <FolderIcon v-if="type == 'directory'" class="text-gray-300" />
      <FileIcon v-if="type == 'file'" class="text-gray-300" />
      <p class="text-wrap max-w-fit" v-if="type == 'file'">
        {{ name }}
      </p>
      <RouterLink class="transition hover:text-blue-400 hover:underline" :to="calcPath(name)" v-else>
        {{ name }}
      </RouterLink>
    </div>

    <!-- File size and edit icon -->
    <div class="flex items-center space-x-2" v-if="type == 'file'">
      <p class="text-gray-400 text-sm">{{ prettyBytes(size) }}</p>
      <button>
        <DotsHandleIcon/>
      </button>
    </div>
  </li>
</template>
