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
import {onMounted} from "#imports";
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
  ogTitle: 'MyCloud - ',
  ogDescription: 'Üdvözöllek a Hungarian Vypers Gaming hivatalos weboldalán.',
  ogType: 'website',
  ogUrl: 'https://vypers.hu',
  themeColor: '#73177b',
  ogImage: 'https://vypers.hu/thumbnail/logo.png',
  ogImageAlt: 'Vypers Logo',
  ogImageWidth: '128',
  ogImageHeight: '128',
  robots: 'index, follow'
})

const socket = useSocket()
onMounted(() => {
  socket.on('download_request_finished', handleSocketDownloadRequestFinished)
  socket.connect()
})

onBeforeUnmount(() => socket.finish())
</script>