<template>
  <LoaderUtil v-if="loader.loading" />
  <NuxtLayout>
    <FadeTransition>
      <main class="text-black dark:text-white max-w-full min-h-screen" v-show="!loader.loading">
        <NuxtPage/>
      </main>
    </FadeTransition>
    <ws-indicator :active="connected" />
  </NuxtLayout>
</template>

<script setup lang="ts">
import LoaderUtil from '~/components/utils/loader-util.vue'
import {useLoaderStore} from "~/stores/loader";
import {onMounted, useAuthStore} from "#imports";
import FadeTransition from "~/components/transitions/fade-transition.vue";
import {handleOnMessage} from "~/scripts/ws";
import {useWebSocket, type UseWebSocketReturn} from "@vueuse/core";

const auth = useAuthStore()
const loader = useLoaderStore()
// https://vueuse.org/core/useWebSocket/
const connected = ref(false)
const conn = ref<UseWebSocketReturn<string>|null>(null)

onMounted(async () => {
  await auth.handle()

  if(auth.isLoggedIn) setupWS()

  loader.finish()
})

watch(auth, () => {
  if(!auth.isLoggedIn) {
    if(conn) conn.value?.close()
    conn.value = null
    connected.value = false
    return
  }
  setupWS()
})

const setupWS = () => {
  if([WebSocket.OPEN, WebSocket.CONNECTING].includes(conn.value?.status)) return
  if(conn.value || connected.value) return

  conn.value = useWebSocket<string>(useRuntimeConfig().public.ws, {
    autoReconnect: true,
    heartbeat: {
      message: JSON.stringify({type: 'echo'}),
      interval: 2500,
      pongTimeout: 1000
    },
    onMessage: handleOnMessage,
    onDisconnected: () => {
      conn.value = null
      connected.value = false
    },
    onConnected: () => connected.value = true,
  })
}

addRouteMiddleware('global-loader', () => loader.start(), { global: true })

useHead({
  htmlAttrs: {
    lang: 'en-US',
  },
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
</script>