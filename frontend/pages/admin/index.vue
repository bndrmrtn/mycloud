<script setup lang="ts">
import { definePageMeta, useLoaderStore } from '#imports';
import SpaceLayout from '~/layouts/space.vue';
import type { Service } from '~/types/service';
import { fetchServiceInfo } from '~/scripts/service';
import { useToast } from 'vue-toastification';

definePageMeta({
  middleware: ['auth', 'admin'],
});

const conf = ref<Service | null>();

const authConf = computed(() => conf.value?.application.authorization)

onMounted(async () => {
  const data = await fetchServiceInfo();
  useLoaderStore().loading = false;
  if (data instanceof Error)
    return process.client ? useToast().warning(data.message) : null;
  conf.value = data;
});
</script>

<template>
  <SpaceLayout>
    <h1 class="fredoka text-3xl mb-5">Admin dashboard</h1>
    <section v-if="conf">
      <div class="md:grid md:grid-cols-2 md:gap-4">
        <div>
          <h2 class="fredoka text-2xl mb-5">Configuration</h2>
          <div class="mt-2 bg-widget p-2 rounded-lg">
            <p>
              Service version:
              <span class="text-gray-400">{{ conf.service.version }}</span>
            </p>
            <p>
              Whitelist:
              <span class="text-gray-400">{{ authConf?.use_whitelist ? 'enabled' : 'disabled' }}</span>
            </p>
            <p>
              Blacklist:
              <span class="text-gray-400">{{ authConf?.use_blacklist ? 'enabled' : 'disabled' }}</span>
            </p>
            <p>
              Multiple admins:
              <span class="text-gray-400">{{ authConf?.admin.enable_multi_admin ? 'enabled' : 'disabled' }}</span>
            </p>
          </div>

          <div v-if="!authConf?.use_whitelist && !authConf?.use_blacklist" class="bg-widget p-2 rounded-lg text-yellow-400 mt-2">
            <p>
              No whitelist or blacklist enabled; all users will have access by default.
            </p>
          </div>
        </div>
        <div class="mt-2 md:mt-0">
          <h2 class="fredoka text-2xl mb-5">Manage</h2>
          <div class="mt-2 p-2">
            <ButtonsButtonBluish class="mb-2" to="/admin/users">
              Manage Users
            </ButtonsButtonBluish>

            <ButtonsButtonBluish v-if="authConf?.admin.enable_multi_admin" class="mb-2" to="/admin/admins">
              Manage Admins
            </ButtonsButtonBluish>

            <ButtonsButtonBluish v-if="authConf?.use_whitelist" class="mb-2" to="/admin/whitelist">
              Manage Whitelist
            </ButtonsButtonBluish>

            <ButtonsButtonBluish v-if="authConf?.use_blacklist" class="mb-2" to="/admin/blacklist">
              Manage Blacklist
            </ButtonsButtonBluish>
          </div>
        </div>
      </div>
    </section>
    <section class="mt-3">
      <admin-analytics/>
    </section>
  </SpaceLayout>
</template>
