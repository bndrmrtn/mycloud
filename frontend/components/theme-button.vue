<template>
  <button :disabled="loading" class="transition w-full py-2 px-3 rounded-lg bg-gray-100 hover:bg-gray-200 focus:bg-gray-200/10 border border-gray-200 text-gray-500 disabled:text-gray-400 fredoka">
    <slot/>
    <spinner-icon v-if="loading" />
  </button>
</template>

<script setup lang="ts">
import SpinnerIcon from "~/components/icons/spinner-icon.vue";

defineProps<{
  loading: boolean
}>()
</script>