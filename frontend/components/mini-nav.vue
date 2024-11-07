<script setup lang="ts">
import {useAuthStore} from "~/stores/auth";
import LogoutIcon from "~/components/icons/logout-icon.vue";
import PlanetIcon from "~/components/icons/planet-icon.vue";
import SpinnerIcon from "~/components/icons/spinner-icon.vue";
import {useRouter} from "#app";
import LogoutModal from "~/components/modals/logout-modal.vue";

const router = useRouter()
const env = useRuntimeConfig()
const auth = useAuthStore()

const logout = ref(false)
</script>

<template>
  <nav class="px-5 max-w-screen-md mx-auto py-5">
    <div class="w-full bg-widget bg-blur rounded-2xl drop-shadow-sm px-4 py-3 flex items-center justify-between space-x-2">
      <div class="flex items-center space-x-2">
        <img class="w-7 h-7 md:w-8 md:h-8 no-select pointer-events-none" src="~/assets/images/logo.svg" alt="Logo" />
        <RouterLink to="/" class="fredoka text-lg md:text-xl inline-block">MyCloud</RouterLink>
      </div>

      <div class="flex items-center space-x-2">
        <button v-tooltip="auth.user.name">
          <img class="w-7 h-7 md:w-8 md:h-8 rounded-lg" :src="`${env.public.api}/profileimage/${auth.user.image_url}`" :alt="auth.user.name" />
        </button>

        <RouterLink to="/spaces" v-tooltip="'Spaces'" class="link">
          <PlanetIcon />
        </RouterLink>

        <button @click="logout = true" v-tooltip="'Logout'" class="link">
          <LogoutIcon />
        </button>
      </div>
    </div>
  </nav>

  <logout-modal v-if="logout" @close="logout = false" />
</template>

<style scoped>
.link {
  @apply transition w-7 h-7 md:w-8 md:h-8 flex items-center justify-center bg-main-from hover:bg-main-to p-1 rounded-lg
}
</style>