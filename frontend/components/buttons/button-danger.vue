<template>
  <BtnRoot :to="to" :is-loading="isLoading" class="from-rose-600 to-pink-600">
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