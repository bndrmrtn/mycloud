<template>
  <ModalUtil :show="true" @close="emit('close')" :disable-close="loggingOut">
    <h2 class="text-2xl fredoka">
      Logout <LogoutIcon class="-mt-1 !w-5 !h-5" />
    </h2>

    <p class="mt-2 mb-4 text-gray-300">Are you sure you want to log out?</p>
    <div class="flex items-center space-x-2">
      <ButtonGreenish :disabled="loggingOut" @click="emit('close')"
        >Retain</ButtonGreenish
      >
      <ButtonDanger :is-loading="loggingOut" @click="logout"
        >Logout</ButtonDanger
      >
    </div>
  </ModalUtil>
</template>

<script setup lang="ts">
import ModalUtil from '~/components/utils/modal-util.vue';
import LogoutIcon from '~/components/icons/logout-icon.vue';
import ButtonGreenish from '~/components/buttons/button-greenish.vue';
import { useAuthStore } from '#imports';
import { useRouter } from '#app';
import ButtonDanger from '~/components/buttons/button-danger.vue';

const emit = defineEmits(['close']);
const router = useRouter();

const auth = useAuthStore();
const loggingOut = ref(false);

const logout = async () => {
  loggingOut.value = true;
  await auth.logout();
  loggingOut.value = false;
  await router.push('/');
  emit('close');
};
</script>
