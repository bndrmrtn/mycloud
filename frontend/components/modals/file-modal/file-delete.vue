<template>
  <div class="mt-1.5" :class="{ 'flex items-center space-x-2': start }">
    <ButtonDanger v-if="!start" @click="start = true">Delete</ButtonDanger>

    <ButtonGreenish v-if="start" @click="start = false">Retain</ButtonGreenish>
    <ButtonDanger
      :is-loading="deleting"
      @click="del"
      v-tooltip="'You won\'t be able to withdraw it'"
      v-if="start"
      >Delete</ButtonDanger
    >
  </div>
</template>

<script setup lang="ts">
import ButtonDanger from '~/components/buttons/button-danger.vue';
import ButtonGreenish from '~/components/buttons/button-greenish.vue';
import { deleteFile } from '~/scripts/space';
import { useToast } from 'vue-toastification';

const props = defineProps<{
  file_id: string;
}>();

const emit = defineEmits(['deleted']);
const start = ref(false);
const deleting = ref(false);

const del = async () => {
  deleting.value = true;
  const err = await deleteFile(props.file_id);
  deleting.value = false;

  if (err instanceof Error) {
    if (process.client) useToast().error(err.message);
    return;
  }

  if (process.client) useToast().success('File deleted successfully');
  emit('deleted');
};
</script>
