<template>
  <BtnRoot :is-loading="false" target="_blank" to="https://discord.gg/" class="from-[#5865F2] to-[#5865F2] mt-2 md:mt-0">
    <span class="flex items-center justify-center space-x-3">
      <IconsDiscordIcon/>
      <span>{{ title }}</span>
    </span>
  </BtnRoot>
</template>

<script setup lang="ts">
import BtnRoot from "~/components/betterbuttons/btn-root.vue";

defineProps<{
  title: string
}>()
</script>