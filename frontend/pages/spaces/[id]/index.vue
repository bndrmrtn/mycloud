<script setup lang="ts">
import {onMounted, useLoaderStore} from "#imports";
import type {Space, SpaceFile} from "~/types/space";
import {dirRoute, fetchDirs, fetchFiles, fetchSpace} from "~/scripts/fetch-spaces";
import {useRoute, useRouter} from "#app";
import {useToast} from "vue-toastification";
import SpaceLayout from "~/layouts/space-layout.vue";
import ReloadIcon from "~/components/icons/reload-icon.vue";

definePageMeta({
  middleware: ['auth'],
})

const loader = useLoaderStore()
const route = useRoute()
const router = useRouter()
const { warning, info } = useToast()

const id = route.params.id as string
const space = ref<Space|null>(null)
const dirs = ref<Array<string>>([])
const files = ref<Array<SpaceFile>>([])

const fetchSpaceName = async () => {
  space.value = await fetchSpace(id)
  if(!space.value) {
    warning('Failed to fetch space')
    await router.push('/spaces')
  }
}

const dir = () => route.query['directory'] as string || '/'

const fetchSpaceDirs = async () => {
  const d = await fetchDirs(id, dir())
  if(d) return dirs.value = d
  warning('Failed to fetch directories')
}


const fetchSpaceFiles = async () => {
  const f = await fetchFiles(id, dir())
  if(f != null) return files.value = f
  warning('Failed to fetch files')
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
  if(reload) info('Reloaded successfully')
}

watch(route, async () => await load())
</script>

<template>
  <SpaceLayout>
      <div class="flex items-center justify-between">
        <h1 class="fredoka text-3xl mb-5">{{ space?.name }}</h1>
        <div class="flex items-center space-x-2">
          <buttons-button-pinkle :to="`/spaces/${id}/upload`" class="!w-min">Upload</buttons-button-pinkle>
          <buttons-button-pinkle @click="load(true)" class="flex items-center justify-center">
            <reload-icon class="-mt-1" />
          </buttons-button-pinkle>
        </div>
      </div>

      <ul class="mt-5 bg-widget rounded-lg drop-shadow-sm">
        <file-list-back :id="id" />
        <FileList :id="id" v-for="dir in dirs" :key="dir" type="directory" :name="dir" />
        <FileList :id="id" v-for="file in files" :key="file.id" type="file" :name="file.file_name" :size="file.fileinfo.info.size" />
      </ul>
  </SpaceLayout>
</template>