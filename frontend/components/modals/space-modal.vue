<template>
  <ModalUtil :show="true" @close="emit('close')">
    <h2 class="text-2xl fredoka mb-3">Manage space</h2>

    <ButtonsButtonPinkle class="my-1" @click="download">
      Download directory ({{ dir() }})
    </ButtonsButtonPinkle>

    <ButtonsButtonPinkle class="my-1">
      Manage collaborators
    </ButtonsButtonPinkle>
  </ModalUtil>
</template>
<script setup lang="ts">
import ModalUtil from '~/components/utils/modal-util.vue';

defineProps<{
  dir: () => string;
}>();

const emit = defineEmits(['close', 'download']);

const download = () => {
  emit('download');
  emit('close');
};
</script>
