<script setup lang="ts">
import type {User} from "~/types/user";

const env = useRuntimeConfig();

defineProps<{
  user: User
}>()
</script>

<template>
  <img
      class="w-7 h-7 md:w-8 md:h-8 rounded-lg"
      :src="`${env.public.api}/profileimage/${user.image_url}`"
      :alt="user.name"
  />
</template>