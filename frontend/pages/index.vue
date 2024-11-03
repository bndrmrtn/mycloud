<template>
     <div class="min-h-screen flex items-center justify-center bg-square">
          <div class="px-5">
               <h1 class="text-3xl md:text-5xl text-center inter font-bold leading-snug drop-shadow">
                    <span class="text-4xl md:text-5xl text-green-300 fredoka">MyCloud</span><br/>
                    <span class="text-purple-400">Web-Based File<span class="fredoka">System</span></span>
               </h1>
               <p class="text-center text-gray-300 mt-2 text-lg">
                 Made with
                 <HeartIcon v-tooltip="'Love'" class="text-red-400 cursor-pointer -mt-0.5" /> by
                 <a target="_blank" class="text-blue-400 hover:underline" href="https://mrtn.vip">Martin Binder</a>
               </p>

               <div class="md:grid md:grid-cols-2 md:gap-2 mt-3 md:px-5">
                    <ButtonsGithubButton />
                    <ButtonsDiscordButton title="Discord" />
               </div>
               
               <div class="md:w-1/2 mx-auto mt-3">
                 <ButtonsButtonBluish to="/spaces" v-if="auth.isLoggedIn">
                   Dashboard
                 </ButtonsButtonBluish>
                 <ButtonsGoogleButton :isLoading="progress" @click="request" v-else />
               </div>
          </div>
     </div>
</template>

<script setup lang="ts">
import {newRequest} from "~/scripts/request";
import {onMounted, useLoaderStore} from "#imports";
import {useToast} from "vue-toastification";
import HeartIcon from "~/components/icons/heart-icon.vue";

const auth = useAuthStore()

const progress = ref(false)

const request = async () => {
  progress.value = true
  try {
    const res = await newRequest('/auth-redirect')
    const data = await res.json() as {redirect_url: string}
    window.location.href = data.redirect_url
  } catch(e: unknown) {
    console.error(e)
    if(process.client) useToast().warning('Failed to create authentication url')
  } finally {
    progress.value = false
  }
}

onMounted(() => useLoaderStore().finish())
</script>