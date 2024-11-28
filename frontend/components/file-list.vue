<script setup lang="ts">
import FolderIcon from '~/components/icons/folder-icon.vue';
import prettyBytes from 'pretty-bytes';
import DotsHandleIcon from '~/components/icons/dots-handle-icon.vue';
import { useRoute } from '#app';
import { getFileIcon } from '~/scripts/filetype';
import FileIcon from '~/components/icons/file-icon.vue';

defineProps<{
  id: string;
  type: 'file' | 'directory';
  name: string;
  size?: number;
  fileType?: string;
}>();

const env = useRuntimeConfig();
const emit = defineEmits(['openModal']);

const route = useRoute();

const calcPath = (d: string) => {
  const [path, query] = route.fullPath.split('?');
  const searchParams = new URLSearchParams(query || '');
  const currentDir = searchParams.get('directory') || '';
  searchParams.set('directory', `${currentDir}/${d}`.replace(/^\/+/, '/'));
  return `${path}?${searchParams.toString()}`;
};
</script>

<template>
  <li
    class="py-2 px-4 first:rounded-t-lg last:rounded-b-lg border-b border-gray-500 last:border-b-0 flex items-center"
  >
    <template v-if="type == 'directory'">
      <FolderIcon />
      <RouterLink
        class="ml-2 mr-2 transition hover:text-blue-400 hover:underline"
        :to="calcPath(name)"
      >
        {{ name }}
      </RouterLink>
    </template>

    <template v-else-if="type == 'file'">
      <component v-if="fileType" :is="getFileIcon(fileType)" />
      <FileIcon v-else />
      <a
        target="_blank"
        :href="`${env.public.api}/files/${id}/download`"
        class="truncate max-w-fit ml-2 mr-2 transition hover:text-blue-400 hover:underline"
        >{{ name }}</a
      >
    </template>

    <template v-if="type == 'file'">
      <p class="text-gray-400 text-sm ml-auto whitespace-nowrap">
        {{ prettyBytes(size || 0) }}
      </p>
      <button @click="emit('openModal', id)" v-tooltip="'Options'" class="ml-1">
        <DotsHandleIcon />
      </button>
    </template>
  </li>
</template>
