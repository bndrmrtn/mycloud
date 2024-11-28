<template>
  <div class="text-left my-2">
    <template v-for="(val, key) in info" :key="key">
      <div class="my-2">
        <p class="ml-1 text-sm text-gray-400">{{ key }}:</p>
        <div
          class="transition group hover:cursor-pointer w-full px-3 py-1.5 rounded-lg bg-main-from hover:bg-main-to flex items-center justify-between"
        >
          <p class="text-nowrap overflow-x-scroll">{{ val }}</p>
          <button
            @click="copyText(val)"
            v-tooltip="'Copy content'"
            class="transition hidden group-hover:flex items-center justify-center w-6 h-6 bg-widget p-1 rounded-lg"
          >
            <ClipboardIcon />
          </button>
        </div>
      </div>
    </template>
    <ButtonBluish @click="emit('back')" class="my-1.5">Back</ButtonBluish>
  </div>
</template>

<script setup lang="ts">
import ClipboardIcon from '~/components/icons/clipboard-icon.vue';
import { useToast } from 'vue-toastification';
import ButtonBluish from '~/components/buttons/button-bluish.vue';

defineProps<{
  info: { [key: string]: string };
}>();

const emit = defineEmits(['back']);

const copyText = (text: string) => {
  navigator.clipboard
    .writeText(text)
    .then(() =>
      process.client
        ? useToast().info('Successfully copied to clipboard')
        : null
    )
    .catch(() =>
      process.client ? useToast().warning('Failed to copy to clipboard') : null
    );
};
</script>
