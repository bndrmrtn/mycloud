<script setup lang="ts">
import SpaceLayout from "~/layouts/space.vue";
import {onMounted, useLoaderStore} from "#imports";
import {useRoute, useRouter} from "#app";
import type {Space} from "~/types/space";
import {fetchSpace} from "~/scripts/fetch-spaces";
import {useToast} from "vue-toastification";
import InputUtil from "~/components/utils/input-util.vue";
import {createFile} from "~/scripts/create-space";
import ButtonBluish from "~/components/buttons/button-bluish.vue";

definePageMeta({
  middleware: ['space', 'auth'],
})

const route = useRoute()
const router = useRouter()
const loader = useLoaderStore()

const id = route.params.id as string
const space = ref<Space | null>(null)

const fetchSpaceName = async () => {
  space.value = await fetchSpace(id)
  if(!space.value) {
    if(process.client) useToast().warning('Failed to fetch space')
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
  const err = await createFile(id, fileData.value, filePath.value, fileName.value)
  processing.value = false

  if(err == null) return process.client ? useToast().success('Successfully uploaded') : null

  if(process.client) useToast().error(err.message)
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
        <InputUtil :disabled="processing" v-model.lazy="filePath" placeholder="Path (default: /)" />
        <InputUtil :disabled="processing" v-model.lazy="fileName" placeholder="New file name" class="mt-3" />

        <ButtonBluish :disabled="processing" @click="(uploadFileEl as HTMLInputElement).click()" type="button" class="mt-3 md:py-2.5">
          <span v-if="!fileData">Choose file</span>
          <span v-else>{{ fileData?.name }}</span>
        </ButtonBluish>

        <buttons-button-pinkle :isLoading="processing" @click="upload" class="mt-4 !py-2.5">Submit</buttons-button-pinkle>
        <input ref="uploadFileEl" type="file" class="hidden" @change="uploadFile">
      </form>
    </div>
  </SpaceLayout>
</template>
