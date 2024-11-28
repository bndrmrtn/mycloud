<template>
  <BtnRoot target="_blank" class="from-gray-50 to-gray-50 text-black">
    <span class="flex items-center justify-center space-x-3">
      <GoogleIcon />
      <span>Login</span>
    </span>
  </BtnRoot>
</template>

<script setup lang="ts">
import GoogleIcon from '~/components/icons/google-icon.vue';
import BtnRoot from '~/components/betterbuttons/btn-root.vue';
</script>
