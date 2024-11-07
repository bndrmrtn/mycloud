<template>
  <BtnRoot :to="to" :is-loading="isLoading" class="from-emerald-400 via-green-400 to-lime-200">
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