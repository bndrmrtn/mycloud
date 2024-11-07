<script setup lang="ts">
import SpinnerIcon from "~/components/icons/spinner-icon.vue";

defineProps<{
  to?: string
  isLoading?: boolean
}>()
</script>

<template>
  <button v-if="!to" class="fredoka btnStyle" :class="{'btnLoading': isLoading}" :disabled="isLoading">
    <span><slot/></span>
    <SpinnerIcon v-if="isLoading" />
  </button>
  <RouterLink v-else-if="to && !to.substring(5, 9).includes('://')" :to="to" class="fredoka btnStyle inline-block">
    <slot/>
  </RouterLink>
  <a v-else-if="to" :href="to" class="fredoka btnStyle inline-block">
    <slot/>
  </a>
</template>

<style>
.btnStyle {
  @apply transition-all text-center hover:opacity-80 focus:opacity-70 bg-gradient-to-r duration-150 w-full px-2.5 py-2 md:px-5 rounded-xl disabled:cursor-not-allowed
}

.btnLoading {
  @apply flex items-center space-x-2 text-center content-center justify-center
}
</style>