<template>
   <div class="min-h-[calc(100vh-var(--nav-height))] flex items-center justify-center w-full">
     <h1 class="text-4xl md:text-5xl fredoka text-gray-800">
       Verifying login
       <spinner-icon class="!w-8 !h-8" />
     </h1>
  </div>
</template>

<script setup lang="ts">
import SpinnerIcon from "~/components/icons/spinner-icon.vue";
import {onMounted} from "vue";
import {useRoute, useRouter} from "#app";
import type {User} from "~/types/user";

const router = useRouter()
const route = useRoute()
const env = useRuntimeConfig().public

onMounted(async () => {
  let q = ""
  Object.keys(route.query).forEach(key => {
    console.log(key)
    q += `${key}=${route.query[key]}&`
  })

  const url = env.api + '/gauth?' + q.substring(0, q.length-1)
  try {
    const res = await fetch(url)
    const data = await res.json() as {user: User}
    useAuthStore().user = data.user
    await router.push('/')
  } catch (e: unknown) {
    console.error(e)
    await router.push('/')
  }
})
</script>