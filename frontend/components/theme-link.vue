<template>
  <RouterLink :to="to" :disabled="loading" class="transition inline-block text-center w-full py-2 px-4 rounded-lg bg-gray-100 hover:bg-gray-200 focus:bg-gray-200/10 border border-gray-200 text-gray-500 disabled:text-gray-400 fredoka">
    <slot/>
    <spinner-icon v-if="loading" />
  </RouterLink>
</template>

<script setup lang="ts">
import SpinnerIcon from "~/components/icons/spinner-icon.vue";

defineProps<{
  loading: boolean
  to: string
}>()
</script>
