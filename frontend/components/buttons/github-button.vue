<template>
     <BtnRoot target="_blank" to="https://github.com/bndrmrtn/mycloud" class="from-black to-black">
          <span class="flex items-center justify-center space-x-3">
              <IconsGithubIcon/>
              <span>{{ 'Open-source repository' }}</span>
          </span>
     </BtnRoot>
</template>

<script setup lang="ts">
import BtnRoot from "~/components/betterbuttons/btn-root.vue";
</script>