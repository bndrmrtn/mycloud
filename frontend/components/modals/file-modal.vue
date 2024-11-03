<template>
  <ModalUtil :show="true">
    <h2 class="text-2xl fredoka">Manage file</h2>
    <div class="text-left my-2">
      <template v-for="(val, key) in dataset()" :key="key">
        <div class="my-2">
          <p class="ml-1 text-sm text-gray-400">{{ key }}:</p>
          <div class="transition group hover:cursor-pointer w-full px-3 py-1.5 rounded-lg bg-main-from hover:bg-main-to flex items-center justify-between">
            <p class="text-nowrap overflow-x-scroll">{{ val }}</p>
            <button @click="copyText(val)" v-tooltip="'Copy content'" class="transition hidden group-hover:flex items-center justify-center w-6 h-6 bg-widget p-1 rounded-lg">
              <ClipboardIcon/>
            </button>
          </div>
        </div>
      </template>
    </div>
  </ModalUtil>
</template>

<script setup lang="ts">
import ModalUtil from "~/components/utils/modal-util.vue";
import type {SpaceFile} from "~/types/space";
import prettyBytes from "pretty-bytes";
import ClipboardIcon from "~/components/icons/clipboard-icon.vue";
import {useToast} from "vue-toastification";

const props = defineProps<{
  file: SpaceFile
}>()

const dataset = () => {
  return {
    'Name': props.file.file_name,
    'Directory': props.file.directory,
    'Size': prettyBytes(props.file.fileinfo.info.size),
    'SHA256': props.file.fileinfo.info.hash,
    'Mime type': props.file.fileinfo.info.type,
  }
}

const copyText = (text: string) => {
  navigator.clipboard.writeText(text).
    then(() => process.client ? useToast().info('Successfully copied to clipboard') : null).
    catch(() => process.client ? useToast().warning('Failed to copy to clipboard') : null)
}
</script>