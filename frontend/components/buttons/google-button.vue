<template>
  <ButtonRootLink target="_blank" class="from-gray-50 to-gray-50 text-black">
          <span class="flex items-center justify-center space-x-3">
              <GoogleIcon/>
              <span>Login</span>
          </span>
  </ButtonRootLink>
</template>

<script setup lang="ts">
import ButtonRootLink from "~/components/buttons/button-root-link.vue";
import GoogleIcon from "~/components/icons/google-icon.vue";
</script>
