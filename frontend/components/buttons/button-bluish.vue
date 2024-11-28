<template>
  <BtnRoot
    :to="to"
    :is-loading="isLoading"
    class="from-indigo-400 via-violet-500 to-blue-400"
  >
    <slot />
  </BtnRoot>
</template>

<script setup lang="ts">
import BtnRoot from '~/components/betterbuttons/btn-root.vue';

defineProps<{
  to?: string;
  isLoading?: boolean;
}>();
</script>
