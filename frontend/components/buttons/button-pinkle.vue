<template>
  <BtnRoot :to="to" :is-loading="isLoading" class="from-pink-400 via-magenta-400 to-violet-500">
    <slot/>
  </BtnRoot>
</template>

<script setup lang="ts">
import BtnRoot from "~/components/betterbuttons/btn-root.vue";

defineProps<{
  to?: string
  isLoading?: boolean
}>()
</script>