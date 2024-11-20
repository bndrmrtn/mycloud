<script setup lang="ts">
import {definePageMeta, useLoaderStore} from "#imports";
import SpaceLayout from "~/layouts/space.vue";
import type {Service} from "~/types/service";
import {fetchServiceInfo} from "~/scripts/service";
import {useToast} from "vue-toastification";

definePageMeta({
  middleware: ['auth', 'admin']
})

const conf = ref<Service|null>()

onMounted(async () => {
  const data = await fetchServiceInfo()
  useLoaderStore().loading = false
  if(data instanceof Error) return process.client ? useToast().warning(data.message) : null
  conf.value = data
})
</script>

<template>
  <SpaceLayout>
    <h1 class="fredoka text-3xl mb-5">Admin dashboard</h1>
    <section v-if="conf">
      <h2 class="fredoka text-2xl mb-5">Config.yaml</h2>
      <div class="md:grid md:grid-cols-2 md:gap-3">
        <div class="bg-widget p-2 rounded-lg">
          <p><span class="fredoka">Service version</span>: {{ conf.service.version }}</p>
        </div>
        <div class="bg-widget p-2 rounded-lg mt-2 md:mt-0">
          <p><span class="fredoka">Whitelist</span>: {{ conf.application.authorization.use_whitelist ? 'enabled' : 'disabled' }}</p>
          <p><span class="fredoka">Blacklist</span>: {{ conf.application.authorization.use_blacklist ? 'enabled' : 'disabled' }}</p>
          <p class="text-yellow-400 font-bold text-sm" v-if="!conf.application.authorization.use_blacklist && !conf.application.authorization.use_blacklist">No whitelist or blacklist enabled; all users will have access by default.</p>
          <p><span class="fredoka">Multiple admins</span>: {{ conf.application.authorization.admin.enable_multi_admin ? 'enabled' : 'disabled' }}</p>
        </div>
      </div>
    </section>
    <section class="mt-3" v-if="conf">
      <h2 class="fredoka text-2xl mb-5">Manage</h2>
      <div class="md:grid md:grid-cols-2 md:gap-3">
        <div>
          <ButtonsButtonBluish to="/admin/users">Manage users</ButtonsButtonBluish>
          <ButtonsButtonBluish to="/admin/admins" v-if="conf.application.authorization.admin.enable_multi_admin" class="mt-3">Manage admins</ButtonsButtonBluish>
        </div>
        <div class="mt-3 md:mt-0">
          <ButtonsButtonBluish v-if="conf.application.authorization.use_whitelist">Manage whitelist</ButtonsButtonBluish>
          <ButtonsButtonBluish v-if="conf.application.authorization.use_blacklist" class="mt-3">Manage blacklist</ButtonsButtonBluish>
        </div>
      </div>
    </section>
  </SpaceLayout>
</template>