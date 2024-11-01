<script setup lang="ts">
import {useRoute} from "#app";
import {apiFetch} from "~/scripts/request";

const route = useRoute()
const id = route.params.id as string

const dir = ref('')
const fileName = ref('')
const fileData = ref<File|null>(null)

const uploadFile = (e: Event) => {
  const input = e.target as HTMLInputElement
  const files = input.files
  if (!files || files.length < 1) {
    return
  }
  fileData.value = files[0]
}

const processing = ref(false)
const upload = async () => {
  if(!fileData.value) return
  processing.value = true
  const data = new FormData()
  data.append('file', fileData.value)
  if(dir.value != '') data.append('directory', dir.value)
  if(fileName.value) data.append('filename', fileName.value)

  try {
    await apiFetch(`/spaces/${id}/upload`, {
      method: 'POST',
      body: data,
    })
    alert('Successfully uploaded')
  } catch (e: unknown) {
    console.error(e)
  } finally {
    processing.value = false
    fileData.value = null
    fileName.value = ''
    dir.value = ''
  }
}
</script>

<template>
  <div class="px-10 py-5 max-w-screen-md mx-auto">
    <div class="flex items-center justify-between">
      <h1 class="fredoka text-3xl mb-5">Upload</h1>
      <green-link :to="`/spaces/${id}`">Back</green-link>
    </div>
    <div>
      <form @submit.prevent>
        <input :disabled="processing" v-model.lazy="dir" type="text" placeholder="Enter directory (default is /)" class="w-full bg-gray-200 text-gray-700 rounded-lg p-3">
        <input :disabled="processing" v-model.lazy="fileName" type="text" placeholder="Rename file (empty for default)" class="w-full bg-gray-200 text-gray-700 rounded-lg p-3 mt-4">
        <input :disabled="processing" @change="uploadFile" type="file" placeholder="Upload file" class="w-full bg-gray-200 text-gray-700 rounded-lg p-3 mt-4">
        <green-button :loading="processing" @click="upload" class="mt-4">Upload</green-button>
      </form>
    </div>
  </div>
</template>