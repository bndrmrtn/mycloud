<script setup lang="ts">
import type { SpaceFile } from '~/types/space';
import InputUtil from '~/components/utils/input-util.vue';
import ButtonBluish from '~/components/buttons/button-bluish.vue';
import ButtonGreenish from '~/components/buttons/button-greenish.vue';
import { updateFileInfo } from '~/scripts/space';
import { useToast } from 'vue-toastification';

const emit = defineEmits(['back']);

const props = defineProps<{
  file: SpaceFile;
}>();

const name = ref(props.file.file_name);
const directory = ref(props.file.directory);
const processing = ref(false);

const save = async () => {
  processing.value = true;
  const err = await updateFileInfo(props.file.id, name.value, directory.value);
  processing.value = false;

  if (!process.client) return;
  const toast = useToast();

  if (err instanceof Error) {
    toast.error(err.message);
    return;
  }

  toast.success('Successfully saved');
};
</script>

<template>
  <form class="my-2 text-left" @submit.prevent>
    <label for="filename" class="text-sm text-gray-300 ml-2">File name</label>
    <InputUtil
      :disabled="processing"
      id="filename"
      placeholder="File name"
      v-model="name"
      :value="name"
    />
    <label for="directory" class="text-sm text-gray-300 ml-2 mt-3"
      >Directory</label
    >
    <InputUtil
      :disabled="processing"
      id="directory"
      placeholder="Directory"
      v-model="directory"
      :value="directory"
    />
    <ButtonGreenish class="mt-3" @click="save" :is-loading="processing"
      >Save</ButtonGreenish
    >
  </form>
  <ButtonBluish class="mt-1" @click="emit('back')">Back</ButtonBluish>
</template>
