<script setup lang="ts">
import { onMounted, useLoaderStore } from '#imports';
import type { Space, SpaceFile } from '~/types/space';
import { useRoute, useRouter } from '#app';
import { useToast } from 'vue-toastification';
import SpaceLayout from '~/layouts/space.vue';
import ReloadIcon from '~/components/icons/reload-icon.vue';
import DownloadIcon from '~/components/icons/download-icon.vue';
import {
  fetchDirs,
  fetchFiles,
  fetchSpace,
  requestSpaceDownload,
} from '~/scripts/space';
import { useEventStore } from '~/stores/event';

definePageMeta({
  middleware: ['space', 'auth'],
});

const loader = useLoaderStore();
const route = useRoute();
const router = useRouter();

const id = route.params.id as string;
const space = ref<Space | null>(null);
const dirs = ref<Array<string>>([]);
const files = ref<Array<SpaceFile>>([]);
const fileModal = ref<SpaceFile | null>(null);
const downloading = ref(false);

const downloadDir = async () => {
  downloading.value = true;
  const err = await requestSpaceDownload(id);
  downloading.value = false;

  if (!process.client) return;

  if (err instanceof Error) {
    useToast().error(err.message);
    return;
  }

  useToast().info('Download requested successfully');
};

const fetchSpaceName = async () => {
  space.value = await fetchSpace(id);
  if (!space.value) {
    if (process.client) useToast().warning('Failed to fetch space');
    await router.push('/spaces');
  }
};

const dir = () => (route.query['directory'] as string) || '/';

const fetchSpaceDirs = async () => {
  const d = await fetchDirs(id, dir());
  if (d) return (dirs.value = d);
  if (process.client) useToast().warning('Failed to fetch directories');
};

const fetchSpaceFiles = async () => {
  const f = await fetchFiles(id, dir());
  if (f != null) return (files.value = f);
  if (process.client) useToast().warning('Failed to fetch files');
};

const event = useEventStore();

onMounted(async () => {
  const helpDir = (d: string) => {
    const currentDir = (dir() == '/' ? '' : dir()) + '/';
    if (!d.startsWith(currentDir)) return;
    const newDir = d.substring(currentDir.length).split('/')[0];
    if (newDir && !dirs.value.includes(newDir)) dirs.value.push(newDir);
  };

  event.register(
    'space_file_update',
    'update-fs',
    (data: Record<string, any>): void => {
      const fileData = (data as { file: SpaceFile }).file;
      helpDir(fileData.directory);
      // Check if the file is in the current directory
      if (fileData.directory != dir()) {
        files.value = files.value.filter((file) => file.id != fileData.id);
        return;
      }
      // Update the file
      const index = files.value.findIndex((file) => file.id === fileData.id);
      if (index != -1) files.value[index] = fileData;
    }
  );

  event.register(
    'space_file_delete',
    'delete-fs',
    (data: Record<string, any>): void => {
      // Delete the file from the list
      const id = data['file_id'] as string;
      files.value = files.value.filter((file) => file.id != id);
    }
  );

  event.register(
    'space_file_upload',
    'upload-fs',
    (data: Record<string, any>): void => {
      // Add the file to the list if it's in the current directory
      const fileData = (data as { file: SpaceFile }).file;
      helpDir(fileData.directory);
      if (fileData?.directory == dir()) files.value.push(fileData);
    }
  );

  await fetchSpaceName();
  await load();
});

onUnmounted(() => {
  event.detach('space_file_update', 'update-fs');
  event.detach('space_file_delete', 'delete-fs');
  event.detach('space_file_upload', 'upload-fs');
});

const load = async (reload: boolean = false) => {
  loader.start();
  await fetchSpaceFiles();
  await fetchSpaceDirs();
  loader.finish();
  if (reload) if (process.client) useToast().info('Reloaded successfully');
};

const openFileModal = (id: string) => {
  for (let i = 0; i < files.value.length; i++) {
    if (files.value[i].id == id) {
      fileModal.value = files.value[i];
      break;
    }
  }
};

watch(route, async () => await load());

const showFuncs = ref(false);
</script>

<template>
  <SpaceLayout>
    <div class="flex items-center justify-between">
      <h1 class="fredoka text-3xl mb-5 pr-1">{{ space?.name }}</h1>
      <div class="relative flex items-center space-x-2">
        <!-- Manage icon -->
        <buttons-button-pinkle
          v-tooltip="'Manage'"
          class="!w-min"
          @click="showFuncs = !showFuncs"
          :is-loading="downloading"
        >
          <IconsSettingBarsIcon class="-mt-1" />
        </buttons-button-pinkle>

        <buttons-button-pinkle
          v-tooltip="'Upload'"
          class="!w-min"
          :to="`/spaces/${id}/upload`"
        >
          <IconsUploadIcon class="-mt-1" />
        </buttons-button-pinkle>
      </div>
    </div>

    <ul class="mt-5 bg-widget bg-blur rounded-lg drop-shadow-sm">
      <FileListBack :id="id" />
      <FileList
        :id="id"
        v-for="dir in dirs"
        :key="dir"
        type="directory"
        :name="dir"
      />
      <FileList
        @openModal="openFileModal"
        v-for="file in files"
        :key="file.id"
        type="file"
        :id="file.id"
        :name="file.file_name"
        :size="file.fileinfo.info.size"
        :file-type="file.fileinfo.info.type"
      />
    </ul>

    <div v-if="dirs.length == 0 && files.length == 0">
      <p>You have no files currently.</p>
    </div>
  </SpaceLayout>

  <ModalsFileModal
    v-if="fileModal != null"
    :file="fileModal"
    @close="fileModal = null"
  />

  <ModalsSpaceModal
    v-if="showFuncs"
    :space="space"
    :dir="dir"
    @close="showFuncs = false"
    @download="downloadDir"
  />
</template>
