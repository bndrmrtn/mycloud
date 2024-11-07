<template>
  <ModalUtil :show="true" @close="emit('close')">
    <h2 class="text-2xl fredoka">Manage file</h2>
    <file-information v-if="show == 'information'" :info="dataset()" @back="show = 'selector'" />
    <file-editor v-if="show == 'edit'" :file="file" @back="show = 'selector'" />

    <div class="text-left my-2" v-if="show == 'selector'">
      <ButtonBluish @click="show = 'information'" class="my-1.5">Information</ButtonBluish>
      <ButtonBluish @click="show = 'edit'" class="my-1.5">Edit information</ButtonBluish>
      <file-delete :file_id="file.id" @deleted="emit('close')" />
    </div>
  </ModalUtil>
</template>

<script setup lang="ts">
import ModalUtil from "~/components/utils/modal-util.vue";
import type {SpaceFile} from "~/types/space";
import prettyBytes from "pretty-bytes";
import ButtonBluish from "~/components/buttons/button-bluish.vue";
import FileInformation from "~/components/modals/file-modal/file-information.vue";
import {onMounted} from "vue";
import FileDelete from "~/components/modals/file-modal/file-delete.vue";
import FileEditor from "~/components/modals/file-modal/file-editor.vue";

const props = defineProps<{
  file: SpaceFile
}>()

const emit = defineEmits(['close'])
const show = ref('selector')

onMounted(() => show.value = 'selector')

const dataset = () => {
  return {
    'Name': props.file.file_name,
    'Directory': props.file.directory,
    'Size': prettyBytes(props.file.fileinfo.info.size),
    'SHA256': props.file.fileinfo.info.hash,
    'Mime type': props.file.fileinfo.info.type,
  }
}
</script>