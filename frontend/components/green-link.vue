<template>
  <RouterLink :to="to" :disabled="loading" class="transition inline-block py-1.5 px-4 rounded-lg bg-theme border border-theme text-black disabled:text-gray-800 fredoka">
    <slot/>
    <spinner-icon v-if="loading" />
  </RouterLink>
</template>

<script setup lang="ts">
import SpinnerIcon from "~/components/icons/spinner-icon.vue";

defineProps<{
  to: string
  loading: boolean
}>()
</script>
