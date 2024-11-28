<script setup lang="ts">
defineProps<{
  request_id: string;
}>();

const emit = defineEmits(['move']);
</script>

<template>
  <div>
    <p>Your download prepared successfully.</p>
    <p>Request ID: {{ request_id }}</p>
    <button @click="emit('move')">Move</button>
  </div>
</template>
