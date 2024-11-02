<script setup lang="ts">
import SpaceLayout from "~/layouts/space-layout.vue";
import {onMounted, useLoaderStore} from "#imports";
import {useRoute, useRouter} from "#app";
import type {Space} from "~/types/space";
import {fetchSpace} from "~/scripts/fetch-spaces";
import {useToast} from "vue-toastification";
import InputUtil from "~/components/utils/input-util.vue";
import {newRequest} from "~/scripts/request";

const route = useRoute()
const router = useRouter()
const loader = useLoaderStore()

const id = route.params.id as string
const space = ref<Space | null>(null)
const { warning, success } = useToast()

const fetchSpaceName = async () => {
  space.value = await fetchSpace(id)
  if(!space.value) {
    warning('Failed to fetch space')
    await router.push('/spaces')
  }
}

onMounted(async () => {
  await fetchSpaceName()
  loader.finish()
})

const filePath = ref('')
const fileName = ref('')
const fileData = ref<File|null>(null)

const uploadFileEl = ref(null)
const processing = ref(false)

const uploadFile = (e: Event) => {
  const input = e.target as HTMLInputElement
  const files = input.files
  if (!files || files.length < 1) {
    return
  }
  fileData.value = files[0]
}

const upload = async () => {
  if(!fileData.value) return
  processing.value = true
  const data = new FormData()
  data.append('file', fileData.value)
  if(filePath.value != '') data.append('directory', filePath.value)
  if(fileName.value) {
    if(!fileName.value.startsWith('/')) fileName.value = '/' + fileName.value
    data.append('filename', fileName.value)
  }


  try {
    await newRequest(`/spaces/${id}/upload`, {
      method: 'POST',
      body: data,
    })
    success('Successfully uploaded')
  } catch (e: unknown) {
    console.error(e)
    warning('Failed to upload file')
  } finally {
    processing.value = false
    fileData.value = null
    fileName.value = ''
    filePath.value = ''
  }
}
</script>

<template>
  <SpaceLayout>
    <div class="flex items-center justify-between">
        <h1 class="fredoka text-3xl mb-5">{{ space?.name }}</h1>
        <div class="flex items-center space-x-2">
          <buttons-button-pinkle :to="`/spaces/${id}`" class="!w-min">Files</buttons-button-pinkle>
        </div>
    </div>

    <div class="mt-5 max-w-sm mx-auto">
      <form @submit.prevent>
        <h2 class="fredoka text-xl mb-2">Upload file</h2>
        <InputUtil v-model.lazy="filePath" placeholder="Path (default: /)" />
        <InputUtil v-model.lazy="fileName" placeholder="New file name" class="mt-3" />
        <buttons-button-bluish @click="(uploadFileEl as HTMLInputElement).click()" type="button" class="mt-3 md:py-2.5">Choose file</buttons-button-bluish>
        <buttons-button-pinkle @click="upload" class="mt-4 !py-2.5">Submit</buttons-button-pinkle>
        <input ref="uploadFileEl" type="file" class="hidden" @change="uploadFile">
      </form>
    </div>
  </SpaceLayout>
</template>
