<script setup lang="ts">
import {newRequest} from "~/scripts/request";
import type {User} from "~/types/user";
import SpinnerIcon from "~/components/icons/spinner-icon.vue";
import {useAuthStore, useLoaderStore} from "#imports";

const router = useRouter()
const route = useRoute()

onMounted(async () => {
  useLoaderStore().finish()
  let q = ""
  Object.keys(route.query).forEach(key => {
    console.log(key)
    q += `${key}=${route.query[key]}&`
  })

  const url = '/gauth?' + q.substring(0, q.length-1)
  try {
    const res = await newRequest(url)
    const data = await res.json() as {user: User}
    useAuthStore().set(data.user)
    await router.push('/')
  } catch (e: unknown) {
    console.error(e)
    await router.push('/')
  }
})
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-square">
    <div class="px-5">
      <h1 class="text-3xl md:text-5xl text-center inter font-bold leading-snug drop-shadow">
        <span class="text-green-300">Verifying login</span>
        <SpinnerIcon class="!w-7 !h-7 ml-2" />
      </h1>
    </div>
  </div>
</template>