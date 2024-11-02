<template>
     <ButtonRootLink target="_blank" to="https://github.com/bndrmrtn/mycloud" class="from-black to-black">
          <span class="flex items-center justify-center space-x-3">
              <IconsGithubIcon/>
              <span>{{ 'Open-source repository' }}</span>
          </span>
     </ButtonRootLink>
</template>

<script setup lang="ts">
import ButtonRootLink from "~/components/buttons/button-root-link.vue";
</script>