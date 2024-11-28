<script setup lang="ts">
import type { NuxtError } from '#app';
import { onMounted, useLoaderStore } from '#imports';

const props = defineProps({
  error: Object as () => NuxtError,
});

const handleError = () => clearError({ redirect: '/' });

useHead({
  title: props.error?.statusCode + ' - ' + props.error?.statusMessage,
  bodyAttrs: {
    class:
      'bg-main-from bg-gradient-to-bl from-main-from to-main-to text-white',
  },
});

onMounted(() => useLoaderStore().finish());
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-square">
    <div class="px-5">
      <h1 class="text-center text-7xl drop-shadow text-green-300 fredoka">
        {{ error?.statusCode }}
      </h1>
      <p class="text-center text-gray-300 mt-2 text-lg">
        {{ error?.statusMessage }}
      </p>
      <p v-html="error?.stack"></p>
      <div class="mx-auto mt-4">
        <ButtonsButtonDanger @click="handleError">Home</ButtonsButtonDanger>
      </div>
    </div>
  </div>
</template>
