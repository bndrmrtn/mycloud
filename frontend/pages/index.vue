<template>
  <div class="min-h-[calc(100vh-var(--nav-height))] flex items-center justify-center w-full">
    <div>
      <div class="flex items-start space-x-3">
        <div>
          <h1 class="text-4xl md:text-5xl fredoka text-gray-800">MyCloud</h1>
          <p class="mt-2 text-gray-500">Easy to use web based<br/>file manager</p>
        </div>
        <img src="../assets/images/cloud.svg" class="w-16 svg-img" alt="Cloud Image">
      </div>
      <theme-button v-if="!auth.isLoggedIn" class="mt-3" :loading="progress" @click="request">
        Login
      </theme-button>
      <theme-link v-else to="/spaces" :loading="false">
        Spaces
      </theme-link>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useAuthStore} from "#imports";
import ThemeLink from "~/components/theme-link.vue";
import {apiFetch} from "~/scripts/request";

const progress = ref(false)

const auth = useAuthStore()

const request = async () => {
  progress.value = true
  try {
    const res = await apiFetch('/auth-redirect')
    const data = await res.json() as {redirect_url: string}
    window.location.href = data.redirect_url
  } catch(e: unknown) {
    console.error(e)
    alert('Failed to create authentication url')
  } finally {
    progress.value = false
  }
}
</script>