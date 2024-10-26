<script setup lang="ts">
import {useAuthStore} from "~/stores/auth";

const open = ref(false)
const env = useRuntimeConfig()

const auth = useAuthStore()
</script>

<template>
  <nav class="bg-white md:h-[var(--nav-height)] border-gray-200 border-b bg-opacity-40 fixed top-0 w-full backdrop-blur z-50">
    <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
      <RouterLink to="/" class="flex items-center space-x-3 rtl:space-x-reverse">
        <img src="../assets/images/cloud.svg" class="h-8 w-8" alt="MyCloud">
        <span class="self-center text-2xl fredoka whitespace-nowrap text-gray-800">MyCloud</span>
      </RouterLink>
      <div class="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
        <div v-if="auth.isLoggedIn" class="flex items-center space-x-1.5 hover:bg-gray-100 rounded-lg p-2">
          <img v-if="!!auth.user?.image" class="w-6 h-6 rounded-lg" :src="`${env.public.api}/profileimage/${auth.user?.image}`" alt="User Image" />
          <p class="fredoka">{{ auth.user?.name }}</p>
        </div>
        <div v-else class="flex items-center space-x-1.5 hover:bg-gray-100 rounded-lg p-2">
          <p class="fredoka">Unauthorized</p>
        </div>
        <button @click="open = !open" type="button" class="transition inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:bg-gray-200 focus:outline-none">
          <span class="sr-only">Open main menu</span>
          <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 17 14">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 1h15M1 7h15M1 13h15"/>
          </svg>
        </button>
      </div>
      <div :class="{'hidden': !open}" class="fredoka items-center justify-between w-full md:flex md:w-auto md:order-1" id="navbar-sticky">
        <ul class="flex flex-col p-4 md:p-0 mt-4 font-medium md:space-x-8 rtl:space-x-reverse md:flex-row md:mt-0 md:border-0">
          <li>
            <RouterLink to="/" class="activeLink" aria-current="page">Home</RouterLink>
          </li>
          <li v-if="auth.isLoggedIn">
            <RouterLink to="/spaces" class="inactiveLink">Spaces</RouterLink>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.activeLink {
  @apply block py-2 px-3 text-white bg-green-400 rounded md:bg-transparent md:text-green-400 md:p-0
}

.inactiveLink {
  @apply block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:hover:text-green-300 md:p-0
}
</style>