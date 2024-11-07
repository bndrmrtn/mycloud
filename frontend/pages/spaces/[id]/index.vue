<script setup lang="ts">
import {onMounted, useLoaderStore} from "#imports";
import type {Space, SpaceFile} from "~/types/space";
import {useRoute, useRouter} from "#app";
import {useToast} from "vue-toastification";
import SpaceLayout from "~/layouts/space.vue";
import ReloadIcon from "~/components/icons/reload-icon.vue";
import DownloadIcon from "~/components/icons/download-icon.vue";
import {fetchDirs, fetchFiles, fetchSpace, requestSpaceDownload} from "~/scripts/space";

definePageMeta({
  middleware: ['space', 'auth'],
})

const loader = useLoaderStore()
const route = useRoute()
const router = useRouter()

const id = route.params.id as string
const space = ref<Space|null>(null)
const dirs = ref<Array<string>>([])
const files = ref<Array<SpaceFile>>([])
const fileModal = ref<SpaceFile | null>(null)
const downloading = ref(false)

const downloadDir = async () => {
  downloading.value = true
  const err = await requestSpaceDownload(id)
  downloading.value = false

  if(!process.client) return
  const toast = useToast()

  if(err instanceof Error) {
    toast.error(err.message)
    return
  }

  toast.info('Download requested successfully')
}

const fetchSpaceName = async () => {
  space.value = await fetchSpace(id)
  if(!space.value) {
    if(process.client) useToast().warning('Failed to fetch space')
    await router.push('/spaces')
  }
}

const dir = () => route.query['directory'] as string || '/'

const fetchSpaceDirs = async () => {
  const d = await fetchDirs(id, dir())
  if(d) return dirs.value = d
  if(process.client) useToast().warning('Failed to fetch directories')
}


const fetchSpaceFiles = async () => {
  const f = await fetchFiles(id, dir())
  if(f != null) return files.value = f
  if(process.client) useToast().warning('Failed to fetch files')
}

onMounted(async () => {
  await fetchSpaceName()
  await load()
})

const load = async (reload: boolean = false) => {
  loader.start()
  await fetchSpaceFiles()
  await fetchSpaceDirs()
  loader.finish()
  if(reload) if(process.client) useToast().info('Reloaded successfully')
}

const openFileModal = (id: string) => {
  for(let i = 0; i < files.value.length; i++) {
    if(files.value[i].id == id) {
      fileModal.value = files.value[i]
      break
    }
  }
}

watch(route, async () => await load())
</script>

<template>
  <SpaceLayout>
      <div class="flex items-center justify-between">
        <h1 class="fredoka text-3xl mb-5">{{ space?.name }}</h1>
        <div class="flex items-center space-x-2">
          <buttons-button-pinkle :to="`/spaces/${id}/upload`" class="!w-min">Upload</buttons-button-pinkle>
          <buttons-button-pinkle @click="downloadDir" :is-loading="downloading" class="flex items-center justify-center">
            <download-icon class="-mt-1" />
          </buttons-button-pinkle>
          <buttons-button-pinkle @click="load(true)" class="flex items-center justify-center">
            <reload-icon class="-mt-1" />
          </buttons-button-pinkle>
        </div>
      </div>

    <ul class="mt-5 bg-widget bg-blur rounded-lg drop-shadow-sm">
        <FileListBack :id="id" />
        <FileList :id="id" v-for="dir in dirs" :key="dir" type="directory" :name="dir" />
        <FileList @openModal="openFileModal" v-for="file in files" :key="file.id" type="file" :id="file.id" :name="file.file_name" :size="file.fileinfo.info.size" :file-type="file.fileinfo.info.type" />
    </ul>

    <div v-if="dirs.length == 0 && files.length == 0">
      <p>You have no files currently.</p>
    </div>
  </SpaceLayout>

  <ModalsFileModal v-if="fileModal != null" :file="fileModal" @close="fileModal = null" />
</template>