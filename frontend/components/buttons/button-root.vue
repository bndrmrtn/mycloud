<template>
  <button
      class="fredoka transition-all hover:opacity-80 focus:opacity-70 bg-gradient-to-r duration-150 w-full px-2 py-2 md:px-5 rounded-xl disabled:cursor-not-allowed"
      :class="{'flex items-center space-x-2 text-center content-center justify-center': isLoading}"
      :disabled="isLoading"
  >
    <span>
      <slot/>
    </span>
    <SpinnerIcon v-if="isLoading" />
  </button>
</template>

<script setup lang="ts">
import SpinnerIcon from "~/components/icons/spinner-icon.vue";

defineProps<{
  isLoading?: boolean
}>()
</script>