<template>
  <div class="w-full min-h-screen bg-square">
    <MiniNav v-if="!hideNav" />
    <div class="p-5 md:px-10 max-w-screen-md mx-auto">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  hideNav?: boolean;
}>();
</script>
