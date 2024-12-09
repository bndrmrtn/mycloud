<script setup lang="ts">
import SpaceLayout from '~/layouts/space.vue';
import {definePageMeta} from "#imports";

definePageMeta({
  middleware: ['auth'],
});
</script>

<template>
  <SpaceLayout>
    hello
  </SpaceLayout>
</template>