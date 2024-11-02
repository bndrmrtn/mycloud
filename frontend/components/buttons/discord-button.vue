<template>
  <ButtonRootLink target="_blank" to="https://discord.gg/2BtaJker6S" class="from-[#5865F2] to-[#5865F2] mt-2 md:mt-0">
            <span class="flex items-center justify-center space-x-3">
              <IconsDiscordIcon/>
              <span>{{ title }}</span>
            </span>
  </ButtonRootLink>
</template>

<script setup lang="ts">
import ButtonRootLink from "~/components/buttons/button-root-link.vue";

defineProps<{
  title: string;
}>()
</script>