<template>
  <LoaderUtil v-if="loader.loading" />
  <NuxtLayout>
    <FadeTransition>
      <main class="text-black dark:text-white max-w-full min-h-screen" v-show="!loader.loading">
        <NuxtPage/>
      </main>
    </FadeTransition>
  </NuxtLayout>
</template>

<script setup lang="ts">
import LoaderUtil from '~/components/utils/loader-util.vue'
import {useLoaderStore} from "~/stores/loader";
import {onMounted, onBeforeUnmount, useAuthStore} from "#imports";
import FadeTransition from "~/components/transitions/fade-transition.vue";
import {useSocket} from "~/composables/useSocket";
import {handleSocketDownloadRequestFinished} from "~/scripts/socket";

const loader = useLoaderStore()

onMounted(async () => {
  await useAuthStore().handle()
  loader.finish()
})

addRouteMiddleware('global-loader', () => loader.start(), { global: true })

useHead({
  title: 'MyCloud - Easy to use, web driven filesystem',
  bodyAttrs: {
    class: 'bg-main-from bg-gradient-to-bl from-main-from to-main-to text-white',
  }
})

useSeoMeta({
  ogSiteName: 'MyCloud',
  ogTitle: 'MyCloud - Mrtn.Vip',
  ogDescription: 'Easy to use Web-Based FileSystem.',
  ogType: 'website',
  ogUrl: 'https://cloud.mrtn.vip',
  themeColor: '#dcbff5',
  robots: 'index, follow'
})

const auth = useAuthStore()
const socket = useSocket()
const socketOk = ref(false)

onMounted(() => {
  startSocket()
})

watch(auth, () => {
  if(auth.isLoggedIn) startSocket()
})

const startSocket = () => {
  if(socketOk.value) return
  if(!auth.isLoggedIn) return
  socketOk.value = true
  socket.on('download_request_finished', handleSocketDownloadRequestFinished)
  socket.connect()
}

onBeforeUnmount(() => socket.finish())
</script>